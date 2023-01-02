package cli_test

import (
	"fmt"

	ostcli "github.com/line/ostracon/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/line/lbm-sdk/client/flags"
	clitestutil "github.com/line/lbm-sdk/testutil/cli"
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
	"github.com/line/lbm-sdk/x/composable/client/cli"
)

func (s *CLITestSuite) TestNewQueryCmdParams() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{},
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdParams()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryParamsResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdClass() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				composable.ClassIDFromOwner(s.vendor),
			},
		},
		"invalid class id": {
			args: []string{
				"",
			},
			err: composable.ErrInvalidClassID,
		},
		"class not found": {
			args: []string{
				composable.ClassIDFromOwner(s.customer),
			},
			err: status.Error(
				codes.NotFound,
				sdkerrors.ErrKeyNotFound.Wrap(
					status.Error(
						codes.NotFound,
						composable.ErrClassNotFound.Wrap(
							composable.ClassIDFromOwner(s.customer),
						).Error(),
					).Error(),
				).Error(),
			),
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdClass()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryClassResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdClasses() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{},
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdClasses()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryClassesResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdNFT() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.vendor),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid id": {
			args: []string{
				idToString(composable.FullID{
					Id: sdk.OneUint(),
				}),
			},
			err: composable.ErrInvalidClassID,
		},
		"nft not found": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
			err: status.Error(
				codes.NotFound,
				sdkerrors.ErrKeyNotFound.Wrap(
					status.Error(
						codes.NotFound,
						composable.ErrNFTNotFound.Wrap(
							(&composable.FullID{
								ClassId: composable.ClassIDFromOwner(s.customer),
								Id:      sdk.OneUint(),
							}).String(),
						).Error(),
					).Error(),
				).Error(),
			),
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdNFT()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryNFTResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdNFTs() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				composable.ClassIDFromOwner(s.vendor),
			},
		},
		"invalid class id": {
			args: []string{
				"",
			},
			err: composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdNFTs()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryNFTsResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdOwner() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.vendor),
					Id:      sdk.OneUint(),
				}),
			},
		},
		"invalid id": {
			args: []string{
				idToString(composable.FullID{
					Id: sdk.OneUint(),
				}),
			},
			err: composable.ErrInvalidClassID,
		},
		"nft not found": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
			err: status.Error(
				codes.NotFound,
				sdkerrors.ErrKeyNotFound.Wrap(
					status.Error(
						codes.NotFound,
						composable.ErrNFTNotFound.Wrap(
							(&composable.FullID{
								ClassId: composable.ClassIDFromOwner(s.customer),
								Id:      sdk.OneUint(),
							}).String(),
						).Error(),
					).Error(),
				).Error(),
			),
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdOwner()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryOwnerResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}

func (s *CLITestSuite) TestNewQueryCmdParent() {
	val := s.network.Validators[0]

	commonArgs := []string{
		fmt.Sprintf("--%s=%d", flags.FlagHeight, s.setupHeight),
		fmt.Sprintf("--%s=json", ostcli.OutputFlag),
	}

	testCases := map[string]struct {
		args []string
		err  error
	}{
		"valid request": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.vendor),
					Id:      sdk.NewUint(s.numNFTs - 2),
				}),
			},
		},
		"invalid id": {
			args: []string{
				idToString(composable.FullID{
					Id: sdk.OneUint(),
				}),
			},
			err: composable.ErrInvalidClassID,
		},
		"parent not found": {
			args: []string{
				idToString(composable.FullID{
					ClassId: composable.ClassIDFromOwner(s.customer),
					Id:      sdk.OneUint(),
				}),
			},
			err: status.Error(
				codes.NotFound,
				sdkerrors.ErrKeyNotFound.Wrap(
					status.Error(
						codes.NotFound,
						composable.ErrParentNotFound.Wrap(
							(&composable.FullID{
								ClassId: composable.ClassIDFromOwner(s.customer),
								Id:      sdk.OneUint(),
							}).String(),
						).Error(),
					).Error(),
				).Error(),
			),
		},
	}

	for name, tc := range testCases {
		tc := tc

		s.Run(name, func() {
			cmd := cli.NewQueryCmdParent()

			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, append(tc.args, commonArgs...))
			s.Require().ErrorIs(err, tc.err)
			if tc.err != nil {
				return
			}

			var res composable.QueryParentResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &res)
			s.Require().NoError(err, out)
		})
	}
}
