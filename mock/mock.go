// Package mock offers mock implementations of interfaces defined in types.go.
// This allows complete control over input / output for any given method that consumes a given type.
package mock

import (
	"github.com/sony/sonyflake/types"
)

// NewSuccessfulInterfaceAddrs returns a single private IP address.
func NewSuccessfulInterfaceAddrs() types.InterfaceAddrs {
	_ = "STUB: not implemented"
	return *new(types.InterfaceAddrs)
}

// NewFailingInterfaceAddrs returns an error.
func NewFailingInterfaceAddrs() types.InterfaceAddrs {
	_ = "STUB: not implemented"
	return *new(types.InterfaceAddrs)
}

// NewNilInterfaceAddrs returns an empty slice of addresses.
func NewNilInterfaceAddrs() types.InterfaceAddrs {
	_ = "STUB: not implemented"
	return *new(types.InterfaceAddrs)
}
