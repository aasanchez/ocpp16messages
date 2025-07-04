package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdTagInfoPayload struct {
	Status      string
	ExpiryDate  string
	ParentIdTag *string
}

type IdTagInfo struct {
	expiryDate  *st.DateTime
	parentIdTag *IdToken
	status      AuthorizationStatus
}

func SetIdTagInfo(input IdTagInfoPayload) (IdTagInfo, error) {
	status, err := SetAuthorizationStatus(input.Status)
	if err != nil {
		return IdTagInfo{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to parse status", err)
	}

	var expiryDate *st.DateTime
	if input.ExpiryDate != "" {
		parsed, err := st.SetDateTime(input.ExpiryDate)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to parse expiryDate", err)
		}
		expiryDate = parsed
	}

	var parentIdTag *IdToken
	if input.ParentIdTag != nil {
		ci, err := st.SetCiString20(*input.ParentIdTag)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to validate parentIdTag as CiString20", err)
		}
		tag, _ := SetIdToken(ci)
		parentIdTag = &tag
	}

	return IdTagInfo{
		status:      status,
		expiryDate:  expiryDate,
		parentIdTag: parentIdTag,
	}, nil
}

func (i IdTagInfo) Value() IdTagInfoPayload {
	var expiry string
	if i.expiryDate != nil {
		expiry = i.expiryDate.String()
	}

	var parent *string
	if i.parentIdTag != nil {
		val := i.parentIdTag.Value()
		parent = &val
	}

	return IdTagInfoPayload{
		Status:      i.status.Value(),
		ExpiryDate:  expiry,
		ParentIdTag: parent,
	}
}
