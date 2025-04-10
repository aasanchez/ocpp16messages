package types

// IdTagType represents a validated authorization tag used in OCPP 1.6J.
// It provides methods for string access and validation.
type IdTagType interface {
	String() string
	Validate() error
}

// idTagImpl is the unexported concrete implementation of IdTagType.
// It wraps a validated IdTokenType internally.
type idTagImpl struct {
	value IdTokenType
}

// IdTag constructs a new IdTagType from a raw string.
// It delegates validation to the IdToken constructor.
func IdTag(s string) (IdTagType, error) {
	idToken, err := IdToken(s)

	if err != nil {
		return nil, err
	}

	return idTagImpl{value: idToken}, nil
}

// String returns the raw string value of the idTag.
func (id idTagImpl) String() string {
	return id.value.String()
}

// Validate re-validates the underlying IdToken.
func (id idTagImpl) Validate() error {
	return id.value.Validate()
}
