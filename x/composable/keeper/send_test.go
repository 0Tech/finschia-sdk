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

			// got, err := s.keeper.getOwner(ctx, fullID)
			// s.Require().NoError(err)
			// s.Require().Equal(s.customer, *got)
		})
	}
}
