package authorizetypes

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdTagInfoPayload struct {
	Status      string
	ExpiryDate  *string
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

	info := IdTagInfo{
		status:      status,
		expiryDate:  nil,
		parentIdTag: nil,
	}

	if input.ExpiryDate != nil {
		parsedDate, err := st.SetDateTime(*input.ExpiryDate)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to parse expiryDate", err)
		}

		info.expiryDate = &parsedDate
	}

	if input.ParentIdTag != nil {
		ci, err := st.SetCiString20(*input.ParentIdTag)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(st.ErrFmtFieldWrapped, "failed to validate parentIdTag as CiString20", err)
		}

		idTag, _ := SetIdToken(ci)
		info.parentIdTag = &idTag
	}

	return info, nil
}

func (i IdTagInfo) Value() IdTagInfoPayload {
	var expiry *string

	if str := i.expiryDate; str != nil {
		s := str.String()
		expiry = &s
	}

	var parent *string

	if p := i.parentIdTag; p != nil {
		val := p.Value()
		parent = &val
	}

	return IdTagInfoPayload{
		Status:      i.status.Value(),
		ExpiryDate:  expiry,
		ParentIdTag: parent,
	}
}
