package cli_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	abci "github.com/line/ostracon/abci/types"
	"github.com/stretchr/testify/suite"

	"github.com/line/lbm-sdk/client/flags"
	"github.com/line/lbm-sdk/crypto/hd"
	"github.com/line/lbm-sdk/crypto/keyring"
	clitestutil "github.com/line/lbm-sdk/testutil/cli"
	"github.com/line/lbm-sdk/testutil/network"
	sdk "github.com/line/lbm-sdk/types"
	bankcli "github.com/line/lbm-sdk/x/bank/client/cli"
	"github.com/line/lbm-sdk/x/composable"
	"github.com/line/lbm-sdk/x/composable/client/cli"
)

type CLITestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	setupHeight int64

	authority sdk.AccAddress

	vendor   sdk.AccAddress
	customer sdk.AccAddress

	numNFTs uint64
}

func TestCLITestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewCLITestSuite(cfg))
}

func NewCLITestSuite(cfg network.Config) *CLITestSuite {
	return &CLITestSuite{cfg: cfg}
}

func (s *CLITestSuite) SetupSuite() {
	s.T().Log("setting up cli test suite")

	// modify genesis
	var composableData composable.GenesisState
	err := s.cfg.Codec.UnmarshalJSON(s.cfg.GenesisState[composable.ModuleName], &composableData)
	s.Assert().NoError(err)

	const maxDescendants = 1
	params := composable.Params{
		MaxDescendants: maxDescendants,
	}
	composableData.Params = params

	composableDataBz, err := s.cfg.Codec.MarshalJSON(&composableData)
	s.Assert().NoError(err)

	s.cfg.GenesisState[composable.ModuleName] = composableDataBz

	s.network = network.New(s.T(), s.cfg)
	_, err = s.network.WaitForHeight(1)
	s.Assert().NoError(err)

	// create accounts
	s.vendor = s.createAccount("vendor")
	s.customer = s.createAccount("customer")

	// vendor creates a class
	classID := composable.ClassIDFromOwner(s.vendor)
	s.newClass(s.vendor)
	s.Assert().NoError(err)

	// vendor mints nfts to all accounts by amount of numNFTs
	s.numNFTs = (maxDescendants + 1) + 1 + 1

	for _, owner := range []sdk.AccAddress{
		s.vendor,
		s.customer,
	} {
		for i := range make([]struct{}, s.numNFTs) {
			id := s.mintNFT(classID, owner)

			// each account attachs its second nft to its first nft
			if i == 1 {
				subject := composable.NFT{
					ClassId: classID,
					Id:      id,
				}
				target := composable.NFT{
					ClassId: classID,
					Id:      id.Decr(),
				}
				s.attach(owner, subject, target)
			}
		}
	}

	s.setupHeight, err = s.network.LatestHeight()
	s.Assert().NoError(err)
	s.Assert().NoError(s.network.WaitForNextBlock())
}

func (s *CLITestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))),
}

func (s *CLITestSuite) pickEvent(events []abci.Event, event proto.Message, fn func(event proto.Message)) {
	getType := func(msg proto.Message) string {
		return proto.MessageName(msg)
	}

	for _, e := range events {
		if e.Type == getType(event) {
			msg, err := sdk.ParseTypedEvent(e)
			s.Assert().NoError(err)

			fn(msg)
			return
		}
	}

	s.Assert().Failf("event not found", "%s", events)
}

// creates an account and send some coins to it for the future transactions.
func (s *CLITestSuite) createAccount(uid string) sdk.AccAddress {
	val := s.network.Validators[0]

	keyInfo, _, err := val.ClientCtx.Keyring.NewMnemonic(uid, keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Assert().NoError(err)

	addr := keyInfo.GetAddress()

	fee := sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(1000000)))
	args := append([]string{
		val.Address.String(),
		addr.String(),
		fee.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address),
	}, commonArgs...)
	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, bankcli.NewSendTxCmd(), args)
	s.Assert().NoError(err)

	var res sdk.TxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	s.Assert().NoError(err, out.String())
	s.Assert().Zero(res.Code, out.String())

	return addr
}

func nftToString(nft composable.NFT) string {
	return fmt.Sprintf("%s:%s", nft.ClassId, nft.Id)
}

func (s *CLITestSuite) attach(owner sdk.AccAddress, subject, target composable.NFT) {
	val := s.network.Validators[0]

	args := append([]string{
		owner.String(),
		subject.String(),
		target.String(),
	}, commonArgs...)

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cli.NewTxCmdAttach(), args)
	s.Assert().NoError(err)

	var res sdk.TxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	s.Assert().NoError(err, out.String())
	s.Assert().Zero(res.Code, out.String())
}

func (s *CLITestSuite) newClass(owner sdk.AccAddress) {
	val := s.network.Validators[0]

	args := append([]string{
		owner.String(),
		"[]",
	}, commonArgs...)

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cli.NewTxCmdNewClass(), args)
	s.Assert().NoError(err)

	var res sdk.TxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	s.Assert().NoError(err, out.String())
	s.Assert().Zero(res.Code, out.String())
}

func (s *CLITestSuite) mintNFT(classID string, recipient sdk.AccAddress) sdk.Uint {
	val := s.network.Validators[0]

	owner := composable.ClassOwner(classID)
	args := append([]string{
		owner.String(),
		"[]",
		recipient.String(),
	}, commonArgs...)

	out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cli.NewTxCmdMintNFT(), args)
	s.Assert().NoError(err)

	var res sdk.TxResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
	s.Assert().NoError(err, out.String())
	s.Assert().Zero(res.Code, out.String())

	var event composable.EventMintNFT
	s.pickEvent(res.Events, &event, func(e proto.Message) {
		event = *e.(*composable.EventMintNFT)
	})

	return event.Nft.Id
}
