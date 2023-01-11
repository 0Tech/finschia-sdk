package keeper_test

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func (s *KeeperTestSuite) TestQueryParams() {
	testCases := map[string]struct {
		code codes.Code
	}{
		"valid request": {},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryParamsRequest{}

			res, err := s.queryServer.Params(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			params := res.Params
			s.Require().Equal(uint32(1), params.MaxDescendants)
		})
	}
}

func (s *KeeperTestSuite) TestQueryClass() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
		"class not found": {
			classID: composable.ClassIDFromOwner(s.customer),
			code:    codes.NotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryClassRequest{
				ClassId: tc.classID,
			}

			res, err := s.queryServer.Class(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			class := res.Class
			s.Require().NotNil(class)
			s.Require().Equal(tc.classID, class.Id)
		})
	}
}

func (s *KeeperTestSuite) TestQueryClasses() {
	testCases := map[string]struct {
		code codes.Code
	}{
		"valid request": {},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryClassesRequest{}

			res, err := s.queryServer.Classes(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			classes := res.Classes
			s.Require().Len(classes, 1)
		})
	}
}

func (s *KeeperTestSuite) TestQueryTrait() {
	testCases := map[string]struct {
		classID string
		traitID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
			traitID: s.immutableTraitID,
		},
		"invalid class id": {
			traitID: s.immutableTraitID,
			code:    codes.InvalidArgument,
		},
		"invalid trait id": {
			classID: composable.ClassIDFromOwner(s.vendor),
			code:    codes.InvalidArgument,
		},
		"trait not found": {
			classID: composable.ClassIDFromOwner(s.customer),
			traitID: s.immutableTraitID,
			code:    codes.NotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryTraitRequest{
				ClassId: tc.classID,
				TraitId: tc.traitID,
			}

			res, err := s.queryServer.Trait(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			trait := res.Trait
			s.Require().NotNil(trait)
			s.Require().Equal(tc.traitID, trait.Id)
		})
	}
}

func (s *KeeperTestSuite) TestQueryTraits() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryTraitsRequest{
				ClassId: tc.classID,
			}

			res, err := s.queryServer.Traits(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			traits := res.Traits
			s.Require().Len(traits, 2)
		})
	}
}

func (s *KeeperTestSuite) TestQueryNFT() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
		"nft not found": {
			classID: composable.ClassIDFromOwner(s.customer),
			code:    codes.NotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryNFTRequest{
				ClassId: tc.classID,
				Id:      sdk.OneUint().String(),
			}

			res, err := s.queryServer.NFT(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			nft := res.Nft
			s.Require().NotNil(nft)
			s.Require().Equal(req.ClassId, nft.ClassId)
			s.Require().Equal(req.Id, nft.Id.String())
		})
	}
}

func (s *KeeperTestSuite) TestQueryNFTs() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryNFTsRequest{
				ClassId: tc.classID,
			}

			res, err := s.queryServer.NFTs(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			nfts := res.Nfts
			s.Require().Len(nfts, int(s.numNFTs)*2)
		})
	}
}

func (s *KeeperTestSuite) TestQueryProperty() {
	testCases := map[string]struct {
		classID    string
		propertyID string
		code       codes.Code
	}{
		"valid request": {
			classID:    composable.ClassIDFromOwner(s.vendor),
			propertyID: s.immutableTraitID,
		},
		"invalid class id": {
			propertyID: s.immutableTraitID,
			code:       codes.InvalidArgument,
		},
		"invalid trait id": {
			classID: composable.ClassIDFromOwner(s.vendor),
			code:    codes.InvalidArgument,
		},
		"trait not found": {
			classID:    composable.ClassIDFromOwner(s.customer),
			propertyID: s.immutableTraitID,
			code:       codes.NotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryPropertyRequest{
				ClassId:    tc.classID,
				Id:         sdk.OneUint().String(),
				PropertyId: tc.propertyID,
			}

			res, err := s.queryServer.Property(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			property := res.Property
			s.Require().NotNil(property)
			s.Require().Equal(tc.propertyID, property.Id)
		})
	}
}

func (s *KeeperTestSuite) TestQueryProperties() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryPropertiesRequest{
				ClassId: tc.classID,
				Id:      sdk.OneUint().String(),
			}

			res, err := s.queryServer.Properties(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			properties := res.Properties
			s.Require().Len(properties, 2)
		})
	}
}

func (s *KeeperTestSuite) TestQueryOwner() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
		"nft not found": {
			classID: composable.ClassIDFromOwner(s.customer),
			code:    codes.NotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryOwnerRequest{
				ClassId: tc.classID,
				Id:      sdk.OneUint().String(),
			}

			res, err := s.queryServer.Owner(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			ownerStr := res.Owner
			_, err = sdk.AccAddressFromBech32(ownerStr)
			s.Require().NoError(err)
		})
	}
}

func (s *KeeperTestSuite) TestQueryParent() {
	testCases := map[string]struct {
		classID string
		code    codes.Code
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"invalid class id": {
			code: codes.InvalidArgument,
		},
		"parent not found": {
			classID: composable.ClassIDFromOwner(s.customer),
			code:    codes.NotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.QueryParentRequest{
				ClassId: tc.classID,
				Id:      sdk.NewUint(s.numNFTs - 2).String(),
			}

			res, err := s.queryServer.Parent(sdk.WrapSDKContext(ctx), req)
			s.Require().Equal(tc.code, status.Code(err))
			if tc.code != codes.OK {
				return
			}
			s.Require().NotNil(res)

			parent := res.Parent
			err = parent.ValidateBasic()
			s.Require().NoError(err)

			s.Require().Equal(tc.classID, parent.ClassId)
			s.Require().Equal(req.Id, parent.Id.Incr().String())
		})
	}
}
