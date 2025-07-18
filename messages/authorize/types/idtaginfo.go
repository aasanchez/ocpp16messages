package authorizetypes

import (
	"errors"
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

var (
	errParsingExpiredDate = errors.New("failed to parse expiryDate")
	errParseStatus        = errors.New("failed to parse status")
	errParseParentId      = errors.New("failed to validate parentIdTag as CiString20")
)

type IdTagInfo struct {
	expiryDate  *st.DateTime
	parentIdTag *IdToken
	status      AuthorizationStatus
}

func SetIdTagInfo(input IdTagInfoPayload) (IdTagInfo, error) {
	status, err := SetAuthorizationStatus(input.Status)
	if err != nil {
		return IdTagInfo{}, fmt.Errorf(st.ErrorWrapper, errParseStatus, err)
	}

	info := IdTagInfo{
		status:      status,
		expiryDate:  nil,
		parentIdTag: nil,
	}

	if input.ExpiryDate != "" {
		parsedDate, err := st.SetDateTime(input.ExpiryDate)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(
				st.ErrorWrapper,
				errParsingExpiredDate,
				err,
			)
		}

		info.expiryDate = &parsedDate
	}

	if input.ParentIdTag != "" {
		cistring, err := st.SetCiString20Type(input.ParentIdTag)
		if err != nil {
			return IdTagInfo{}, fmt.Errorf(
				st.ErrorWrapper,
				errParseParentId,
				err,
			)
		}

		idTag, _ := SetIdToken(cistring)
		info.parentIdTag = &idTag
	}

	return info, nil
}

func (idTagInfo IdTagInfo) Value() IdTagInfoPayload {

}
