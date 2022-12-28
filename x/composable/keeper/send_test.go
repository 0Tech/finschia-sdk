package keeper_test

import (
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func (s *KeeperTestSuite) TestSend() {
	testCases := map[string]struct {
		sender sdk.AccAddress
		id     sdk.Uint
		err    error
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

			err = s.keeper.Send(ctx, s.vendor, s.customer, fullID)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			got, err := s.keeper.GetRootOwner(ctx, fullID)
			s.Require().NoError(err)
			s.Require().Equal(s.customer, *got)
		})
	}
}

func (s *KeeperTestSuite) TestAttach() {
	testCases := map[string]struct {
		targetID sdk.Uint
		err      error
	}{
		"valid request": {
			targetID: sdk.NewUint(s.numNFTs - 1),
		},
		"insufficient nft": {
			targetID: sdk.NewUint(s.numNFTs + 1),
			err:      composable.ErrInsufficientNFT,
		},
		"too many descendants": {
			targetID: sdk.OneUint(),
			err:      composable.ErrTooManyDescendants,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			subjectID := composable.FullID{
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      sdk.NewUint(s.numNFTs),
			}
			err := subjectID.ValidateBasic()
			s.Assert().NoError(err)

			targetID := composable.FullID{
				ClassId: subjectID.ClassId,
				Id:      tc.targetID,
			}
			err = targetID.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.Attach(ctx, s.vendor, subjectID, targetID)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			// TODO: check state change
		})
	}
}

func (s *KeeperTestSuite) TestDetach() {
	testCases := map[string]struct {
		id  sdk.Uint
		err error
	}{
		"valid request": {
			id: sdk.NewUint(s.numNFTs - 2),
		},
		"insufficient nft": {
			id:  sdk.NewUint(s.numNFTs*2 - 2),
			err: composable.ErrInsufficientNFT,
		},
		"parent not found": {
			id:  sdk.OneUint(),
			err: composable.ErrParentNotFound,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			id := composable.FullID{
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := id.ValidateBasic()
			s.Assert().NoError(err)

			err = s.keeper.Detach(ctx, s.vendor, id)
			s.Require().ErrorIs(err, tc.err)
			if err != nil {
				return
			}

			// TODO: check state change
		})
	}
}
