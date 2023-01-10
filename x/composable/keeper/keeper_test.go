package keeper_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	ocproto "github.com/line/ostracon/proto/ostracon/types"
	"github.com/stretchr/testify/suite"

	"github.com/line/lbm-sdk/simapp"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
	"github.com/line/lbm-sdk/x/composable/keeper"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx   sdk.Context
	goCtx context.Context

	keeper keeper.Keeper

	queryServer composable.QueryServer
	msgServer   composable.MsgServer

	vendor   sdk.AccAddress
	customer sdk.AccAddress

	mutableTraitID   string
	immutableTraitID string

	numNFTs uint64
}

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
}

func createClassIDs(size int, prefix string) []string {
	owners := createAddresses(size, prefix)
	ids := make([]string, len(owners))
	for i, owner := range owners {
		ids[i] = composable.ClassIDFromOwner(owner)
	}

	return ids
}

func randomString(size int) string {
	res := make([]rune, size)

	letters := []rune("0123456789abcdef")
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}

	return string(res)
}

func (s *KeeperTestSuite) SetupTest() {
	rand.Seed(time.Now().UnixNano())

	checkTx := false
	app := simapp.Setup(checkTx)

	s.ctx = app.BaseApp.NewContext(checkTx, ocproto.Header{})
	s.goCtx = sdk.WrapSDKContext(s.ctx)

	s.keeper = app.ComposableKeeper

	s.queryServer = keeper.NewQueryServer(s.keeper)
	s.msgServer = keeper.NewMsgServer(s.keeper)

	// create accounts
	addresses := []*sdk.AccAddress{
		&s.vendor,
		&s.customer,
	}
	for i, address := range createAddresses(len(addresses), "addr") {
		*addresses[i] = address
	}

	const maxDescendants = 1
	s.keeper.SetParams(s.ctx, composable.Params{
		MaxDescendants: maxDescendants,
	})

	// vendor creates a class
	class := composable.Class{
		Id: composable.ClassIDFromOwner(s.vendor),
	}
	err := class.ValidateBasic()
	s.Assert().NoError(err)

	s.mutableTraitID = "level"
	s.immutableTraitID = "color"

	traits := []composable.Trait{
		{
			Id:      s.mutableTraitID,
			Mutable: true,
		},
		{
			Id: s.immutableTraitID,
		},
	}

	err = s.keeper.NewClass(s.ctx, class, traits)
	s.Assert().NoError(err)

	// vendor mints nfts to all accounts by amount of numNFTs
	s.numNFTs = (maxDescendants + 1) + 1 + 1

	for _, owner := range []sdk.AccAddress{
		s.vendor,
		s.customer,
	} {
		for i := range make([]struct{}, s.numNFTs) {
			properties := []composable.Property{
				{
					Id: s.mutableTraitID,
				},
				{
					Id: s.immutableTraitID,
				},
			}

			id, err := s.keeper.MintNFT(s.ctx, owner, class.Id, properties)
			s.Assert().NoError(err)

			// each account attachs its second nft to its first nft
			if i == 1 {
				subject := composable.NFT{
					ClassId: class.Id,
					Id:      *id,
				}
				err := subject.ValidateBasic()
				s.Assert().NoError(err)

				target := composable.NFT{
					ClassId: class.Id,
					Id:      id.Decr(),
				}
				err = target.ValidateBasic()
				s.Assert().NoError(err)

				err = s.keeper.Attach(s.ctx, owner, subject, target)
				s.Assert().NoError(err)
			}
		}
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
