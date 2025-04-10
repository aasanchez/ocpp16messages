package types

import (
	"errors"
	"fmt"
)

const (
	maxLenCiString20  = 20
	maxLenCiString25  = 25
	maxLenCiString50  = 50
	maxLenCiString255 = 255
	maxLenCiString500 = 500
)

var (
	ErrExceedsMaxLength  = errors.New("value exceeds maximum allowed length")
	ErrNonPrintableASCII = errors.New("value contains non-printable ASCII characters")
)

type ciString struct {
	Value  string
	MaxLen int
}

func newCiString(value string, maxLen int) (ciString, error) {
	cs := ciString{Value: value, MaxLen: maxLen}
	if err := cs.validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

func (cs ciString) validate() error {
	if len(cs.Value) > cs.MaxLen {
		return fmt.Errorf("%w: actual length %d, max %d", ErrExceedsMaxLength, len(cs.Value), cs.MaxLen)
	}

	for _, r := range cs.Value {
		if r < 32 || r > 126 {
			return ErrNonPrintableASCII
		}
	}

	return nil
}

func (cs ciString) String() string {
	return cs.Value
}

// CiString20Type represents a case-insensitive ASCII string with max 20 chars.
type CiString20Type struct{ inner ciString }

func CiString20(value string) (CiString20Type, error) {
	cs, err := newCiString(value, maxLenCiString20)

	return CiString20Type{cs}, err
}
func (c CiString20Type) String() string  { return c.inner.String() }
func (c CiString20Type) Validate() error { return c.inner.validate() }

// // CiString25Type represents a case-insensitive ASCII string with max 25 chars.
// type CiString25Type struct{ inner ciString }

// func CiString25(value string) (CiString25Type, error) {
// 	cs, err := newCiString(value, maxLenCiString25)

// 	return CiString25Type{cs}, err
// }
// func (c CiString25Type) String() string  { return c.inner.String() }
// func (c CiString25Type) Validate() error { return c.inner.validate() }

// // CiString50Type represents a case-insensitive ASCII string with max 50 chars.
// type CiString50Type struct{ inner ciString }

// func CiString50(value string) (CiString50Type, error) {
// 	cs, err := newCiString(value, maxLenCiString50)

// 	return CiString50Type{cs}, err
// }
// func (c CiString50Type) String() string  { return c.inner.String() }
// func (c CiString50Type) Validate() error { return c.inner.validate() }

// // CiString255Type represents a case-insensitive ASCII string with max 255 chars.
// type CiString255Type struct{ inner ciString }

// func CiString255(value string) (CiString255Type, error) {
// 	cs, err := newCiString(value, maxLenCiString255)

// 	return CiString255Type{cs}, err
// }
// func (c CiString255Type) String() string  { return c.inner.String() }
// func (c CiString255Type) Validate() error { return c.inner.validate() }

// // CiString500Type represents a case-insensitive ASCII string with max 500 chars.
// type CiString500Type struct{ inner ciString }

// func CiString500(value string) (CiString500Type, error) {
// 	cs, err := newCiString(value, maxLenCiString500)

// 	return CiString500Type{cs}, err
// }
// func (c CiString500Type) String() string  { return c.inner.String() }
// func (c CiString500Type) Validate() error { return c.inner.validate() }
