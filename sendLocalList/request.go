package sendLocalList

import (
	"errors"
	"fmt"

	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

// ReqInput represents the raw input data for creating a SendLocalList.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: Version number of the local authorization list.
	ListVersion int
	// Optional: Authorization data entries to update.
	// When empty with Full updateType, the local list is cleared.
	LocalAuthorizationList []slt.AuthorizationDataInput
	// Required: Type of update ("Full" or "Differential").
	UpdateType string
}

// ReqMessage represents an OCPP 1.6 SendLocalList.req message.
type ReqMessage struct {
	ListVersion            st.Integer
	LocalAuthorizationList []slt.AuthorizationData
	UpdateType             slt.UpdateType
}

// reqValidation holds validated fields during construction.
type reqValidation struct {
	listVersion            st.Integer
	localAuthorizationList []slt.AuthorizationData
	updateType             slt.UpdateType
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
			ListVersion:            st.Integer{},
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
) (st.Integer, []error) {
	intVal, err := st.NewIntegerFromInt(listVersion)
	if err != nil {
		return st.Integer{}, append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "ListVersion", err),
		)
	}

	return intVal, errs
}

func validateReqUpdateType(
	updateType string,
	errs []error,
) (slt.UpdateType, []error) {
	updateTypeVal := slt.UpdateType(updateType)

	if !updateTypeVal.IsValid() {
		return "", append(
			errs,
			fmt.Errorf(st.ErrorFieldFormat, "UpdateType", st.ErrInvalidValue),
		)
	}

	return updateTypeVal, errs
}

const authListLenZero = 0

func validateReqAuthorizationList(
	authList []slt.AuthorizationDataInput,
	errs []error,
) ([]slt.AuthorizationData, []error) {
	if len(authList) == authListLenZero {
		return nil, errs
	}

	var validEntries []slt.AuthorizationData

	for i, entry := range authList {
		authData, err := slt.NewAuthorizationData(entry)
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
