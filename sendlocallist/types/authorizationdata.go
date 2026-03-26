package types

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/types"
)

// AuthorizationDataInput represents the raw input data for creating an
// AuthorizationData entry in the local authorization list.
type AuthorizationDataInput struct {
	// Required: The identifier to be authorized.
	IdTag string
	// Optional: Authorization status information for this idTag.
	// When omitted in a Differential update, the entry is deleted.
	IdTagInfo *IdTagInfoInput
}

// IdTagInfoInput represents the raw input data for IdTagInfo.
type IdTagInfoInput struct {
	// Required: Authorization status.
	Status string
	// Optional: Expiry date in RFC3339 format.
	ExpiryDate *string
	// Optional: Parent identifier tag.
	ParentIdTag *string
}

// AuthorizationData represents an entry in the local authorization list.
type AuthorizationData struct {
	IdTag     st.IdToken
	IdTagInfo *st.IdTagInfo
}

// NewAuthorizationData creates a new AuthorizationData from the given input.
// It validates all fields and returns an error if:
//   - IdTag is empty or invalid
//   - IdTagInfo.Status is invalid (when IdTagInfo is provided)
func NewAuthorizationData(
	input AuthorizationDataInput,
) (AuthorizationData, error) {
	ciString, err := st.NewCiString20Type(input.IdTag)
	if err != nil {
		return AuthorizationData{}, fmt.Errorf(
			"NewAuthorizationData: "+st.ErrorFieldFormat,
			"IdTag",
			err,
		)
	}

	idToken := st.NewIdToken(ciString)

	if input.IdTagInfo == nil {
		return AuthorizationData{
			IdTag:     idToken,
			IdTagInfo: nil,
		}, nil
	}

	idTagInfo, err := buildIdTagInfo(*input.IdTagInfo)
	if err != nil {
		return AuthorizationData{}, fmt.Errorf(
			"NewAuthorizationData: "+st.ErrorFieldFormat,
			"IdTagInfo",
			err,
		)
	}

	return AuthorizationData{
		IdTag:     idToken,
		IdTagInfo: &idTagInfo,
	}, nil
}

func buildIdTagInfo(input IdTagInfoInput) (st.IdTagInfo, error) {
	status := st.AuthorizationStatus(input.Status)

	idTagInfo, err := st.NewIdTagInfo(status)
	if err != nil {
		return st.IdTagInfo{}, fmt.Errorf("buildIdTagInfo: %w", err)
	}

	if input.ExpiryDate != nil {
		expiryDate, err := st.NewDateTime(*input.ExpiryDate)
		if err != nil {
			return st.IdTagInfo{}, fmt.Errorf(
				st.ErrorFieldFormat,
				"ExpiryDate",
				err,
			)
		}

		idTagInfo = idTagInfo.WithExpiryDate(expiryDate)
	}

	if input.ParentIdTag != nil {
		ciString, err := st.NewCiString20Type(*input.ParentIdTag)
		if err != nil {
			return st.IdTagInfo{}, fmt.Errorf(
				st.ErrorFieldFormat,
				"ParentIdTag",
				err,
			)
		}

		parentIdToken := st.NewIdToken(ciString)
		idTagInfo = idTagInfo.WithParentIdTag(parentIdToken)
	}

	return idTagInfo, nil
}
