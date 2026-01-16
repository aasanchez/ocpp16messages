package types

import (
	"fmt"
	"strconv"
)

const (
	// ListVersionUnsupported indicates the Charge Point does not support
	// Local Authorization Lists.
	ListVersionUnsupported = -1

	// ListVersionEmpty indicates the local authorization list is empty.
	ListVersionEmpty = 0

	// decimalBase is the base-10 radix used for parsing integer strings.
	decimalBase = 10

	// bitSize32 is the bit size for int32 parsing validation.
	bitSize32 = 32
)

// ListVersionNumber represents the version number of the local authorization
// list as defined in OCPP 1.6.
//
// Special values:
//   - -1: The Charge Point does not support Local Authorization Lists
//   - 0: The local authorization list is empty
//   - >0: The version number of the local authorization list
type ListVersionNumber struct {
	value int32
}

// NewListVersionNumber creates a new ListVersionNumber from a string value.
// Returns an error if the value cannot be parsed as a valid int32.
func NewListVersionNumber(value string) (ListVersionNumber, error) {
	parsedValue, err := strconv.ParseInt(value, decimalBase, bitSize32)
	if err != nil {
		return ListVersionNumber{}, fmt.Errorf("%w", err)
	}

	return ListVersionNumber{value: int32(parsedValue)}, nil
}

// Value returns the underlying int32 value of the ListVersionNumber.
func (listVersionNumber ListVersionNumber) Value() int32 {
	return listVersionNumber.value
}

// IsUnsupported returns true if the Charge Point does not support
// Local Authorization Lists.
func (listVersionNumber ListVersionNumber) IsUnsupported() bool {
	return listVersionNumber.value == ListVersionUnsupported
}

// IsEmpty returns true if the local authorization list is empty.
func (listVersionNumber ListVersionNumber) IsEmpty() bool {
	return listVersionNumber.value == ListVersionEmpty
}

// String returns the string representation of the ListVersionNumber.
func (listVersionNumber ListVersionNumber) String() string {
	return strconv.FormatInt(int64(listVersionNumber.value), decimalBase)
}
