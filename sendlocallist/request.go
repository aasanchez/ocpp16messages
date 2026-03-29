package sendlocallist

import (
	"errors"
	"fmt"

	types "github.com/aasanchez/ocpp16types"
)

// ReqInput represents the raw input data for creating a SendLocalList.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: Version number of the local authorization list.
	ListVersion int
	// Optional: Authorization data entries to update.
	// When empty with Full updateType, the local list is cleared.
	LocalAuthorizationList []types.AuthorizationDataInput
	// Required: Type of update ("Full" or "Differential").
	UpdateType string
}

// ReqMessage represents an OCPP 1.6 SendLocalList.req message.
type ReqMessage struct {
	ListVersion            types.Integer
	LocalAuthorizationList []types.AuthorizationData
	UpdateType             types.UpdateType
}

// reqValidation holds validated fields during construction.
type reqValidation struct {
	listVersion            types.Integer
	localAuthorizationList []types.AuthorizationData
	updateType             types.UpdateType
}

// Req creates a SendLocalList.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ListVersion is negative
//   - UpdateType is not a valid value
//   - Any AuthorizationData entry is invalid
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if errs != nil {
		return ReqMessage{
			ListVersion:            types.Integer{},
			LocalAuthorizationList: nil,
			UpdateType:             "",
		}, errors.Join(errs...)
	}

	return ReqMessage{
		ListVersion:            validated.listVersion,
		LocalAuthorizationList: validated.localAuthorizationList,
		UpdateType:             validated.updateType,
	}, nil
}

func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	validated.listVersion, errs = validateReqListVersion(
		input.ListVersion,
		errs,
	)

	validated.updateType, errs = validateReqUpdateType(
		input.UpdateType,
		errs,
	)

	validated.localAuthorizationList, errs = validateReqAuthorizationList(
		input.LocalAuthorizationList,
		errs,
	)

	return validated, errs
}

func validateReqListVersion(
	listVersion int,
	errs []error,
) (types.Integer, []error) {
	intVal, err := types.NewInteger(listVersion)
	if err != nil {
		return types.Integer{}, append(
			errs,
			fmt.Errorf(types.ErrorFieldFormat, "ListVersion", err),
		)
	}

	return intVal, errs
}

func validateReqUpdateType(
	updateType string,
	errs []error,
) (types.UpdateType, []error) {
	updateTypeVal := types.UpdateType(updateType)

	if !updateTypeVal.IsValid() {
		return "", append(
			errs,
			fmt.Errorf(types.ErrorFieldFormat, "UpdateType", types.ErrInvalidValue),
		)
	}

	return updateTypeVal, errs
}

const authListLenZero = 0

func validateReqAuthorizationList(
	authList []types.AuthorizationDataInput,
	errs []error,
) ([]types.AuthorizationData, []error) {
	if authList == nil {
		return nil, errs
	}

	if len(authList) == authListLenZero {
		return []types.AuthorizationData{}, errs
	}

	var validEntries []types.AuthorizationData

	for i, entry := range authList {
		authData, err := types.NewAuthorizationData(entry)
		if err != nil {
			errs = append(errs, fmt.Errorf(
				"localAuthorizationList[%d]: %w",
				i,
				err,
			))
		} else {
			validEntries = append(validEntries, authData)
		}
	}

	return validEntries, errs
}
