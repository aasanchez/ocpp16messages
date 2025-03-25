package types

import (
	"fmt"
)

// CiString is the interface that all CiString types implement.
type CiString interface {
	IsValid() bool
	String() string
}

// ===== Helpers =====

func validateLength(fieldName string, value string, maxLen int) error {
	if len(value) > maxLen {
		return fmt.Errorf("%s exceeds %d characters: %q", fieldName, maxLen, value)
	}
	return nil
}

func marshalCiString(fieldName string, value string, maxLen int) ([]byte, error) {
	if err := validateLength(fieldName, value, maxLen); err != nil {
		return nil, err
	}
	return []byte(value), nil
}

func unmarshalCiString(fieldName string, data []byte, maxLen int) (string, error) {
	value := string(data)
	if err := validateLength(fieldName, value, maxLen); err != nil {
		return "", err
	}
	return value, nil
}

// ===== CiString20 =====

type CiString20 string

func (s CiString20) IsValid() bool  { return len(s) <= 20 }
func (s CiString20) String() string { return string(s) }
func (s CiString20) MarshalText() ([]byte, error) {
	return marshalCiString("CiString20", string(s), 20)
}
func (s *CiString20) UnmarshalText(data []byte) error {
	str, err := unmarshalCiString("CiString20", data, 20)
	if err != nil {
		return err
	}
	*s = CiString20(str)
	return nil
}

// ===== CiString25 =====

type CiString25 string

func (s CiString25) IsValid() bool  { return len(s) <= 25 }
func (s CiString25) String() string { return string(s) }
func (s CiString25) MarshalText() ([]byte, error) {
	return marshalCiString("CiString25", string(s), 25)
}
func (s *CiString25) UnmarshalText(data []byte) error {
	str, err := unmarshalCiString("CiString25", data, 25)
	if err != nil {
		return err
	}
	*s = CiString25(str)
	return nil
}

// ===== CiString50 =====

type CiString50 string

func (s CiString50) IsValid() bool  { return len(s) <= 50 }
func (s CiString50) String() string { return string(s) }
func (s CiString50) MarshalText() ([]byte, error) {
	return marshalCiString("CiString50", string(s), 50)
}
func (s *CiString50) UnmarshalText(data []byte) error {
	str, err := unmarshalCiString("CiString50", data, 50)
	if err != nil {
		return err
	}
	*s = CiString50(str)
	return nil
}

// ===== CiString255 =====

type CiString255 string

func (s CiString255) IsValid() bool  { return len(s) <= 255 }
func (s CiString255) String() string { return string(s) }
func (s CiString255) MarshalText() ([]byte, error) {
	return marshalCiString("CiString255", string(s), 255)
}
func (s *CiString255) UnmarshalText(data []byte) error {
	str, err := unmarshalCiString("CiString255", data, 255)
	if err != nil {
		return err
	}
	*s = CiString255(str)
	return nil
}

// ===== CiString500 =====

type CiString500 string

func (s CiString500) IsValid() bool  { return len(s) <= 500 }
func (s CiString500) String() string { return string(s) }
func (s CiString500) MarshalText() ([]byte, error) {
	return marshalCiString("CiString500", string(s), 500)
}
func (s *CiString500) UnmarshalText(data []byte) error {
	str, err := unmarshalCiString("CiString500", data, 500)
	if err != nil {
		return err
	}
	*s = CiString500(str)
	return nil
}
