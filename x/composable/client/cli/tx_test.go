package cli_test

import (
	"fmt"

	"github.com/line/lbm-sdk/client/flags"
	clitestutil "github.com/line/lbm-sdk/testutil/cli"
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	// txtypes "github.com/line/lbm-sdk/types/tx"
	"github.com/line/lbm-sdk/x/composable"
	"github.com/line/lbm-sdk/x/composable/client/cli"
)

// func (s *CLITestSuite) TestNewTxCmdUpdateParams() {
// 	val := s.network.Validators[0]

// 	commonArgs := []string{
// 		fmt.Sprintf("--%s", flags.FlagGenerateOnly),
// 	}

// 	testCases := map[string]struct {
// 		args  []string
// 		err error
// 	}{
// 		"valid request": {
// 			args: []string{
// 				s.authority.String(),
// 				fmt.Sprintf(`{"max_descendants": "%d"}`, composable.DefaultMaxDescendants),
// 			},
// 		},
// 		"wrong number of args": {
// 			args: []string{
// 				s.authority.String(),
// 				fmt.Sprintf(`{"max_descendants": "%d"}`, composable.DefaultMaxDescendants),
// 				"extra",
// 			},
// 			err: nil,
// 		},
// 	}

// 	for name, tc := range testCases {
// 		tc := tc

// 		s.Run(name, func() {
// 			cmd := cli.NewTxCmdUpdateParams()

// 			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
// 			s.Require().ErrorIs(err, tc.err)
// 			if tc.err != nil  {
// 				return
// 			}

// 			var res txtypes.Tx
// 			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
// 			s.Require().NoError(err, out)
// 		})
// 	}
// }

func (s *CLITestSuite) TestNewTxCmdSend() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				s.customer.String(),
				s.vendor.String(),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid id": {
			args: []string{
				s.customer.String(),
				s.vendor.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				s.vendor.String(),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdSend()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdAttach() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				s.customer.String(),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.vendor),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid subject id": {
			args: []string{
				s.customer.String(),
				"",
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.vendor),
					Id:      sdk.OneUint(),
				}),
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid target id": {
			args: []string{
				s.customer.String(),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.vendor),
					Id:      sdk.OneUint(),
				}),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdAttach()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdDetach() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				s.customer.String(),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid id": {
			args: []string{
				s.customer.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdDetach()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdNewClass() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				s.vendor.String(),
			},
		},
		"invalid msg": {
			args: []string{
				"",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdNewClass()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdUpdateClass() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				composable.ClassIDFromOwner(s.customer),
			},
		},
		"invalid class id": {
			args: []string{
				"",
			},
			err: composable.ErrInvalidClassID,
		},
		"invalid msg": {
			args: []string{
				composable.ClassIDFromOwner(s.customer),
				fmt.Sprintf("--%s=%s", cli.FlagUri, "https://ipfs.io/ipfs/tIBeTianfOX"),
			},
			err: composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdUpdateClass()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdMintNFT() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				composable.ClassIDFromOwner(s.customer),
				s.customer.String(),
			},
		},
		"invalid class id": {
			args: []string{
				"",
				s.customer.String(),
			},
			err: composable.ErrInvalidClassID,
		},
		"invalid msg": {
			args: []string{
				composable.ClassIDFromOwner(s.customer),
				s.customer.String(),
				fmt.Sprintf("--%s=%s", cli.FlagUri, "https://ipfs.io/ipfs/tIBeTianfOX"),
			},
			err: composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdMintNFT()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdBurnNFT() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				s.customer.String(),
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid id": {
			args: []string{
				s.customer.String(),
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid msg": {
			args: []string{
				"",
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdBurnNFT()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}

func (s *CLITestSuite) TestNewTxCmdUpdateNFT() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10)))),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid failing request": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid id": {
			args: []string{
				"",
			},
			err: sdkerrors.ErrInvalidType,
		},
		"invalid class id": {
			args: []string{
				idToString(composable.FullID{
					Id: sdk.OneUint(),
				}),
			},
			err: composable.ErrInvalidClassID,
		},
		"invalid msg": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
				fmt.Sprintf("--%s=%s", cli.FlagUri, "https://ipfs.io/ipfs/tIBeTianfOX"),
			},
			err: composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewTxCmdUpdateNFT()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res sdk.TxResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
			s.Require().NotZero(res.Code, out)
		})
	}
}
