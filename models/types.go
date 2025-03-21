package models

import (
	"fmt"
)

func validateCiString(value string, max int, field string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s: cannot be empty", field)
	}
	if len(value) > max {
		return fmt.Errorf("%s: exceeds %d characters", field, max)
	}
	return nil
}

type CiString20Type string

func (s CiString20Type) Validate() error {
	return validateCiString(string(s), 20, "CiString20Type")
}

type CiString25Type string

func (s CiString25Type) Validate() error {
	return validateCiString(string(s), 25, "CiString25Type")
}

type CiString50Type string

func (s CiString50Type) Validate() error {
	return validateCiString(string(s), 50, "CiString50Type")
}
