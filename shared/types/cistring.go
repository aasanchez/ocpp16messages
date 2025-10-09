package sharedtypes

import (
	"errors"
	"fmt"
)

const (
	ci20  = 20
	ci25  = 25
	ci50  = 50
	ci255 = 255
	ci500 = 500
)

var (
	errExceedsMaxLength = errors.New("exceeds maximum length")
	errInvalidASCII     = errors.New("CiString has non-printable ASCII")
)

type ciString struct {
	value string
}

func setCiString(input string, maxLen int) (ciString, error) {
	cis := ciString{value: input}

	if len(cis.value) > maxLen {
		return ciString{}, fmt.Errorf(
			"CiString len=%d > max=%d (error): %w",
			len(cis.value),
			maxLen,
			errExceedsMaxLength,
		)
	}

	for _, r := range input {
		if r < 32 || r > 126 {
			return ciString{}, errInvalidASCII
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

func SetCiString20Type(value string) (CiString20Type, error) {
	cs, err := setCiString(value, ci20)

	return CiString20Type{value: cs}, err
}

func (c CiString20Type) Value() string {
	return c.value.val()
}

type CiString25Type struct {
	value ciString
}

func SetCiString25Type(value string) (CiString25Type, error) {
	cs, err := setCiString(value, ci25)

	return CiString25Type{value: cs}, err
}

func (c CiString25Type) Value() string {
	return c.value.val()
}

type CiString50Type struct {
	value ciString
}

func SetCiString50Type(value string) (CiString50Type, error) {
	cs, err := setCiString(value, ci50)

	return CiString50Type{value: cs}, err
}

func (c CiString50Type) Value() string {
	return c.value.val()
}

type CiString255Type struct {
	value ciString
}

func SetCiString255Type(value string) (CiString255Type, error) {
	cs, err := setCiString(value, ci255)

	return CiString255Type{value: cs}, err
}

func (c CiString255Type) Value() string {
	return c.value.val()
}

type CiString500Type struct {
	value ciString
}

func SetCiString500Type(value string) (CiString500Type, error) {
	cis, err := setCiString(value, ci500)

	return CiString500Type{value: cis}, err
}

func (c CiString500Type) Value() string {
	return c.value.val()
}
