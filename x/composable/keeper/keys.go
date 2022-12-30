package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
)

var (
	paramsKey = []byte{0x00}

	classKeyPrefix      = []byte{0x10}
	previousIDKeyPrefix = []byte{0x11}

	nftKeyPrefix            = []byte{0x20}
	ownerKeyPrefix          = []byte{0x21}
	parentKeyPrefix         = []byte{0x22}
	numDescendantsKeyPrefix = []byte{0x23}
)

func concatenate(components ...[]byte) []byte {
	size := 0
	for _, component := range components {
		size += len(component)
	}

	res := make([]byte, 0, size)
	for _, component := range components {
		res = append(res, component...)
	}

	return res
}

func lengthPrefixedBytes(bz []byte) []byte {
	return concatenate(
		[]byte{byte(len(bz))},
		bz,
	)
}

func classIDBytes(id string) []byte {
	bz := []byte(id)

	return lengthPrefixedBytes(bz)
}

func nftIDBytes(id sdk.Uint) []byte {
	bz, err := id.Marshal()
	if err != nil {
		panic(err)
	}

	return lengthPrefixedBytes(bz)
}

func classKey(id string) []byte {
	return concatenate(
		classKeyPrefix,
		classIDBytes(id),
	)
}

func previousIDKey(classID string) []byte {
	return concatenate(
		previousIDKeyPrefix,
		classIDBytes(classID),
	)
}

func nftKeyPrefixOfClass(classID string) []byte {
	return concatenate(
		nftKeyPrefix,
		classIDBytes(classID),
	)
}

func nftKey(classID string, id sdk.Uint) []byte {
	return concatenate(
		nftKeyPrefixOfClass(classID),
		nftIDBytes(id),
	)
}

func ownerKey(classID string, id sdk.Uint) []byte {
	return concatenate(
		ownerKeyPrefix,
		classIDBytes(classID),
		nftIDBytes(id),
	)
}

func parentKey(classID string, id sdk.Uint) []byte {
	return concatenate(
		parentKeyPrefix,
		classIDBytes(classID),
		nftIDBytes(id),
	)
}

func numDescendantsKey(classID string, id sdk.Uint) []byte {
	return concatenate(
		numDescendantsKeyPrefix,
		classIDBytes(classID),
		nftIDBytes(id),
	)
}
