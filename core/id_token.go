package core

// IdTag represents a unique identifier for a user or charging session.
//
// According to the OCPP 1.6J specification, this is defined as a CiString20.
type IdTag struct {
	CiString20
}

// NewIdTag creates a new IdTag and validates its length.
func NewIdTag(value string) (*IdTag, error) {
	cs := CiString20(value)
	if !cs.IsValid() {
		return nil, NewFieldError("idTag", "idTag exceeds 20 characters")
	}
	return &IdTag{cs}, nil
}

// String returns the string value of the IdTag.
func (id *IdTag) String() string {
	if id == nil {
		return ""
	}
	return string(id.CiString20)
}
