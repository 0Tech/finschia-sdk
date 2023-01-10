package keeper_test

import (
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func (s *KeeperTestSuite) TestNewClass() {
	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.customer),
		},
		"class already exists": {
			classID: composable.ClassIDFromOwner(s.vendor),
			err:     composable.ErrClassAlreadyExists,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			class := composable.Class{
				Id: tc.classID,
			}
			err := class.ValidateBasic()
			s.Assert().NoError(err)

			traits := []composable.Trait{
				{
					Id: "uri",
				},
			}

			err = s.keeper.NewClass(ctx, class, traits)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetClass(ctx, tc.classID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(class, *got)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateClass() {
	testCases := map[string]struct {
		classID string
		err     error
	}{
		"valid request": {
			classID: composable.ClassIDFromOwner(s.vendor),
		},
		"class not found": {
			classID: composable.ClassIDFromOwner(s.customer),
			err:     composable.ErrClassNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			class := composable.Class{
				Id: tc.classID,
			}
			err := class.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.UpdateClass(ctx, class)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetClass(ctx, tc.classID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(class, *got)
		})
	}
}

func (s *KeeperTestSuite) TestMintNFT() {
	testCases := map[string]struct {
		classID    string
		propertyID string
		err        error
	}{
		"valid request": {
			classID:    composable.ClassIDFromOwner(s.vendor),
			propertyID: s.immutableTraitID,
		},
		"class not found": {
			classID:    composable.ClassIDFromOwner(s.customer),
			propertyID: s.immutableTraitID,
			err:        composable.ErrClassNotFound,
		},
		"trait not found": {
			classID:    composable.ClassIDFromOwner(s.vendor),
			propertyID: "no-such-a-trait",
			err:        composable.ErrTraitNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			err := composable.ValidateClassID(tc.classID)
			s.Assert().NoError(err)

			properties := []composable.Property{
				{
					Id:   tc.propertyID,
					Fact: randomString(32),
				},
			}
			err = composable.Properties(properties).ValidateBasic()
			s.Assert().NoError(err)

			id, err := s.keeper.MintNFT(ctx, s.vendor, tc.classID, properties)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				s.Require().Nil(id)
				return
			}
			s.Require().NotNil(id)

			nft := composable.NFT{
				ClassId: tc.classID,
				Id:      *id,
			}
			err = nft.ValidateBasic()
			s.Require().NoError(err)

			_, err = s.keeper.GetNFT(ctx, nft)
			s.Require().NoError(err)

			got, err := s.keeper.GetProperty(ctx, nft, tc.propertyID)
			s.Require().NotNil(got)
			s.Require().Equal(properties[0], *got)
		})
	}
}

func (s *KeeperTestSuite) TestBurnNFT() {
	testCases := map[string]struct {
		id  sdk.Uint
		err error
	}{
		"valid request": {
			id: sdk.OneUint(),
		},
		"insufficient nft": {
			id:  sdk.NewUint(s.numNFTs + 1),
			err: composable.ErrInsufficientNFT,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			nft := composable.NFT{
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := nft.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.BurnNFT(ctx, s.vendor, nft)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetNFT(ctx, nft)
			s.Require().Error(err)
			s.Require().Nil(got)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateNFT() {
	testCases := map[string]struct {
		id         sdk.Uint
		propertyID string
		err        error
	}{
		"valid request": {
			id:         sdk.OneUint(),
			propertyID: s.mutableTraitID,
		},
		"nft not found": {
			id:         sdk.NewUint(s.numNFTs*2 + 1),
			propertyID: s.mutableTraitID,
			err:        composable.ErrNFTNotFound,
		},
		"trait not found": {
			id:         sdk.OneUint(),
			propertyID: "no-such-a-trait",
			err:        composable.ErrTraitNotFound,
		},
		"trait immutable": {
			id:         sdk.OneUint(),
			propertyID: s.immutableTraitID,
			err:        composable.ErrTraitImmutable,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			nft := composable.NFT{
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := nft.ValidateBasic()
			s.Assert().NoError(err)

			property := composable.Property{
				Id:   tc.propertyID,
				Fact: randomString(32),
			}
			err = property.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.UpdateNFT(ctx, nft, []composable.Property{property})
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetProperty(ctx, nft, tc.propertyID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(property, *got)
		})
	}
}
