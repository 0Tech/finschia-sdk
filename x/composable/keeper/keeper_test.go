package keeper_test

import (
	"context"
	"fmt"
	"testing"

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

	classID string

	numNFTs int
}

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
}

func (s *KeeperTestSuite) SetupTest() {
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

	const maxDescendants = 3
	s.keeper.SetParams(s.ctx, composable.Params{
		MaxDescendants: maxDescendants,
	})
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
