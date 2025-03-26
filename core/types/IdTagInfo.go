package types

type IdTagInfo struct {
	ExpiryDate  *string `json:"expiryDate,omitempty"`
	ParentIdTag *string `json:"parentIdTag,omitempty"`
	Status      string  `json:"status"`
}

func NewIdTagInfo(value string) (*IdToken, error) {
	cs := CiString20(value)
	if !cs.IsValid() {
		return nil, NewFieldError("idToken", "idToken exceeds 20 characters")
	}
	return &IdToken{cs}, nil
}

// func (id *IdToken) String() string {
// 	if id == nil {
// 		return ""
// 	}
// 	return string(id.CiString20)
// }
