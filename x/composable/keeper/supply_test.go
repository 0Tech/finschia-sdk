package keeper_test

import (
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
