// Package core provides shared types, constants, and utilities used across multiple OCPP 1.6 message types.
// This file defines CiString types with length constraints, as specified in the OCPP 1.6J specification.
package core

import (
	"fmt"
)

// CiString20 is a string with a maximum length of 20 characters.
type CiString20 string

// CiString25 is a string with a maximum length of 25 characters.
type CiString25 string

// CiString50 is a string with a maximum length of 50 characters.
type CiString50 string

// CiString255 is a string with a maximum length of 255 characters.
type CiString255 string

// IsValid checks if the CiString20 is valid (length <= 20).
func (s CiString20) IsValid() bool {
	return len(s) <= 20
}

// String returns the string value of CiString20.
func (s CiString20) String() string {
	return string(s)
}

// MarshalText implements encoding.TextMarshaler.
func (s CiString20) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("CiString20 exceeds 20 characters: %q", s)
	}
	return []byte(s), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CiString20) UnmarshalText(data []byte) error {
	if len(data) > 20 {
		return fmt.Errorf("CiString20 exceeds 20 characters: %q", string(data))
	}
	*s = CiString20(data)
	return nil
}
