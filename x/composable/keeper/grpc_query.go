package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/line/lbm-sdk/store/prefix"
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/types/query"
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

func (s queryServer) grpcError(err error) error {
	if err == nil {
		return nil
	}

	mapping := map[*sdkerrors.Error]codes.Code{
		composable.ErrInvalidClassID: codes.InvalidArgument,
		composable.ErrInvalidNFTID:   codes.InvalidArgument,
		composable.ErrInvalidUriHash: codes.InvalidArgument,
		// composable.ErrInvalidComposition: codes.InvalidArgument,
		composable.ErrClassNotFound: codes.NotFound,
		// composable.ErrClassAlreadyExists: codes.AlreadyExists,
		composable.ErrNFTNotFound: codes.NotFound,
		// composable.ErrInsufficientNFT: codes.FailedPrecondition,
		// composable.ErrTooManyDescendants: codes.FailedPrecondition,
		composable.ErrParentNotFound: codes.NotFound,
	}

	for sdkerror, code := range mapping {
		if sdkerror.Is(err) {
			return status.Error(code, err.Error())
		}
	}

	return status.Convert(err).Err()
}

// Params queries the module params.
func (s queryServer) Params(c context.Context, req *composable.QueryParamsRequest) (_ *composable.QueryParamsResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	params := s.keeper.GetParams(ctx)

	return &composable.QueryParamsResponse{
		Params: params,
	}, nil
}

// Class queries an NFT class based on its id
func (s queryServer) Class(c context.Context, req *composable.QueryClassRequest) (_ *composable.QueryClassResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := composable.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	class, err := s.keeper.GetClass(ctx, req.ClassId)
	if err != nil {
		return nil, err
	}

	return &composable.QueryClassResponse{
		Class: class,
	}, nil
}

// Classes queries all NFT classes
func (s queryServer) Classes(c context.Context, req *composable.QueryClassesRequest) (_ *composable.QueryClassesResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(s.keeper.storeKey)
	classStore := prefix.NewStore(store, classKeyPrefix)

	var classes []composable.Class
	pageRes, err := query.Paginate(classStore, req.Pagination, func(_ []byte, value []byte) error {
		var class composable.Class
		s.keeper.cdc.MustUnmarshal(value, &class)

		classes = append(classes, class)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &composable.QueryClassesResponse{
		Classes:    classes,
		Pagination: pageRes,
	}, nil
}

// NFT queries an NFT based on its class and id.
func (s queryServer) NFT(c context.Context, req *composable.QueryNFTRequest) (_ *composable.QueryNFTResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := composable.NFTIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	fullID := composable.FullID{
		ClassId: req.ClassId,
		Id:      *id,
	}
	if err := fullID.ValidateBasic(); err != nil {
		return nil, err
	}

	nft, err := s.keeper.GetNFT(ctx, fullID)
	if err != nil {
		return nil, err
	}

	return &composable.QueryNFTResponse{
		Nft: nft,
	}, nil
}

// NFTs queries all NFTs of a given class
func (s queryServer) NFTs(c context.Context, req *composable.QueryNFTsRequest) (_ *composable.QueryNFTsResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := composable.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	nftStore := prefix.NewStore(store, nftKeyPrefix)

	var nfts []composable.NFT
	pageRes, err := query.Paginate(nftStore, req.Pagination, func(_ []byte, value []byte) error {
		var nft composable.NFT
		s.keeper.cdc.MustUnmarshal(value, &nft)

		nfts = append(nfts, nft)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &composable.QueryNFTsResponse{
		Nfts:       nfts,
		Pagination: pageRes,
	}, nil
}

// Owner queries the owner of the NFT based on its class and id, same as ownerOf in ERC721
func (s queryServer) Owner(c context.Context, req *composable.QueryOwnerRequest) (_ *composable.QueryOwnerResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := composable.NFTIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	fullID := composable.FullID{
		ClassId: req.ClassId,
		Id:      *id,
	}
	if err := fullID.ValidateBasic(); err != nil {
		return nil, err
	}

	owner, err := s.keeper.GetRootOwner(ctx, fullID)
	if err != nil {
		return nil, err
	}

	return &composable.QueryOwnerResponse{
		Owner: owner.String(),
	}, nil
}

// Parent queries the parent of the NFT based on its class and id
func (s queryServer) Parent(c context.Context, req *composable.QueryParentRequest) (_ *composable.QueryParentResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := composable.NFTIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	fullID := composable.FullID{
		ClassId: req.ClassId,
		Id:      *id,
	}
	if err := fullID.ValidateBasic(); err != nil {
		return nil, err
	}

	parent, err := s.keeper.getParent(ctx, fullID)
	if err != nil {
		return nil, err
	}

	return &composable.QueryParentResponse{
		Parent: parent,
	}, nil
}
