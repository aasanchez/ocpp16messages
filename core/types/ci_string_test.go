package types

import (
	"strings"
	"testing"
)

// generateStr creates a string of the given length using the character 'X'
func generateStr(length int) string {
	return strings.Repeat("X", length)
}

func TestCiString20_IsValid(t *testing.T) {
	valid := CiString20(generateStr(20)) // Generates a 20-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString20(generateStr(21)) // Generates a 21-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString20_MarshalText(t *testing.T) {
	s := CiString20(generateStr(20)) // Generates a 20-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != generateStr(20) {
		t.Errorf("Expected '%s', got: %s", generateStr(20), data)
	}

	s = CiString20(generateStr(21)) // Generates a 21-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString20_UnmarshalText(t *testing.T) {
	var s CiString20
	err := s.UnmarshalText([]byte(generateStr(20))) // Generates a 20-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != generateStr(20) { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", generateStr(20), s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(21))) // Generates a 21-character string
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString25_IsValid(t *testing.T) {
	valid := CiString25(generateStr(25)) // Generates a 25-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString25(generateStr(26)) // Generates a 26-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString25_MarshalText(t *testing.T) {
	s := CiString25(generateStr(25)) // Generates a 25-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != generateStr(25) {
		t.Errorf("Expected '%s', got: %s", generateStr(25), data)
	}

	s = CiString25(generateStr(26)) // Generates a 26-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 25 characters")
	}
}

func TestCiString25_UnmarshalText(t *testing.T) {
	var s CiString25
	err := s.UnmarshalText([]byte(generateStr(25))) // Generates a 25-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != generateStr(25) { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", generateStr(25), s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(26))) // Generates a 26-character string
	if err == nil {
		t.Error("Expected error for string exceeding 25 characters")
	}
}

func TestCiString50_IsValid(t *testing.T) {
	valid := CiString50(generateStr(50)) // Generates a 50-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString50(generateStr(51)) // Generates a 51-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString50_MarshalText(t *testing.T) {
	s := CiString50(generateStr(50)) // Generates a 50-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != generateStr(50) {
		t.Errorf("Expected '%s', got: %s", generateStr(50), data)
	}

	s = CiString50(generateStr(51)) // Generates a 51-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 50 characters")
	}
}

func TestCiString50_UnmarshalText(t *testing.T) {
	var s CiString50
	err := s.UnmarshalText([]byte(generateStr(50))) // Generates a 50-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != generateStr(50) { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", generateStr(50), s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(51))) // Generates a 51-character string
	if err == nil {
		t.Error("Expected error for string exceeding 50 characters")
	}
}

func TestCiString255_IsValid(t *testing.T) {
	valid := CiString255(generateStr(255)) // Generates a 255-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString255(generateStr(256)) // Generates a 256-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString255_MarshalText(t *testing.T) {
	s := CiString255(generateStr(255)) // Generates a 255-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != generateStr(255) {
		t.Errorf("Expected '%s', got: %s", generateStr(255), data)
	}

	s = CiString255(generateStr(256)) // Generates a 256-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 255 characters")
	}
}

func TestCiString255_UnmarshalText(t *testing.T) {
	var s CiString255
	err := s.UnmarshalText([]byte(generateStr(255))) // Generates a 255-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != generateStr(255) { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", generateStr(255), s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(256))) // Generates a 256-character string
	if err == nil {
		t.Error("Expected error for string exceeding 255 characters")
	}
}

func TestCiString500_IsValid(t *testing.T) {
	valid := CiString500(generateStr(500)) // Generates a 500-character string
	if !valid.IsValid() {
		t.Errorf("Expected string %q to be valid", valid)
	}

	invalid := CiString500(generateStr(501)) // Generates a 501-character string
	if invalid.IsValid() {
		t.Errorf("Expected string %q to be invalid", invalid)
	}
}

func TestCiString500_MarshalText(t *testing.T) {
	s := CiString500(generateStr(500)) // Generates a 500-character string
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if string(data) != generateStr(500) {
		t.Errorf("Expected '%s', got: %s", generateStr(500), data)
	}

	s = CiString500(generateStr(501)) // Generates a 501-character string
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 500 characters")
	}
}

func TestCiString500_UnmarshalText(t *testing.T) {
	var s CiString500
	err := s.UnmarshalText([]byte(generateStr(500))) // Generates a 500-character string
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if s.String() != generateStr(500) { // Compare using the String() method
		t.Errorf("Expected '%s', got: %q", generateStr(500), s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(501))) // Generates a 501-character string
	if err == nil {
		t.Error("Expected error for string exceeding 500 characters")
	}
}
