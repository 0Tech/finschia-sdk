package keeper

import (
	"context"
	"strings"

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
		composable.ErrInvalidTraitID: codes.InvalidArgument,
		composable.ErrInvalidNFTID:   codes.InvalidArgument,
		// composable.ErrInvalidComposition: codes.InvalidArgument,
		composable.ErrClassNotFound: codes.NotFound,
		composable.ErrTraitNotFound: codes.NotFound,
		// composable.ErrClassAlreadyExists: codes.AlreadyExists,
		composable.ErrNFTNotFound: codes.NotFound,
		// composable.ErrInsufficientNFT: codes.FailedPrecondition,
		// composable.ErrTooManyDescendants: codes.FailedPrecondition,
		composable.ErrParentNotFound: codes.NotFound,

		sdkerrors.ErrInvalidRequest: codes.InvalidArgument,
	}

	for sdkerror, code := range mapping {
		if sdkerror.Is(err) {
			return status.Error(code, err.Error())
		}
	}

	return status.Convert(err).Err()
}

const didDelimiter = ":"

// Params queries the module params.
func (s queryServer) Params(c context.Context, req *composable.QueryParamsRequest) (_ *composable.QueryParamsResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
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
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
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
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
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
func (s queryServer) Trait(c context.Context, req *composable.QueryTraitRequest) (_ *composable.QueryTraitResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := composable.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	if err := composable.ValidateTraitID(req.TraitId); err != nil {
		return nil, err
	}

	trait, err := s.keeper.GetTrait(ctx, req.ClassId, req.TraitId)
	if err != nil {
		return nil, err
	}

	return &composable.QueryTraitResponse{
		Trait: trait,
	}, nil
}

// Traits queries all traits of a class.
func (s queryServer) Traits(c context.Context, req *composable.QueryTraitsRequest) (_ *composable.QueryTraitsResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := composable.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	traitStore := prefix.NewStore(store, traitKeyPrefixOfClass(req.ClassId))

	var traits []composable.Trait
	pageRes, err := query.Paginate(traitStore, req.Pagination, func(_ []byte, value []byte) error {
		var trait composable.Trait
		s.keeper.cdc.MustUnmarshal(value, &trait)

		traits = append(traits, trait)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &composable.QueryTraitsResponse{
		Traits:     traits,
		Pagination: pageRes,
	}, nil
}

// NFT queries an nft.
func (s queryServer) NFT(c context.Context, req *composable.QueryNFTRequest) (_ *composable.QueryNFTResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := composable.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	if err := s.keeper.hasNFT(ctx, *nft); err != nil {
		return nil, err
	}

	return &composable.QueryNFTResponse{
		Nft: nft,
	}, nil
}

// NFTs queries all nfts.
func (s queryServer) NFTs(c context.Context, req *composable.QueryNFTsRequest) (_ *composable.QueryNFTsResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := composable.ValidateClassID(req.ClassId); err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	nftStore := prefix.NewStore(store, nftKeyPrefixOfClass(req.ClassId))

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
func (s queryServer) Property(c context.Context, req *composable.QueryPropertyRequest) (_ *composable.QueryPropertyResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := composable.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	if err := composable.ValidateTraitID(req.PropertyId); err != nil {
		return nil, err
	}

	property, err := s.keeper.GetProperty(ctx, *nft, req.PropertyId)
	if err != nil {
		return nil, err
	}

	return &composable.QueryPropertyResponse{
		Property: property,
	}, nil
}

// Properties queries all properties of a class.
func (s queryServer) Properties(c context.Context, req *composable.QueryPropertiesRequest) (_ *composable.QueryPropertiesResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := composable.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.keeper.storeKey)
	propertyStore := prefix.NewStore(store, propertyKeyPrefixOfNFT(nft.ClassId, nft.Id))

	var properties []composable.Property
	pageRes, err := query.Paginate(propertyStore, req.Pagination, func(_ []byte, value []byte) error {
		var property composable.Property
		s.keeper.cdc.MustUnmarshal(value, &property)

		properties = append(properties, property)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &composable.QueryPropertiesResponse{
		Properties: properties,
		Pagination: pageRes,
	}, nil
}

// Owner queries the owner of an nft.
func (s queryServer) Owner(c context.Context, req *composable.QueryOwnerRequest) (_ *composable.QueryOwnerResponse, err error) {
	defer func() { err = s.grpcError(err) }()

	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := composable.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	owner, err := s.keeper.GetRootOwner(ctx, *nft)
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
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did := strings.Join([]string{
		req.ClassId,
		req.Id,
	}, didDelimiter)

	nft, err := composable.NFTFromString(did)
	if err != nil {
		return nil, err
	}

	parent, err := s.keeper.getParent(ctx, *nft)
	if err != nil {
		return nil, err
	}

	return &composable.QueryParentResponse{
		Parent: parent,
	}, nil
}
