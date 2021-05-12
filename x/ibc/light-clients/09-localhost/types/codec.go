package types

import (
	codectypes "github.com/line/lbm-sdk/v2/codec/types"
	"github.com/line/lbm-sdk/v2/x/ibc/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}