package types

import (
	"crypto/rand"
	"strings"
	"testing"
)

const (
	expectedValidMsg   = "Expected string %q to be valid"
	expectedInvalidMsg = "Expected string %q to be invalid"
	expectedNoErrorMsg = "Expected no error, got: %v"
	expectedStringMsg  = "Expected '%s', got: %s"
	expectedStringQMsg = "Expected '%s', got: %q"
)

func generateStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	sb.Grow(length)

	byteBuf := make([]byte, length)
	_, err := rand.Read(byteBuf)
	if err != nil {
		panic("failed to generate secure random string: " + err.Error())
	}

	for _, b := range byteBuf {
		sb.WriteByte(charset[int(b)%len(charset)])
	}
	return sb.String()
}

func TestCiString20IsValid(t *testing.T) {
	valid := CiString20(generateStr(20))
	if !valid.IsValid() {
		t.Errorf(expectedValidMsg, valid)
	}

	invalid := CiString20(generateStr(21))
	if invalid.IsValid() {
		t.Errorf(expectedInvalidMsg, invalid)
	}
}

func TestCiString20MarshalText(t *testing.T) {
	string20 := generateStr(20)
	s := CiString20(string20)
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if string(data) != string20 {
		t.Errorf(expectedStringMsg, string20, data)
	}

	s = CiString20(generateStr(21))
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString20UnmarshalText(t *testing.T) {
	string20 := generateStr(20)
	var s CiString20
	err := s.UnmarshalText([]byte(string20))
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if s.String() != string20 {
		t.Errorf(expectedStringQMsg, string20, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(21)))
	if err == nil {
		t.Error("Expected error for string exceeding 20 characters")
	}
}

func TestCiString25IsValid(t *testing.T) {
	valid := CiString25(generateStr(25))
	if !valid.IsValid() {
		t.Errorf(expectedValidMsg, valid)
	}

	invalid := CiString25(generateStr(26))
	if invalid.IsValid() {
		t.Errorf(expectedInvalidMsg, invalid)
	}
}

func TestCiString25MarshalText(t *testing.T) {
	string25 := generateStr(25)
	s := CiString25(string25)
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if string(data) != string25 {
		t.Errorf(expectedStringMsg, string25, data)
	}

	s = CiString25(generateStr(26))
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 25 characters")
	}
}

func TestCiString25UnmarshalText(t *testing.T) {
	string25 := generateStr(25)
	var s CiString25
	err := s.UnmarshalText([]byte(string25))
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if s.String() != string25 {
		t.Errorf(expectedStringQMsg, string25, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(26)))
	if err == nil {
		t.Error("Expected error for string exceeding 25 characters")
	}
}

func TestCiString50IsValid(t *testing.T) {
	valid := CiString50(generateStr(50))
	if !valid.IsValid() {
		t.Errorf(expectedValidMsg, valid)
	}

	invalid := CiString50(generateStr(51))
	if invalid.IsValid() {
		t.Errorf(expectedInvalidMsg, invalid)
	}
}

func TestCiString50MarshalText(t *testing.T) {
	string50 := generateStr(50)
	s := CiString50(string50)
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if string(data) != string50 {
		t.Errorf(expectedStringMsg, string50, data)
	}

	s = CiString50(generateStr(51))
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 50 characters")
	}
}

func TestCiString50UnmarshalText(t *testing.T) {
	string50 := generateStr(50)
	var s CiString50
	err := s.UnmarshalText([]byte(string50))
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if s.String() != string50 {
		t.Errorf(expectedStringQMsg, string50, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(51)))
	if err == nil {
		t.Error("Expected error for string exceeding 50 characters")
	}
}

func TestCiString255IsValid(t *testing.T) {
	valid := CiString255(generateStr(255))
	if !valid.IsValid() {
		t.Errorf(expectedValidMsg, valid)
	}

	invalid := CiString255(generateStr(256))
	if invalid.IsValid() {
		t.Errorf(expectedInvalidMsg, invalid)
	}
}

func TestCiString255MarshalText(t *testing.T) {
	string255 := generateStr(255)
	s := CiString255(string255)
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if string(data) != string255 {
		t.Errorf(expectedStringMsg, string255, data)
	}

	s = CiString255(generateStr(256))
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 255 characters")
	}
}

func TestCiString255UnmarshalText(t *testing.T) {
	string255 := generateStr(255)
	var s CiString255
	err := s.UnmarshalText([]byte(string255))
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if s.String() != string255 {
		t.Errorf(expectedStringQMsg, string255, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(256)))
	if err == nil {
		t.Error("Expected error for string exceeding 255 characters")
	}
}

func TestCiString500IsValid(t *testing.T) {
	valid := CiString500(generateStr(500))
	if !valid.IsValid() {
		t.Errorf(expectedValidMsg, valid)
	}

	invalid := CiString500(generateStr(501))
	if invalid.IsValid() {
		t.Errorf(expectedInvalidMsg, invalid)
	}
}

func TestCiString500MarshalText(t *testing.T) {
	string500 := generateStr(500)
	s := CiString500(string500)
	data, err := s.MarshalText()
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if string(data) != string500 {
		t.Errorf(expectedStringMsg, string500, data)
	}

	s = CiString500(generateStr(501))
	_, err = s.MarshalText()
	if err == nil {
		t.Error("Expected error for string exceeding 500 characters")
	}
}

func TestCiString500UnmarshalText(t *testing.T) {
	string500 := generateStr(500)
	var s CiString500
	err := s.UnmarshalText([]byte(string500))
	if err != nil {
		t.Fatalf(expectedNoErrorMsg, err)
	}
	if s.String() != string500 {
		t.Errorf(expectedStringQMsg, string500, s.String())
	}

	err = s.UnmarshalText([]byte(generateStr(501)))
	if err == nil {
		t.Error("Expected error for string exceeding 500 characters")
	}
}
