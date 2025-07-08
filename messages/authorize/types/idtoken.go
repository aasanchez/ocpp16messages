package authorizetypes

import (
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// IdTokenType represents the IdToken type in OCPP 1.6J.
// It is a case-insensitive string with a maximum length of 20 characters.
// See OCPP 1.6J Part 2, Section 4.1.2.1, Table 12 for details.
type IdTokenType struct {
	value sharedtypes.CiString20Type // value holds the underlying CiString20Type. (Required)
}

// IdToken creates a new IdTokenType from a CiString20Type.
// This function is used to encapsulate the IdToken string with its validation logic.
//
// Parameters:
//   s: The CiString20Type representing the IdToken.
//
// Returns:
//   IdTokenType: A new IdTokenType instance.
//   error: An error if the provided CiString20Type is invalid (e.g., too long).
func IdToken(s sharedtypes.CiString20Type) (IdTokenType, error) {
	return IdTokenType{value: s}, nil
}

// Value returns the string representation of the IdToken.
//
// Returns:
//   string: The string content of the IdToken.
func (id IdTokenType) Value() string {
	return id.value.Value()
}
