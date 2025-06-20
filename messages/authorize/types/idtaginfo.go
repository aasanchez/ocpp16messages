package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type IdTagInfoPayload struct {
	Status      string
	ExpiryDate  *string
	ParentIdTag *string
}

type IdTagInfoType struct {
	expiryDate  *sharedtypes.DateTimeType
	parentIdTag *IdTokenType
	status      AuthorizationStatusType
}

const errWrapFormat = "%s: %w"

func IdTagInfo(input IdTagInfoPayload) (IdTagInfoType, error) {
	status, err := AuthorizationStatus(input.Status)
	if err != nil {
		return IdTagInfoType{}, fmt.Errorf(errWrapFormat, "failed to parse status", err)
	}

	info := IdTagInfoType{
		status:      status,
		expiryDate:  nil,
		parentIdTag: nil,
	}

	if input.ExpiryDate != nil {
		parsedDate, err := sharedtypes.DateTime(*input.ExpiryDate)
		if err != nil {
			return IdTagInfoType{}, fmt.Errorf(errWrapFormat, "failed to parse expiryDate", err)
		}

		info.expiryDate = &parsedDate
	}

	if input.ParentIdTag != nil {
		ci, err := sharedtypes.CiString20(*input.ParentIdTag)
		if err != nil {
			return IdTagInfoType{}, fmt.Errorf(errWrapFormat, "failed to validate parentIdTag as CiString20", err)
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
