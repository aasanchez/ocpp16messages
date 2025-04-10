// Package types defines validated OCPP 1.6J types used in message payloads.
package types

// IdToken represents the identifier for a user, such as a card or app code.
// According to the OCPP 1.6J spec, IdToken is a CiString20Type.
type IdToken struct {
	value CiString20Type
}

// IdTokenFromString constructs and validates an IdToken from a raw string.
func IdTokenFromString(s string) (IdToken, error) {
	ci, err := CiString20(s)
	if err != nil {
		return IdToken{}, err
	}

	return IdToken{value: ci}, nil
}

// String returns the string value of the IdToken.
func (id IdToken) String() string {
	return id.value.String()
}

// Validate returns the result of re-validating the wrapped CiString20Type.
func (id IdToken) Validate() error {
	return id.value.Validate()
}
