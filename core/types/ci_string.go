// Package types provides shared types, constants, and utilities used across multiple OCPP 1.6 message types.
// This file defines CiString types with length constraints, as specified in the OCPP 1.6J specification.
package types

import (
	"fmt"
)

// CiString is the interface that all CiString types implement.
type CiString interface {
	IsValid() bool
	String() string
}

// ===== CiString20 =====

type CiString20 string

func (s CiString20) IsValid() bool {
	return len(s) <= 20
}

func (s CiString20) String() string {
	return string(s)
}

func (s CiString20) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("CiString20 exceeds 20 characters: %q", s)
	}
	return []byte(s), nil
}

func (s *CiString20) UnmarshalText(data []byte) error {
	if len(data) > 20 {
		return fmt.Errorf("CiString20 exceeds 20 characters: %q", string(data))
	}
	*s = CiString20(data)
	return nil
}

// ===== CiString25 =====

type CiString25 string

func (s CiString25) IsValid() bool {
	return len(s) <= 25
}

func (s CiString25) String() string {
	return string(s)
}

func (s CiString25) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("CiString25 exceeds 25 characters: %q", s)
	}
	return []byte(s), nil
}

func (s *CiString25) UnmarshalText(data []byte) error {
	if len(data) > 25 {
		return fmt.Errorf("CiString25 exceeds 25 characters: %q", string(data))
	}
	*s = CiString25(data)
	return nil
}

// ===== CiString50 =====

type CiString50 string

func (s CiString50) IsValid() bool {
	return len(s) <= 50
}

func (s CiString50) String() string {
	return string(s)
}

func (s CiString50) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("CiString50 exceeds 50 characters: %q", s)
	}
	return []byte(s), nil
}

func (s *CiString50) UnmarshalText(data []byte) error {
	if len(data) > 50 {
		return fmt.Errorf("CiString50 exceeds 50 characters: %q", string(data))
	}
	*s = CiString50(data)
	return nil
}

// ===== CiString255 =====

type CiString255 string

func (s CiString255) IsValid() bool {
	return len(s) <= 255
}

func (s CiString255) String() string {
	return string(s)
}

func (s CiString255) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("CiString255 exceeds 255 characters: %q", s)
	}
	return []byte(s), nil
}

func (s *CiString255) UnmarshalText(data []byte) error {
	if len(data) > 255 {
		return fmt.Errorf("CiString255 exceeds 255 characters: %q", string(data))
	}
	*s = CiString255(data)
	return nil
}

// ===== CiString500 =====

type CiString500 string

func (s CiString500) IsValid() bool {
	return len(s) <= 500
}

func (s CiString500) String() string {
	return string(s)
}

func (s CiString500) MarshalText() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("CiString500 exceeds 500 characters: %q", s)
	}
	return []byte(s), nil
}

func (s *CiString500) UnmarshalText(data []byte) error {
	if len(data) > 500 {
		return fmt.Errorf("CiString500 exceeds 500 characters: %q", string(data))
	}
	*s = CiString500(data)
	return nil
}
