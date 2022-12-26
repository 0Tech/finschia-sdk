package keeper

var (
	paramsKey = []byte{0x00}

	classKeyPrefix = []byte{0x10}
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

func classIDBytes(id string) []byte {
	bz := []byte(id)

	return concatenate(
		[]byte{byte(len(bz))},
		bz,
	)
}

func classKey(id string) []byte {
	return concatenate(
		classKeyPrefix,
		classIDBytes(id),
	)
}
