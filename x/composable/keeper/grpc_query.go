package keeper

import (
	// "context"

	// sdk "github.com/line/lbm-sdk/types"
	// sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

type queryServer struct {
	keeper Keeper

	// TODO: remove it!
	composable.UnimplementedQueryServer
}

var _ composable.QueryServer = (*queryServer)(nil)

// NewQueryServer returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServer(keeper Keeper) composable.QueryServer {
	return &queryServer{
		keeper: keeper,
	}
}
