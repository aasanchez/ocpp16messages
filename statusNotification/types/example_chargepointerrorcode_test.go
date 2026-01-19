package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/statusNotification/types"
)

const (
	labelErrorCode = "ErrorCode:"
	labelIsValid   = "IsValid:"
)

// ExampleChargePointErrorCode demonstrates using ChargePointErrorCode values.
func ExampleChargePointErrorCode() {
	errorCode := types.ErrCodeNoError

	fmt.Println(labelErrorCode, errorCode.String())
	fmt.Println(labelIsValid, errorCode.IsValid())
	// Output:
	// ErrorCode: NoError
	// IsValid: true
}

// ExampleChargePointErrorCode_groundFailure demonstrates a GroundFailure error.
func ExampleChargePointErrorCode_groundFailure() {
	errorCode := types.ErrCodeGroundFailure

	fmt.Println(labelErrorCode, errorCode.String())
	fmt.Println(labelIsValid, errorCode.IsValid())
	// Output:
	// ErrorCode: GroundFailure
	// IsValid: true
}

// ExampleChargePointErrorCode_invalid demonstrates an invalid error code.
func ExampleChargePointErrorCode_invalid() {
	errorCode := types.ChargePointErrorCode("InvalidCode")

	fmt.Println(labelErrorCode, errorCode.String())
	fmt.Println(labelIsValid, errorCode.IsValid())
	// Output:
	// ErrorCode: InvalidCode
	// IsValid: false
}
