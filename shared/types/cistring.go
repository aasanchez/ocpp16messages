package sharedtypes

import (
	"errors"
	"fmt"
)

var errExceedsMaxLength = errors.New("exceeds maximum length")

type ciString struct {
	value string
}

func setCiString(input string, maxLen int) (ciString, error) {
	cis := ciString{value: input}

	if len(cis.value) > maxLen {
		return ciString{}, fmt.Errorf(
			"ciString Error on Construc (len=%d, max=%d): %w",
			len(cis.value),
			maxLen,
			errExceedsMaxLength,
		)
	}

	for _, r := range input {
		if r < 32 || r > 126 {
			return ciString{}, fmt.Errorf("ciString.Validate: %w", ErrNonPrintableASCII)
		}
	}

	return cis, nil
}

func (cis ciString) val() string {
	return cis.value
}

type CiString20Type struct {
	value ciString
}

func (c CiString20Type) Value() string {
	return c.value.val()
}

func SetCiString20Type(value string) (CiString20Type, error) {
	cs, err := setCiString(value, 20)

	return CiString20Type{value: cs}, err
}

type CiString25Type struct {
	value ciString
}

func (c CiString25Type) Value() string {
	return c.value.val()
}

func SetCiString25Type(value string) (CiString25Type, error) {
	cs, err := setCiString(value, 25)

	return CiString25Type{value: cs}, err
}

type CiString50Type struct {
	value ciString
}

func (c CiString50Type) Value() string {
	return c.value.val()
}

func SetCiString50Type(value string) (CiString50Type, error) {
	cs, err := setCiString(value, 50)

	return CiString50Type{value: cs}, err
}

type CiString255Type struct {
	value ciString
}

func (c CiString255Type) Value() string {
	return c.value.val()
}

func SetCiString255Type(value string) (CiString255Type, error) {
	cs, err := setCiString(value, 255)

	return CiString255Type{value: cs}, err
}

type CiString500Type struct {
	value ciString
}

func (c CiString500Type) Value() string {
	return c.value.val()
}

func SetCiString500Type(value string) (CiString500Type, error) {
	cis, err := setCiString(value, 500)

	return CiString500Type{value: cis}, err
}
