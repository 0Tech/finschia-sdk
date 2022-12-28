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
				Id:      tc.classID,
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err := class.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.NewClass(ctx, class)
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
				Id:      tc.classID,
				Uri:     randomString(32),
				UriHash: randomString(32),
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

			err := composable.ValidateClassID(tc.classID)
			s.Assert().NoError(err)

			nft := composable.NFT{
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err = composable.ValidateURIHash(nft.Uri, nft.UriHash)
			s.Assert().NoError(err)

			id, err := s.keeper.MintNFT(ctx, s.vendor, tc.classID, nft)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				s.Require().Nil(id)
				return
			}
			s.Require().NotNil(id)

			nft.Id = *id
			fullID := composable.FullID{
				ClassId: tc.classID,
				Id:      nft.Id,
			}
			s.Require().NoError(fullID.ValidateBasic())

			got, err := s.keeper.GetNFT(ctx, fullID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(nft, *got)
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

			fullID := composable.FullID{
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := fullID.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.BurnNFT(ctx, s.vendor, fullID)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetNFT(ctx, fullID)
			s.Require().Error(err)
			s.Require().Nil(got)
		})
	}
}

func (s *KeeperTestSuite) TestUpdateNFT() {
	testCases := map[string]struct {
		id  sdk.Uint
		err error
	}{
		"valid request": {
			id: sdk.OneUint(),
		},
		"nft not found": {
			id:  sdk.NewUint(s.numNFTs*2 + 1),
			err: composable.ErrNFTNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			classID := composable.ClassIDFromOwner(s.vendor)
			err := composable.ValidateClassID(classID)
			s.Assert().NoError(err)

			nft := composable.NFT{
				Id:      tc.id,
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err = nft.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.UpdateNFT(ctx, classID, nft)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			fullID := composable.FullID{
				ClassId: classID,
				Id:      nft.Id,
			}
			err = fullID.ValidateBasic()
			s.Assert().NoError(err)

			got, err := s.keeper.GetNFT(ctx, fullID)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(nft, *got)
		})
	}
}
