package keeper

import (
	sdk "github.com/line/lbm-sdk/types"
)

var (
	paramsKey = []byte{0x00}

	classKeyPrefix      = []byte{0x10}
	traitKeyPrefix      = []byte{0x11}
	previousIDKeyPrefix = []byte{0x12}

	nftKeyPrefix      = []byte{0x20}
	propertyKeyPrefix = []byte{0x21}
	ownerKeyPrefix    = []byte{0x22}

	parentKeyPrefix         = []byte{0x30}
	numDescendantsKeyPrefix = []byte{0x31}
)

func concatenate(prefix []byte, components ...[]byte) []byte {
	size := len(prefix) + len(components)
	for _, component := range components {
		size += len(component)
	}

	res := make([]byte, size)
	copy(res, prefix)

	cur := len(prefix)
	for _, component := range components {
		length := len(component)

		res[cur] = byte(length)
		copy(res[cur+1:], component)

		cur += 1 + length
	}

	return res
}

func split(prefix []byte, bz []byte) [][]byte {
	var res [][]byte

	for cur := len(prefix); cur < len(bz); {
		length := int(bz[cur])

		component := bz[cur+1 : cur+1+length]
		res = append(res, component)

		cur += 1 + length
	}

	return res
}

func classIDBytes(id string) []byte {
	bz := []byte(id)

	return bz
}

func traitIDBytes(id string) []byte {
	bz := []byte(id)

	return bz
}

func nftIDBytes(id sdk.Uint) []byte {
	bz, err := id.Marshal()
	if err != nil {
		panic(err)
	}

	return bz
}

func propertyIDBytes(id string) []byte {
	bz := []byte(id)

	return bz
}

func classKey(id string) []byte {
	return concatenate(
		classKeyPrefix,
		classIDBytes(id),
	)
}

func traitKey(classID string, traitID string) []byte {
	return concatenate(
		traitKeyPrefix,
		classIDBytes(classID),
		traitIDBytes(traitID),
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

func propertyKey(classID string, id sdk.Uint, propertyID string) []byte {
	return concatenate(
		propertyKeyPrefix,
		classIDBytes(classID),
		nftIDBytes(id),
		propertyIDBytes(propertyID),
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
