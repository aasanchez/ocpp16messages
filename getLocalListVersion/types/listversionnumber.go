package types

import (
	"fmt"
	"math"
	"strconv"

	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	// ListVersionUnsupported indicates the Charge Point does not support
	// Local Authorization Lists.
	ListVersionUnsupported = -1

	// ListVersionEmpty indicates the local authorization list is empty.
	ListVersionEmpty = 0

	// decimalBase is the base-10 radix used for formatting.
	decimalBase = 10
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

// NewListVersionNumber creates a new ListVersionNumber from an int value.
// Returns an error if the value is outside int32 range.
func NewListVersionNumber(value int) (ListVersionNumber, error) {
	if value < math.MinInt32 || value > math.MaxInt32 {
		return ListVersionNumber{}, fmt.Errorf(
			"NewListVersionNumber: "+st.ErrorFieldFormat,
			"value",
			fmt.Errorf(
				"%w: %d out of range (int32)",
				st.ErrInvalidValue,
				value,
			),
		)
	}

	return ListVersionNumber{value: int32(value)}, nil
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
