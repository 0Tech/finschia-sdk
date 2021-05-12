package types_test

import (
	"testing"
	"time"

	ostbytes "github.com/line/ostracon/libs/bytes"
	ostproto "github.com/line/ostracon/proto/ostracon/types"
	osttypes "github.com/line/ostracon/types"
	"github.com/stretchr/testify/suite"

	"github.com/line/lbm-sdk/v2/codec"
	"github.com/line/lbm-sdk/v2/simapp"
	sdk "github.com/line/lbm-sdk/v2/types"
	clienttypes "github.com/line/lbm-sdk/v2/x/ibc/core/02-client/types"
	ibctmtypes "github.com/line/lbm-sdk/v2/x/ibc/light-clients/07-tendermint/types"
	ibctesting "github.com/line/lbm-sdk/v2/x/ibc/testing"
	ibctestingmock "github.com/line/lbm-sdk/v2/x/ibc/testing/mock"
)

const (
	chainID                        = "gaia"
	chainIDRevision0               = "gaia-revision-0"
	chainIDRevision1               = "gaia-revision-1"
	clientID                       = "gaiamainnet"
	trustingPeriod   time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod        time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift    time.Duration = time.Second * 10
)

var (
	height          = clienttypes.NewHeight(0, 4)
	newClientHeight = clienttypes.NewHeight(1, 1)
	upgradePath     = []string{"upgrade", "upgradedIBCState"}
)

type TendermintTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	// TODO: deprecate usage in favor of testing package
	ctx        sdk.Context
	cdc        codec.Marshaler
	privVal    osttypes.PrivValidator
	valSet     *osttypes.ValidatorSet
	valsHash   ostbytes.HexBytes
	header     *ibctmtypes.Header
	now        time.Time
	headerTime time.Time
	clientTime time.Time
}

func (suite *TendermintTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	// commit some blocks so that QueryProof returns valid proof (cannot return valid query if height <= 1)
	suite.coordinator.CommitNBlocks(suite.chainA, 2)
	suite.coordinator.CommitNBlocks(suite.chainB, 2)

	// TODO: deprecate usage in favor of testing package
	checkTx := false
	app := simapp.Setup(checkTx)

	suite.cdc = app.AppCodec()

	// now is the time of the current chain, must be after the updating header
	// mocks ctx.BlockTime()
	suite.now = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	suite.clientTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	// Header time is intended to be time for any new header used for updates
	suite.headerTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)

	suite.privVal = ibctestingmock.NewPV()

	pubKey, err := suite.privVal.GetPubKey()
	suite.Require().NoError(err)

	heightMinus1 := clienttypes.NewHeight(0, height.RevisionHeight-1)

	val := osttypes.NewValidator(pubKey, 10)
	suite.valSet = osttypes.NewValidatorSet([]*osttypes.Validator{val})
	suite.valsHash = suite.valSet.Hash()
	suite.header = suite.chainA.CreateTMClientHeader(chainID, int64(height.RevisionHeight), heightMinus1, suite.now, suite.valSet, suite.valSet, []osttypes.PrivValidator{suite.privVal})
	suite.ctx = app.BaseApp.NewContext(checkTx, ostproto.Header{Height: 1, Time: suite.now})
}

func TestTendermintTestSuite(t *testing.T) {
	suite.Run(t, new(TendermintTestSuite))
}