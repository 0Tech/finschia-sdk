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

	numNFTs uint64
}

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
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
		Id:      composable.ClassIDFromOwner(s.vendor),
		Uri:     randomString(32),
		UriHash: randomString(32),
	}
	err := class.ValidateBasic()
	s.Assert().NoError(err)

	err = s.keeper.NewClass(s.ctx, class)
	s.Assert().NoError(err)

	// vendor mints nfts to all accounts by amount of numNFTs
	s.numNFTs = (maxDescendants + 1) + 1 + 1

	for _, owner := range []sdk.AccAddress{
		s.vendor,
		s.customer,
	} {
		for range make([]struct{}, s.numNFTs) {
			nft := composable.NFT{
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err := composable.ValidateURIHash(nft.Uri, nft.UriHash)
			s.Assert().NoError(err)

			_, err = s.keeper.MintNFT(s.ctx, owner, class.Id, nft)
			s.Assert().NoError(err)
		}
	}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
