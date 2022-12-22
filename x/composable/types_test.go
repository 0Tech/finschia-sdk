package composable_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/line/lbm-sdk/types"
	"github.com/line/lbm-sdk/x/composable"
)

func createAddresses(size int, prefix string) []string {
	addrs := make([]string, size)
	for i := range addrs {
		addrs[i] = sdk.AccAddress(fmt.Sprintf("%s%d", prefix, i)).String()
	}

	return addrs
}

func TestClass(t *testing.T) {
	id := createAddresses(1, "class")[0]
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"
	hash := "tIBeTianfOX"

	testCases := map[string]struct {
		id   string
		hash string
		err  error
	}{
		"valid class": {
			id:   id,
			hash: hash,
		},
		"invalid id": {
			hash: hash,
			err:  composable.ErrInvalidClassID,
		},
		"invalid uri hash": {
			id:  id,
			err: composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			class := composable.Class{
				Id:      tc.id,
				Uri:     uri,
				UriHash: tc.hash,
			}

			err := class.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestNFT(t *testing.T) {
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"
	hash := "tIBeTianfOX"

	testCases := map[string]struct {
		id   sdk.Uint
		hash string
		err  error
	}{
		"valid nft": {
			id:   sdk.OneUint(),
			hash: hash,
		},
		"invalid id": {
			id:   sdk.ZeroUint(),
			hash: hash,
			err:  composable.ErrInvalidNFTID,
		},
		"invalid uri hash": {
			id:  sdk.OneUint(),
			err: composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			nft := composable.NFT{
				Id:      tc.id,
				Uri:     uri,
				UriHash: tc.hash,
			}

			err := nft.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestFullID(t *testing.T) {
	classIDs := createAddresses(2, "class")

	testCases := map[string]struct {
		classID string
		id      sdk.Uint
		err     error
	}{
		"valid full id": {
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
			id := composable.FullID{
				ClassId: tc.classID,
				Id:      tc.id,
			}

			err := id.ValidateBasic()
			require.ErrorIs(t, err, tc.err)
		})
	}

	l := composable.FullID{
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
			r := composable.FullID{
				ClassId: tc.classID,
				Id:      tc.id,
			}

			require.NoError(t, l.ValidateBasic())
			require.NoError(t, r.ValidateBasic())
			require.Equal(t, tc.equals, l.Equal(r))
		})
	}
}

func TestURIHash(t *testing.T) {
	uri := "https://ipfs.io/ipfs/tIBeTianfOX"
	hash := "tIBeTianfOX"

	testCases := map[string]struct {
		uri  string
		hash string
		err  error
	}{
		"valid uri and hash": {
			uri:  uri,
			hash: hash,
		},
		"empty uri and empty hash": {},
		"non-empty uri but empty hash": {
			uri: uri,
			err: composable.ErrInvalidUriHash,
		},
		"empty uri but non-empty hash": {
			hash: hash,
			err:  composable.ErrInvalidUriHash,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := composable.ValidateURIHash(tc.uri, tc.hash)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
