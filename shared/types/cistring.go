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
	ErrNilCiString          = errors.New("optional ciString is nil")
)

// ciString is the internal representation of a validated ASCII string.
type ciString struct {
	raw    string
	MaxLen int
}

func CiString(value string, maxLen int) (ciString, error) {
	cs := ciString{raw: value, MaxLen: maxLen}
	if err := cs.Validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

func (cs ciString) Value() string {
	return cs.raw
}

func (cs ciString) Validate() error {
	if len(cs.raw) == 0 {
		return fmt.Errorf("ciString.Validate: %w", errEmptyValueNotAllowed)
	}

	if len(cs.raw) > cs.MaxLen {
		return fmt.Errorf("ciString.Validate: %w (got length %d, max %d)", errExceedsMaxLength, len(cs.raw), cs.MaxLen)
	}

	for _, r := range cs.raw {
		if r < 32 || r > 126 {
			return fmt.Errorf("ciString.Validate: %w", errNonPrintableASCII)
		}
	}

	return nil
}

// ---- CiString20 ----

type CiString20Type struct{ value ciString }

func CiString20(value string) (CiString20Type, error) {
	cs, err := CiString(value, maxLenCiString20)

	return CiString20Type{value: cs}, err
}

func (c CiString20Type) Value() string   { return c.value.Value() }
func (c CiString20Type) Validate() error { return c.value.Validate() }

func CiString20Optional(name string, input *string) (*CiString20Type, error) {
	if input == nil {
		return nil, fmt.Errorf("%s: %w", name, ErrNilCiString)
	}

	val, err := CiString20(*input)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}

	return &val, nil
}

// ---- CiString25 ----

type CiString25Type struct{ value ciString }

func CiString25(value string) (CiString25Type, error) {
	cs, err := CiString(value, maxLenCiString25)

	return CiString25Type{value: cs}, err
}

func (c CiString25Type) Value() string   { return c.value.Value() }
func (c CiString25Type) Validate() error { return c.value.Validate() }

func CiString25Optional(name string, input *string) (*CiString25Type, error) {
	if input == nil {
		return nil, fmt.Errorf("%s: %w", name, ErrNilCiString)
	}

	val, err := CiString25(*input)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}

	return &val, nil
}

// ---- CiString50 ----

type CiString50Type struct{ value ciString }

func CiString50(value string) (CiString50Type, error) {
	cs, err := CiString(value, maxLenCiString50)

	return CiString50Type{value: cs}, err
}

func (c CiString50Type) Value() string   { return c.value.Value() }
func (c CiString50Type) Validate() error { return c.value.Validate() }

func CiString50Optional(name string, input *string) (*CiString50Type, error) {
	if input == nil {
		return nil, fmt.Errorf("%s: %w", name, ErrNilCiString)
	}

	val, err := CiString50(*input)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}

	return &val, nil
}

// ---- CiString255 ----

type CiString255Type struct{ value ciString }

func CiString255(value string) (CiString255Type, error) {
	cs, err := CiString(value, maxLenCiString255)

	return CiString255Type{value: cs}, err
}

func (c CiString255Type) Value() string   { return c.value.Value() }
func (c CiString255Type) Validate() error { return c.value.Validate() }

func CiString255Optional(name string, input *string) (*CiString255Type, error) {
	if input == nil {
		return nil, fmt.Errorf("%s: %w", name, ErrNilCiString)
	}

	val, err := CiString255(*input)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}

	return &val, nil
}

// ---- CiString500 ----

type CiString500Type struct{ value ciString }

func CiString500(value string) (CiString500Type, error) {
	cs, err := CiString(value, maxLenCiString500)

	return CiString500Type{value: cs}, err
}

func (c CiString500Type) Value() string   { return c.value.Value() }
func (c CiString500Type) Validate() error { return c.value.Validate() }

func CiString500Optional(name string, input *string) (*CiString500Type, error) {
	if input == nil {
		return nil, fmt.Errorf("%s: %w", name, ErrNilCiString)
	}

	val, err := CiString500(*input)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", name, err)
	}

	return &val, nil
}
