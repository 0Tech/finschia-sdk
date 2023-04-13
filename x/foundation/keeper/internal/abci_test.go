package internal_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/foundation"
	"github.com/line/lbm-sdk/x/foundation/keeper/internal"
)

func (s *KeeperTestSuite) TestBeginBlocker() {
	for name, tc := range map[string]struct {
		taxRatio sdk.Dec
		valid    bool
	}{
		"valid ratio": {
			taxRatio: sdk.OneDec(),
			valid:    true,
		},
		"ratio > 1": {
			taxRatio: sdk.MustNewDecFromStr("1.00000001"),
		},
	} {
		s.Run(name, func() {
			ctx, _ := s.ctx.CacheContext()

			s.impl.SetParams(ctx, foundation.Params{
				FoundationTax: tc.taxRatio,
			})

			// collect
			testing := func() {
				internal.BeginBlocker(ctx, s.impl)
			}
			if !tc.valid {
				s.Require().Panics(testing)
				return
			}
			s.Require().NotPanics(testing)

			if s.deterministic {
				expectedEvents := sdk.Events{sdk.Event{Type: "coin_spent", Attributes: []abci.EventAttribute{{Key: []uint8{0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72}, Value: []uint8{0x6c, 0x69, 0x6e, 0x6b, 0x31, 0x37, 0x78, 0x70, 0x66, 0x76, 0x61, 0x6b, 0x6d, 0x32, 0x61, 0x6d, 0x67, 0x39, 0x36, 0x32, 0x79, 0x6c, 0x73, 0x36, 0x66, 0x38, 0x34, 0x7a, 0x33, 0x6b, 0x65, 0x6c, 0x6c, 0x38, 0x63, 0x35, 0x6c, 0x39, 0x68, 0x72, 0x7a, 0x73, 0x34}, Index: false}, {Key: []uint8{0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74}, Value: []uint8{0x39, 0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x73, 0x74, 0x61, 0x6b, 0x65}, Index: false}}}, sdk.Event{Type: "coin_received", Attributes: []abci.EventAttribute{{Key: []uint8{0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72}, Value: []uint8{0x6c, 0x69, 0x6e, 0x6b, 0x31, 0x76, 0x6d, 0x61, 0x66, 0x6c, 0x38, 0x66, 0x33, 0x73, 0x36, 0x75, 0x75, 0x7a, 0x77, 0x6e, 0x78, 0x6b, 0x71, 0x7a, 0x30, 0x65, 0x7a, 0x61, 0x34, 0x37, 0x76, 0x36, 0x65, 0x63, 0x6e, 0x30, 0x74, 0x75, 0x77, 0x72, 0x36, 0x79, 0x6b}, Index: false}, {Key: []uint8{0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74}, Value: []uint8{0x39, 0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x73, 0x74, 0x61, 0x6b, 0x65}, Index: false}}}, sdk.Event{Type: "transfer", Attributes: []abci.EventAttribute{{Key: []uint8{0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74}, Value: []uint8{0x6c, 0x69, 0x6e, 0x6b, 0x31, 0x76, 0x6d, 0x61, 0x66, 0x6c, 0x38, 0x66, 0x33, 0x73, 0x36, 0x75, 0x75, 0x7a, 0x77, 0x6e, 0x78, 0x6b, 0x71, 0x7a, 0x30, 0x65, 0x7a, 0x61, 0x34, 0x37, 0x76, 0x36, 0x65, 0x63, 0x6e, 0x30, 0x74, 0x75, 0x77, 0x72, 0x36, 0x79, 0x6b}, Index: false}, {Key: []uint8{0x73, 0x65, 0x6e, 0x64, 0x65, 0x72}, Value: []uint8{0x6c, 0x69, 0x6e, 0x6b, 0x31, 0x37, 0x78, 0x70, 0x66, 0x76, 0x61, 0x6b, 0x6d, 0x32, 0x61, 0x6d, 0x67, 0x39, 0x36, 0x32, 0x79, 0x6c, 0x73, 0x36, 0x66, 0x38, 0x34, 0x7a, 0x33, 0x6b, 0x65, 0x6c, 0x6c, 0x38, 0x63, 0x35, 0x6c, 0x39, 0x68, 0x72, 0x7a, 0x73, 0x34}, Index: false}, {Key: []uint8{0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74}, Value: []uint8{0x39, 0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x73, 0x74, 0x61, 0x6b, 0x65}, Index: false}}}, sdk.Event{Type: "message", Attributes: []abci.EventAttribute{{Key: []uint8{0x73, 0x65, 0x6e, 0x64, 0x65, 0x72}, Value: []uint8{0x6c, 0x69, 0x6e, 0x6b, 0x31, 0x37, 0x78, 0x70, 0x66, 0x76, 0x61, 0x6b, 0x6d, 0x32, 0x61, 0x6d, 0x67, 0x39, 0x36, 0x32, 0x79, 0x6c, 0x73, 0x36, 0x66, 0x38, 0x34, 0x7a, 0x33, 0x6b, 0x65, 0x6c, 0x6c, 0x38, 0x63, 0x35, 0x6c, 0x39, 0x68, 0x72, 0x7a, 0x73, 0x34}, Index: false}}}}
				s.Require().Equal(expectedEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (s *KeeperTestSuite) TestEndBlocker() {
	ctx, _ := s.ctx.CacheContext()

	// check preconditions
	for name, tc := range map[string]struct {
		id     uint64
		status foundation.ProposalStatus
	}{
		"active proposal": {
			s.activeProposal,
			foundation.PROPOSAL_STATUS_SUBMITTED,
		},
		"voted proposal": {
			s.votedProposal,
			foundation.PROPOSAL_STATUS_SUBMITTED,
		},
		"withdrawn proposal": {
			s.withdrawnProposal,
			foundation.PROPOSAL_STATUS_WITHDRAWN,
		},
		"invalid proposal": {
			s.invalidProposal,
			foundation.PROPOSAL_STATUS_SUBMITTED,
		},
	} {
		s.Run(name, func() {
			proposal, err := s.impl.GetProposal(ctx, tc.id)
			s.Require().NoError(err)
			s.Require().NotNil(proposal)
			s.Require().Equal(tc.status, proposal.Status)
		})
	}

	// voting periods end
	votingPeriod := s.impl.GetFoundationInfo(ctx).GetDecisionPolicy().GetVotingPeriod()
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(votingPeriod))
	internal.EndBlocker(ctx, s.impl)

	for name, tc := range map[string]struct {
		id      uint64
		removed bool
		status  foundation.ProposalStatus
	}{
		"active proposal": {
			id:     s.activeProposal,
			status: foundation.PROPOSAL_STATUS_ACCEPTED,
		},
		"voted proposal": {
			id:     s.votedProposal,
			status: foundation.PROPOSAL_STATUS_REJECTED,
		},
		"withdrawn proposal": {
			id:      s.withdrawnProposal,
			removed: true,
		},
		"invalid proposal": {
			id:     s.invalidProposal,
			status: foundation.PROPOSAL_STATUS_ACCEPTED,
		},
	} {
		s.Run(name, func() {
			proposal, err := s.impl.GetProposal(ctx, tc.id)
			if tc.removed {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			s.Require().NotNil(proposal)
			s.Require().Equal(tc.status, proposal.Status)
		})
	}

	// proposals expire
	maxExecutionPeriod := foundation.DefaultConfig().MaxExecutionPeriod
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(maxExecutionPeriod))
	internal.EndBlocker(ctx, s.impl)

	// all proposals must be pruned
	s.Require().Empty(s.impl.GetProposals(ctx))

	if s.deterministic {
		expectedEvents := sdk.Events{}
		s.Require().Equal(expectedEvents, ctx.EventManager().Events())
	}
}
