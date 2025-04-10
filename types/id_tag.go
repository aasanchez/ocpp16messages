package types

// IdTag represents a validated authorization tag used in OCPP 1.6J.
type IdTagType struct {
	value IdTokenType
}

// IdTag returns a validated IdTag from a raw string.
func IdTag(s string) (IdTagType, error) {
	idToken, err := IdToken(s)
	if err != nil {
		return IdTagType{}, err
	}

	return IdTagType{value: idToken}, nil
}

// String returns the raw string.
func (id IdTagType) String() string {
	return id.value.String()
}

// Validate re-validates the wrapped token.
func (id IdTagType) Validate() error {
	return id.value.Validate()
}
