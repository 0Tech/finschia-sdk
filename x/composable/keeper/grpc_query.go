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

func nftIDFromString(str string) (*sdk.Uint, error) {
	id, err := sdk.ParseUint(str)
	if err != nil {
		return nil, composable.ErrInvalidNFTID.Wrap(err.Error())
	}

	if err := composable.ValidateNFTID(id); err != nil {
		return nil, err
	}

	return &id, nil
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

// Class queries a class.
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

// Classes queries all classes.
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

// Trait queries a trait of a class.
func (s queryServer) Trait(ctx context.Context, req *composable.QueryTraitRequest) (*composable.QueryTraitResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Trait(ctx, req)
}

// Traits queries all traits of a class.
func (s queryServer) Traits(ctx context.Context, req *composable.QueryTraitsRequest) (*composable.QueryTraitsResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Traits(ctx, req)
}

// NFT queries an nft.
func (s queryServer) NFT(c context.Context, req *composable.QueryNFTRequest) (_ *composable.QueryNFTResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := nftIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	nft := composable.NFT{
		ClassId: req.ClassId,
		Id:      *id,
	}
	if err := nft.ValidateBasic(); err != nil {
		return nil, err
	}

	if err := s.keeper.hasNFT(ctx, nft); err != nil {
		return nil, err
	}

	return &composable.QueryNFTResponse{
		Nft: &nft,
	}, nil
}

// NFTs queries all nfts.
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

// Property queries a property of a class.
func (s queryServer) Property(ctx context.Context, req *composable.QueryPropertyRequest) (*composable.QueryPropertyResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Property(ctx, req)
}

// Properties queries all properties of a class.
func (s queryServer) Properties(ctx context.Context, req *composable.QueryPropertiesRequest) (*composable.QueryPropertiesResponse, error) {
	d := composable.UnimplementedQueryServer{}
	return d.Properties(ctx, req)
}

// Owner queries the owner of an nft.
func (s queryServer) Owner(c context.Context, req *composable.QueryOwnerRequest) (_ *composable.QueryOwnerResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := nftIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	nft := composable.NFT{
		ClassId: req.ClassId,
		Id:      *id,
	}
	if err := nft.ValidateBasic(); err != nil {
		return nil, err
	}

	owner, err := s.keeper.GetRootOwner(ctx, nft)
	if err != nil {
		return nil, err
	}

	return &composable.QueryOwnerResponse{
		Owner: owner.String(),
	}, nil
}

// Parent queries the parent of an nft.
func (s queryServer) Parent(c context.Context, req *composable.QueryParentRequest) (_ *composable.QueryParentResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	id, err := nftIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	nft := composable.NFT{
		ClassId: req.ClassId,
		Id:      *id,
	}
	if err := nft.ValidateBasic(); err != nil {
		return nil, err
	}

	parent, err := s.keeper.getParent(ctx, nft)
	if err != nil {
		return nil, err
	}

	return &composable.QueryParentResponse{
		Parent: parent,
	}, nil
}
