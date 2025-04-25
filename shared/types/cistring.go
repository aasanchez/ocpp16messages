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
	errExceedsMaxLength     = errors.New("value exceeds maximum allowed length")
	errNonPrintableASCII    = errors.New("value contains non-printable ASCII characters")
	errEmptyValueNotAllowed = errors.New("value must not be empty")
)

type ciString struct {
	Value  string
	MaxLen int
}

// CiString creates a new ciString with validation.
func CiString(value string, maxLen int) (ciString, error) {
	cs := ciString{Value: value, MaxLen: maxLen}
	if err := cs.Validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

// String returns the raw string value.
func (cs ciString) String() string {
	return cs.Value
}

// Validate checks if the ciString meets length and ASCII constraints.
func (cs ciString) Validate() error {
	if len(cs.Value) == 0 {
		return fmt.Errorf("ciString.Validate: %w", errEmptyValueNotAllowed)
	}

	if len(cs.Value) > cs.MaxLen {
		return fmt.Errorf(
			"ciString.Validate: %w: got length %d, max %d",
			errExceedsMaxLength, len(cs.Value), cs.MaxLen,
		)
	}

	for _, r := range cs.Value {
		if r < 32 || r > 126 {
			return fmt.Errorf("ciString.Validate: %w", errNonPrintableASCII)
		}
	}

	return nil
}

type CiString20Type struct{ inner ciString }

func CiString20(value string) (CiString20Type, error) {
	cs, err := CiString(value, maxLenCiString20)

	return CiString20Type{inner: cs}, err
}
func (c CiString20Type) String() string  { return c.inner.String() }
func (c CiString20Type) Validate() error { return c.inner.Validate() }
func (c CiString20Type) IsValid() bool   { return c.Validate() == nil }

type CiString25Type struct{ inner ciString }

func CiString25(value string) (CiString25Type, error) {
	cs, err := CiString(value, maxLenCiString25)

	return CiString25Type{inner: cs}, err
}
func (c CiString25Type) String() string  { return c.inner.String() }
func (c CiString25Type) Validate() error { return c.inner.Validate() }
func (c CiString25Type) IsValid() bool   { return c.Validate() == nil }

type CiString50Type struct{ inner ciString }

func CiString50(value string) (CiString50Type, error) {
	cs, err := CiString(value, maxLenCiString50)

	return CiString50Type{inner: cs}, err
}

func (c CiString50Type) String() string  { return c.inner.String() }
func (c CiString50Type) Validate() error { return c.inner.Validate() }
func (c CiString50Type) IsValid() bool   { return c.Validate() == nil }

type CiString255Type struct{ inner ciString }

func CiString255(value string) (CiString255Type, error) {
	cs, err := CiString(value, maxLenCiString255)

	return CiString255Type{inner: cs}, err
}

func (c CiString255Type) String() string  { return c.inner.String() }
func (c CiString255Type) Validate() error { return c.inner.Validate() }
func (c CiString255Type) IsValid() bool   { return c.Validate() == nil }

type CiString500Type struct{ inner ciString }

func CiString500(value string) (CiString500Type, error) {
	cs, err := CiString(value, maxLenCiString500)

	return CiString500Type{inner: cs}, err
}

func (c CiString500Type) String() string  { return c.inner.String() }
func (c CiString500Type) Validate() error { return c.inner.Validate() }
func (c CiString500Type) IsValid() bool   { return c.Validate() == nil }
