package authorizetypes

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

// IdTagInfoPayload represents the raw input fields for constructing IdTagInfoType.
type IdTagInfoPayload struct {
	Status      string
	ExpiryDate  *string
	ParentIdTag *string
}

// IdTagInfoType holds the validated internal representation of IdTagInfo.
type IdTagInfoType struct {
	expiryDate  *sharedtypes.DateTimeType
	parentIdTag *IdTokenType
	status      AuthorizationStatusType
}

const errWrapFormat = "%s: %w"

// IdTagInfo constructs a validated IdTagInfoType from the raw payload.
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

		idTag, _ := IdToken(ci) // IdToken is guaranteed to succeed
		info.parentIdTag = &idTag
	}

	return info, nil
}

// IdTagInfoValue is the public value object returned from IdTagInfoType.
type IdTagInfoValue struct {
	Status      string
	ExpiryDate  *string
	ParentIdTag *string
}

// Value converts the internal IdTagInfoType into its public representation.
func (i IdTagInfoType) Value() IdTagInfoValue {
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

	return IdTagInfoValue{
		Status:      i.status.Value(),
		ExpiryDate:  expiry,
		ParentIdTag: parent,
	}
}
