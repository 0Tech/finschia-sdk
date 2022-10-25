package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	dbm "github.com/tendermint/tm-db"
	"github.com/line/ostracon/libs/log"
	ocproto "github.com/line/ostracon/proto/ostracon/types"

	"github.com/line/lbm-sdk/codec"
	codectypes "github.com/line/lbm-sdk/codec/types"
	"github.com/line/lbm-sdk/store"
	storetypes "github.com/line/lbm-sdk/store/types"
	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/foundation"
	"github.com/line/lbm-sdk/x/foundation/keeper"
)

type invariantTestSuite struct {
	suite.Suite

	ctx sdk.Context
	cdc *codec.ProtoCodec
	key *storetypes.KVStoreKey
	keeper keeper.Keeper
}

func TestInvariantTestSuite(t *testing.T) {
	suite.Run(t, new(invariantTestSuite))
}

func (s *invariantTestSuite) SetupSuite() {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	foundation.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	storeKey := sdk.NewKVStoreKey(foundation.ModuleName)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	_ = cms.LoadLatestVersion()
	s.ctx = sdk.NewContext(cms, ocproto.Header{}, false, log.NewNopLogger())

	s.cdc = cdc
	s.key = key

	s.keeper = keeper.NewKeeper(cdc, storeKey, nil, nil, nil, nil, foundation.DefaultConfig(), nil)
}

func (s *invariantTestSuite) TestTreasuryInvariant() {
	sdkCtx, _ := s.ctx.CacheContext()
	curCtx, cdc, key := sdkCtx, s.cdc, s.key

	// Group Table
	groupTable, err := orm.NewAutoUInt64Table([2]byte{keeper.GroupTablePrefix}, keeper.GroupTableSeqPrefix, &group.GroupInfo{}, cdc)
	s.Require().NoError(err)

	// Group Member Table
	groupMemberTable, err := orm.NewPrimaryKeyTable([2]byte{keeper.GroupMemberTablePrefix}, &group.GroupMember{}, cdc)
	s.Require().NoError(err)

	groupMemberByGroupIndex, err := orm.NewIndex(groupMemberTable, keeper.GroupMemberByGroupIndexPrefix, func(val interface{}) ([]interface{}, error) {
		group := val.(*group.GroupMember).GroupId
		return []interface{}{group}, nil
	}, group.GroupMember{}.GroupId)
	s.Require().NoError(err)

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	amount := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.OneInt()))
	testCases := map[string]struct {
		treasury sdk.DecCoins
		balance sdk.Coins
		valid    bool
	}{
		"invariant not broken": {
			treasury: sdk.NewDecCoinsFromCoins(amount...),
			balance: amount,
			valid: true,
		},
		"treasury differs from the balance": {
			balance: amount,
		},
	}

	for _, spec := range specs {
		cacheCurCtx, _ := curCtx.CacheContext()

		
		_, err := groupTable.Create(cacheCurCtx.KVStore(key), groupsInfo)
		s.Require().NoError(err)

		for i := 0; i < len(groupMembers); i++ {
			err := groupMemberTable.Create(cacheCurCtx.KVStore(key), groupMembers[i])
			s.Require().NoError(err)
		}

		invariant := keeper.TreasuryInvariant(
		_, broken := keeper.TreasuryInvariant(cacheCurCtx, key, *groupTable, groupMemberByGroupIndex)
		s.Require().Equal(spec.expBroken, broken)

	}
}
