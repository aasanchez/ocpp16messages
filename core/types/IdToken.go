package types

// IdToken represents a unique identifier for a user or charging session.
//
// According to the OCPP 1.6J specification, this is defined as a CiString20.
type IdToken struct {
	CiString20
}

// NewIdToken creates a new IdToken and validates its length.
func NewIdToken(value string) (*IdToken, error) {
	cs := CiString20(value)
	if !cs.IsValid() {
		return nil, NewFieldError("idToken", "idToken exceeds 20 characters")
	}
	return &IdToken{cs}, nil
}

// String returns the string value of the IdToken.
func (id *IdToken) String() string {
	if id == nil {
		return ""
	}
	return string(id.CiString20)
}
