package composable_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	"github.com/line/lbm-sdk/x/composable"
)

func createAddresses(size int, prefix string) []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i))
	}

	return addrs
}

func createClassIDs(size int, prefix string) []string {
	owners := createAddresses(size, prefix)
	ids := make([]string, len(owners))
	for i, owner := range owners {
		ids[i] = composable.ClassIDFromOwner(owner)
	}

	return ids
}

func TestClass(t *testing.T) {
	id := createClassIDs(1, "class")[0]

	testCases := map[string]struct {
		id  string
		err error
	}{
		"valid class": {
			id: id,
		},
		"invalid id": {
			err: composable.ErrInvalidClassID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			class := composable.Class{
				Id: tc.id,
			}

			err := class.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestTraits(t *testing.T) {
	testCases := map[string]struct {
		ids []string
		err error
	}{
		"valid traits": {},
		"invalid id": {
			ids: []string{
				"",
			},
			err: composable.ErrInvalidTraitID,
		},
		"duplicate id": {
			ids: []string{
				"uri",
				"uri",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			traits := make([]composable.Trait, len(tc.ids))
			for i, id := range tc.ids {
				traits[i] = composable.Trait{
					Id: id,
				}
			}

			err := composable.Traits(traits).ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestNFT(t *testing.T) {
	classIDs := createClassIDs(2, "class")

	testCases := map[string]struct {
		classID string
		id      sdk.Uint
		err     error
	}{
		"valid nft": {
			classID: classIDs[0],
			id:      sdk.OneUint(),
		},
		"invalid class id": {
			id:  sdk.OneUint(),
			err: composable.ErrInvalidClassID,
		},
		"invalid id": {
			classID: classIDs[0],
			id:      sdk.ZeroUint(),
			err:     composable.ErrInvalidNFTID,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			nft := composable.NFT{
				ClassId: tc.classID,
				Id:      tc.id,
			}

			err := nft.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}

	l := composable.NFT{
		ClassId: classIDs[0],
		Id:      sdk.OneUint(),
	}
	testCases2 := map[string]struct {
		classID string
		id      sdk.Uint
		equals  bool
	}{
		"equals": {
			classID: l.ClassId,
			id:      l.Id,
			equals:  true,
		},
		"different class id": {
			classID: classIDs[1],
			id:      l.Id,
		},
		"different id": {
			classID: l.ClassId,
			id:      l.Id.Incr(),
		},
	}

	for name, tc := range testCases2 {
		t.Run(name, func(t *testing.T) {
			r := composable.NFT{
				ClassId: tc.classID,
				Id:      tc.id,
			}

			require.NoError(t, l.ValidateBasic())
			require.NoError(t, r.ValidateBasic())
			require.Equal(t, tc.equals, l.Equal(r))
		})
	}
}

func TestProperties(t *testing.T) {
	testCases := map[string]struct {
		ids []string
		err error
	}{
		"valid properties": {},
		"invalid id": {
			ids: []string{
				"",
			},
			err: composable.ErrInvalidTraitID,
		},
		"duplicate id": {
			ids: []string{
				"uri",
				"uri",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			properties := make([]composable.Property, len(tc.ids))
			for i, id := range tc.ids {
				properties[i] = composable.Property{
					Id: id,
				}
			}

			err := composable.Properties(properties).ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}
