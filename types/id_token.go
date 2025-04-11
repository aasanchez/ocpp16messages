// Package types defines validated OCPP 1.6J types used in message payloads.
//
// These types implement the strict constraints defined by the Open Charge Point Protocol (OCPP)
// version 1.6J for fields such as strings, identifiers, and enums. All types include built-in
// validation logic to ensure message integrity and protocol compliance when interacting with
// EV chargers and central systems.
package types

// IdTokenType represents a validated user identifier used in OCPP 1.6J messages.
//
// The `idTag` or identifier token (IdToken) is a critical part of the OCPP protocol.
// It is commonly used in messages such as `Authorize.req`, `StartTransaction.req`,
// and `StopTransaction.req` to identify the user initiating or stopping a charging session.
//
// According to the OCPP 1.6J specification, the IdToken must conform to the
// CiString[20] format, which enforces the following rules:
//   - Maximum length of 20 characters.
//   - Only printable ASCII characters (decimal 32–126).
//   - Case-insensitive string semantics.
//
// Specification Reference:
//   - OCPP 1.6J, Section 5.2: Authorize.req
//   - OCPP 1.6J, Data Types: CiString[20]
type IdTokenType struct {
	value CiString20Type
}

// IdToken constructs a new IdTokenType instance from a raw string,
// applying validation rules defined for CiString[20].
//
// It returns a validated IdTokenType or an error if the input string:
//   - Is empty.
//   - Exceeds 20 characters.
//   - Contains non-printable or non-ASCII characters.
//
// Example usage:
//
//	token, err := types.IdToken("ABC123TOKEN")
//	if err != nil {
//	    return fmt.Errorf("invalid IdToken: %w", err)
//	}
func IdToken(s string) (IdTokenType, error) {
	ci, err := CiString20(s)
	if err != nil {
		return IdTokenType{}, err
	}

	return IdTokenType{value: ci}, nil
}

// String returns the underlying string representation of the IdToken.
//
// This method allows safe access to the token value when serializing or
// building OCPP-compliant payloads.
func (id IdTokenType) String() string {
	return id.value.String()
}

// Validate re-validates the IdToken value against OCPP 1.6J constraints.
//
// This is typically used when the token has been deserialized or imported
// from an external source and needs to be rechecked for compliance.
func (id IdTokenType) Validate() error {
	return id.value.Validate()
}
