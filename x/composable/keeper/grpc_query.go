package keeper

import (
	"context"

	// sdk "github.com/line/lbm-sdk/types"
	// sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

type queryServer struct {
	keeper Keeper
}

var _ composable.QueryServer = (*queryServer)(nil)

// NewQueryServer returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServer(keeper Keeper) composable.QueryServer {
	return &queryServer{
		keeper: keeper,
	}
}

// Params queries the module params.
func (s queryServer) Params(c context.Context, req *composable.QueryParamsRequest) (*composable.QueryParamsResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Params(c, req)
}

// Class queries an NFT class based on its id
func (s queryServer) Class(c context.Context, req *composable.QueryClassRequest) (*composable.QueryClassResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Class(c, req)
}

// Classes queries all NFT classes
func (s queryServer) Classes(c context.Context, req *composable.QueryClassesRequest) (*composable.QueryClassesResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Classes(c, req)
}

// NFT queries an NFT based on its class and id.
func (s queryServer) NFT(c context.Context, req *composable.QueryNFTRequest) (*composable.QueryNFTResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.NFT(c, req)
}

// NFTs queries all NFTs of a given class
func (s queryServer) NFTs(c context.Context, req *composable.QueryNFTsRequest) (*composable.QueryNFTsResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.NFTs(c, req)
}

// Owner queries the owner of the NFT based on its class and id, same as ownerOf in ERC721
func (s queryServer) Owner(c context.Context, req *composable.QueryOwnerRequest) (*composable.QueryOwnerResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Owner(c, req)
}

// Parent queries the parent of the NFT based on its class and id
func (s queryServer) Parent(c context.Context, req *composable.QueryParentRequest) (*composable.QueryParentResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Parent(c, req)
}
