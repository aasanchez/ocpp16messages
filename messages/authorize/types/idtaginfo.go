package authorizetypes

// import (
// 	"fmt"

// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// type IdTagInfo struct {
// 	expiryDate  *sharedtypes.DateTime
// 	parentIdTag *IdToken
// 	status      AuthorizationStatus
// }

// func SetIdTagInfo(input IdTagInfoPayload) (IdTagInfo, error) {
// 	status, err := SetAuthorizationStatus(input.Status)
// 	if err != nil {
// 		return IdTagInfo{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped,"failed to parse status",err)
// 	}

// 	info := IdTagInfo{
// 		status:      status,
// 		expiryDate:  nil,
// 		parentIdTag: nil,
// 	}

// 	if input.ExpiryDate != nil {
// 		parsedDate, err := sharedtypes.SetDateTime(*input.ExpiryDate)
// 		if err != nil {
// 			return IdTagInfo{}, fmt.Errorf(
// 				sharedtypes.ErrFmtFieldWrapped,
// 				"failed to parse expiryDate",
// 				err,
// 			)
// 		}

// 		info.expiryDate = &parsedDate
// 	}

// 	if input.ParentIdTag != nil {
// 		cistring, err := sharedtypes.SetCiString20Type(*input.ParentIdTag)
// 		if err != nil {
// 			return IdTagInfo{}, fmt.Errorf(
// 				sharedtypes.ErrFmtFieldWrapped,
// 				"failed to validate parentIdTag as CiString20",
// 				err,
// 			)
// 		}

// 		idTag, _ := SetIdToken(cistring)
// 		info.parentIdTag = &idTag
// 	}

// 	return info, nil
// }

// func (idTagInfo IdTagInfo) Value() IdTagInfoPayload {
// 	var expiry *string

// 	if str := idTagInfo.expiryDate; str != nil {
// 		s := str.String()
// 		expiry = &s
// 	}

// 	var parent *string

// 	if parentTag := idTagInfo.parentIdTag; parentTag != nil {
// 		val := parentTag.Value()
// 		parent = &val
// 	}

// 	return IdTagInfoPayload{
// 		Status:      idTagInfo.status.Value(),
// 		ExpiryDate:  expiry,
// 		ParentIdTag: parent,
// 	}
// }
