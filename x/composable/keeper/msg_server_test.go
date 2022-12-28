package keeper_test

import (
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func (s *KeeperTestSuite) TestMsgSend() {
	testCases := map[string]struct {
		id  sdk.Uint
		err error
	}{
		"valid request": {
			id: sdk.OneUint(),
		},
		"insufficient funds": {
			id:  sdk.NewUint(s.numNFTs + 1),
			err: composable.ErrInsufficientNFT,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.MsgSend{
				Sender:    s.vendor.String(),
				Recipient: s.customer.String(),
				ClassId:   composable.ClassIDFromOwner(s.vendor),
				Id:        tc.id,
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.Send(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgAttach() {
	testCases := map[string]struct {
		targetID sdk.Uint
		err      error
	}{
		"valid request": {
			targetID: sdk.NewUint(s.numNFTs - 1),
		},
		"insufficient funds": {
			targetID: sdk.NewUint(s.numNFTs + 1),
			err:      composable.ErrInsufficientNFT,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			classID := composable.ClassIDFromOwner(s.vendor)
			req := &composable.MsgAttach{
				Owner:          s.vendor.String(),
				SubjectClassId: classID,
				SubjectId:      sdk.NewUint(s.numNFTs),
				TargetClassId:  classID,
				TargetId:       tc.targetID,
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.Attach(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgDetach() {
	testCases := map[string]struct {
		id  sdk.Uint
		err error
	}{
		"valid request": {
			id: sdk.NewUint(s.numNFTs - 2),
		},
		"insufficient funds": {
			id:  sdk.NewUint(s.numNFTs*2 - 2),
			err: composable.ErrInsufficientNFT,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.MsgDetach{
				Owner:   s.vendor.String(),
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.Detach(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgNewClass() {
	testCases := map[string]struct {
		owner sdk.AccAddress
		err   error
	}{
		"valid request": {
			owner: s.customer,
		},
		"class already exists": {
			owner: s.vendor,
			err:   composable.ErrClassAlreadyExists,
		},
	}

	for name, tc := range testCases {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			req := &composable.MsgNewClass{
				Owner:   tc.owner.String(),
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.NewClass(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgUpdateClass() {
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

			req := &composable.MsgUpdateClass{
				ClassId: tc.classID,
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.UpdateClass(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgMintNFT() {
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

			req := &composable.MsgMintNFT{
				ClassId:   tc.classID,
				Uri:       randomString(32),
				UriHash:   randomString(32),
				Recipient: s.customer.String(),
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.MintNFT(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgBurnNFT() {
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

			req := &composable.MsgBurnNFT{
				Owner:   s.vendor.String(),
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.BurnNFT(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}

func (s *KeeperTestSuite) TestMsgUpdateNFT() {
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

			req := &composable.MsgUpdateNFT{
				ClassId: composable.ClassIDFromOwner(s.vendor),
				Id:      tc.id,
				Uri:     randomString(32),
				UriHash: randomString(32),
			}
			err := req.ValidateBasic()
			s.Assert().NoError(err)

			res, err := s.msgServer.UpdateNFT(sdk.WrapSDKContext(ctx), req)
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}
			s.Require().NotNil(res)
		})
	}
}
