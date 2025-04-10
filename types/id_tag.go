package types

// idTag represents an internal OCPP 1.6J authorization tag used for operations like Authorize.req.
// This type wraps IdToken for internal use and is kept unexported.
type IdTagType struct {
	value IdTokenType
}

// IdTagFromString validates the string and returns an idTag.
// It uses IdToken's validation internally to ensure it meets OCPP constraints.
func IdTag(s string) (IdTagType, error) {
	idToken, err := IdToken(s) // Re-use IdTokenFromString for validation
	if err != nil {
		return IdTagType{}, err
	}

	return IdTagType{value: idToken}, nil
}

// String returns the raw string value of the idTag.
func (id IdTagType) String() string {
	return id.value.String()
}

// Validate re-validates the IdToken wrapped inside the idTag.
func (id IdTagType) Validate() error {
	return id.value.Validate()
}
