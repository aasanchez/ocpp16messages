package types

import (
	"fmt"
)

const (
	maxLenCiString20  = 20
	maxLenCiString25  = 25
	maxLenCiString50  = 50
	maxLenCiString255 = 255
	maxLenCiString500 = 500
)

type ciString struct {
	raw    string
	MaxLen int
}

func SetCiString(value string, maxLen int) (ciString, error) {
	cs := ciString{raw: value, MaxLen: maxLen}
	if err := cs.validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

func (cs ciString) value() string {
	return cs.raw
}

func (cs ciString) validate() error {
	if len(cs.raw) == 0 {
		return fmt.Errorf("ciString.Validate: %w", ErrEmptyValueNotAllowed)
	}

	if len(cs.raw) > cs.MaxLen {
		return fmt.Errorf("ciString.Validate: %w (got length %d, max %d)", ErrExceedsMaxLength, len(cs.raw), cs.MaxLen)
	}

	for _, r := range cs.raw {
		if r < 32 || r > 126 {
			return fmt.Errorf("ciString.Validate: %w", ErrNonPrintableASCII)
		}
	}

	return nil
}

// ---- CiString20 ----

type CiString20 struct{ value ciString }

func (c CiString20) Value() string   { return c.value.value() }
func (c CiString20) Validate() error { return c.value.validate() }
func SetCiString20(value string) (CiString20, error) {
	cs, err := SetCiString(value, maxLenCiString20)

	return CiString20{value: cs}, err
}

// ---- CiString25 ----

type CiString25 struct{ value ciString }

func (c CiString25) Value() string   { return c.value.value() }
func (c CiString25) Validate() error { return c.value.validate() }
func SetCiString25(value string) (CiString25, error) {
	cs, err := SetCiString(value, maxLenCiString25)

	return CiString25{value: cs}, err
}

// ---- CiString50 ----

type CiString50 struct{ value ciString }

func (c CiString50) Value() string   { return c.value.value() }
func (c CiString50) Validate() error { return c.value.validate() }
func SetCiString50(value string) (CiString50, error) {
	cs, err := SetCiString(value, maxLenCiString50)

	return CiString50{value: cs}, err
}

// ---- CiString255 ----

type CiString255 struct{ value ciString }

func (c CiString255) Value() string   { return c.value.value() }
func (c CiString255) Validate() error { return c.value.validate() }
func SetCiString255(value string) (CiString255, error) {
	cs, err := SetCiString(value, maxLenCiString255)

	return CiString255{value: cs}, err
}

// ---- CiString500 ----

type CiString500 struct{ value ciString }

func (c CiString500) Value() string   { return c.value.value() }
func (c CiString500) Validate() error { return c.value.validate() }
func SetCiString500(value string) (CiString500, error) {
	cs, err := SetCiString(value, maxLenCiString500)

	return CiString500{value: cs}, err
}
