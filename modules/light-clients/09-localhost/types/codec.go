package types

import (
	codectypes "github.com/cheqd/cosmos-sdk/codec/types"

	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
}
