package types

// IdTag represents a validated authorization tag used in OCPP 1.6J.
type IdTag struct {
	value IdTokenType
}

// IdTag returns a validated IdTag from a raw string.
func NewIdTag(s string) (IdTag, error) {
	idToken, err := IdToken(s)
	if err != nil {
		return IdTag{}, err
	}

	return IdTag{value: idToken}, nil
}

// String returns the raw string.
func (id IdTag) String() string {
	return id.value.String()
}

// Validate re-validates the wrapped token.
func (id IdTag) Validate() error {
	return id.value.Validate()
}
