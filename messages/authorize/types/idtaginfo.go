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

// Reusable error wrap format to comply with SonarQube rule go:S1192.
const errWrapFormat = "%s: %w"

// IdTagInfo creates a validated IdTagInfoType from raw string inputs.
func IdTagInfo(input IdTagInfoPayload) (IdTagInfoType, error) {
	status, err := AuthorizationStatus(input.Status)
	if err != nil {
		return IdTagInfoType{}, fmt.Errorf(errWrapFormat, "failed to parse status", err)
	}

	info := IdTagInfoType{status: status}

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

		idTag, _ := IdToken(ci) // IdToken is guaranteed to succeed
		info.parentIdTag = &idTag
	}

	return info, nil
}

type IdTagInfoValue struct {
	Status      string
	ExpiryDate  *string
	ParentIdTag *string
}

func (i IdTagInfoType) Value() IdTagInfoValue {
	var expiry *string
	if i.expiryDate != nil {
		str := i.expiryDate.String()
		expiry = &str
	}

	var parent *string
	if i.parentIdTag != nil {
		val := i.parentIdTag.Value()
		parent = &val
	}

	return IdTagInfoValue{
		Status:      i.status.Value(),
		ExpiryDate:  expiry,
		ParentIdTag: parent,
	}
}
