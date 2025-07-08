// Package authorizetypes provides types for the Authorize message.
package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// IdTagInfoPayload is a struct that contains information about the authorization status of an IdTag.
// See OCPP 1.6J Part 2, Section 4.1.2.2, Table 14 for details.
type IdTagInfoPayload struct {
	Status      string  // Status of the IdTag. (Required)
	ExpiryDate  *string // ExpiryDate is the date and time after which the IdTag is no longer valid. (Optional)
	ParentIdTag *string // ParentIdTag is the parent IdTag. (Optional)
}

type IdTagInfoType struct {
	expiryDate  *sharedtypes.DateTime
	parentIdTag *IdTokenType
	status      AuthorizationStatusType
}

func IdTagInfo(input IdTagInfoPayload) (IdTagInfoType, error) {
	status, err := AuthorizationStatus(input.Status)
	if err != nil {
		return IdTagInfoType{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse status", err)
	}

	info := IdTagInfoType{
		status:      status,
		expiryDate:  nil,
		parentIdTag: nil,
	}

	if input.ExpiryDate != nil {
		parsedDate, err := sharedtypes.SetDateTime(*input.ExpiryDate)
		if err != nil {
			return IdTagInfoType{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to parse expiryDate", err)
		}

		info.expiryDate = &parsedDate
	}

	if input.ParentIdTag != nil {
		ci, err := sharedtypes.SetCiString20Type(*input.ParentIdTag)
		if err != nil {
			return IdTagInfoType{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, "failed to validate parentIdTag as CiString20", err)
		}

		idTag, _ := IdToken(ci)
		info.parentIdTag = &idTag
	}

	return info, nil
}

func (i IdTagInfoType) Value() IdTagInfoPayload {
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
