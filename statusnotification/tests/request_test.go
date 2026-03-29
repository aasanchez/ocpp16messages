package statusnotification_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16messages/statusnotification"
	types "github.com/aasanchez/ocpp16types"
)

const (
	testConnectorIdOne   = 1
	testConnectorIdZero  = 0
	testValueNegativeOne = -1
	testStatusAvailable  = "Available"
	testStatusCharging   = "Charging"
	testErrorCodeNoError = "NoError"
	errFieldConnectorId  = "connectorId"
	errFieldErrorCode    = "errorCode"
	errFieldStatus       = "status"
	errFieldInfo         = "info"
	errFieldTimestamp    = "timestamp"
	errWantInvalidValue  = "Req() error = %v, want ErrInvalidValue"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != testConnectorIdOne {
		t.Errorf(
			types.ErrorMismatch, testConnectorIdOne, req.ConnectorId.Value(),
		)
	}

	if req.ErrorCode.String() != testErrorCodeNoError {
		t.Errorf(
			types.ErrorMismatch, testErrorCodeNoError, req.ErrorCode.String(),
		)
	}

	if req.Status.String() != testStatusAvailable {
		t.Errorf(types.ErrorMismatch, testStatusAvailable, req.Status.String())
	}
}

func TestReq_ValidConnectorIdZero(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdZero,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorId.Value() != testConnectorIdZero {
		t.Errorf(
			types.ErrorMismatch, testConnectorIdZero, req.ConnectorId.Value(),
		)
	}
}

func TestReq_ValidCharging(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusCharging,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != testStatusCharging {
		t.Errorf(types.ErrorMismatch, testStatusCharging, req.Status.String())
	}
}

func TestReq_ValidWithInfo(t *testing.T) {
	t.Parallel()

	info := "Ground fault detected"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       "GroundFailure",
		Status:          "Faulted",
		Info:            &info,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Info == nil {
		t.Error("Req() Info = nil, want non-nil")
	}

	if req.Info.Value() != info {
		t.Errorf(types.ErrorMismatch, info, req.Info.Value())
	}
}

func TestReq_ValidWithTimestamp(t *testing.T) {
	t.Parallel()

	timestamp := "2025-01-15T10:30:00Z"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       &timestamp,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Timestamp == nil {
		t.Error("Req() Timestamp = nil, want non-nil")
	}
}

func TestReq_ValidWithVendorFields(t *testing.T) {
	t.Parallel()

	vendorId := "VendorX"
	vendorErrorCode := "V001"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        &vendorId,
		VendorErrorCode: &vendorErrorCode,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.VendorId == nil {
		t.Error("Req() VendorId = nil, want non-nil")
	}

	if req.VendorErrorCode == nil {
		t.Error("Req() VendorErrorCode = nil, want non-nil")
	}

	if req.VendorId.Value() != vendorId {
		t.Errorf(types.ErrorMismatch, vendorId, req.VendorId.Value())
	}
}

func TestReq_ConnectorIdNegative(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testValueNegativeOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative connectorId")
	}

	if !strings.Contains(err.Error(), errFieldConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorId)
	}
}

func TestReq_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          "InvalidStatus",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}

	if !strings.Contains(err.Error(), errFieldStatus) {
		t.Errorf(types.ErrorWantContains, err, errFieldStatus)
	}
}

func TestReq_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          "",
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}
}

func TestReq_InvalidErrorCode(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       "InvalidCode",
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid errorCode")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}

	if !strings.Contains(err.Error(), errFieldErrorCode) {
		t.Errorf(types.ErrorWantContains, err, errFieldErrorCode)
	}
}

func TestReq_EmptyErrorCode(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       "",
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty errorCode")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}
}

func TestReq_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	invalidTimestamp := "not-a-timestamp"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       &invalidTimestamp,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_InfoTooLong(t *testing.T) {
	t.Parallel()

	// 51 chars, max is 50
	longInfo := "This info string is way too long for the OCPP spec!"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            &longInfo,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for info too long")
	}

	if !strings.Contains(err.Error(), errFieldInfo) {
		t.Errorf(types.ErrorWantContains, err, errFieldInfo)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidTimestamp := "invalid"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testValueNegativeOne,
		ErrorCode:       "InvalidCode",
		Status:          "InvalidStatus",
		Info:            nil,
		Timestamp:       &invalidTimestamp,
		VendorId:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConnectorId) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorId)
	}

	if !strings.Contains(errStr, errFieldErrorCode) {
		t.Errorf(types.ErrorWantContains, err, errFieldErrorCode)
	}

	if !strings.Contains(errStr, errFieldStatus) {
		t.Errorf(types.ErrorWantContains, err, errFieldStatus)
	}

	if !strings.Contains(errStr, errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_AllStatuses(t *testing.T) {
	t.Parallel()

	statuses := []string{
		"Available",
		"Preparing",
		"Charging",
		"SuspendedEV",
		"SuspendedEVSE",
		"Finishing",
		"Reserved",
		"Unavailable",
		"Faulted",
	}

	for _, status := range statuses {
		req, err := statusnotification.Req(statusnotification.ReqInput{
			ConnectorId:     testConnectorIdOne,
			ErrorCode:       testErrorCodeNoError,
			Status:          status,
			Info:            nil,
			Timestamp:       nil,
			VendorId:        nil,
			VendorErrorCode: nil,
		})
		if err != nil {
			t.Errorf("Req() error = %v for status %s", err, status)
		}

		if req.Status.String() != status {
			t.Errorf(types.ErrorMismatch, status, req.Status.String())
		}
	}
}

func TestReq_AllErrorCodes(t *testing.T) {
	t.Parallel()

	errorCodes := []string{
		"ConnectorLockFailure",
		"EVCommunicationError",
		"GroundFailure",
		"HighTemperature",
		"InternalError",
		"LocalListConflict",
		"NoError",
		"OtherError",
		"OverCurrentFailure",
		"OverVoltage",
		"PowerMeterFailure",
		"PowerSwitchFailure",
		"ReaderFailure",
		"ResetFailure",
		"UnderVoltage",
		"WeakSignal",
	}

	for _, errorCode := range errorCodes {
		req, err := statusnotification.Req(statusnotification.ReqInput{
			ConnectorId:     testConnectorIdOne,
			ErrorCode:       errorCode,
			Status:          testStatusAvailable,
			Info:            nil,
			Timestamp:       nil,
			VendorId:        nil,
			VendorErrorCode: nil,
		})
		if err != nil {
			t.Errorf("Req() error = %v for errorCode %s", err, errorCode)
		}

		if req.ErrorCode.String() != errorCode {
			t.Errorf(types.ErrorMismatch, errorCode, req.ErrorCode.String())
		}
	}
}

func TestReq_InvalidVendorId(t *testing.T) {
	t.Parallel()

	// VendorId must be valid CiString255, test with invalid ASCII
	invalidVendorId := "Vendor\x00Id"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        &invalidVendorId,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid vendorId")
	}

	if !strings.Contains(err.Error(), "vendorId") {
		t.Errorf(types.ErrorWantContains, err, "vendorId")
	}
}

func TestReq_InvalidVendorErrorCode(t *testing.T) {
	t.Parallel()

	// VendorErrorCode must be valid CiString50, test with invalid ASCII
	invalidVendorErrorCode := "Error\x00Code"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorId:     testConnectorIdOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorId:        nil,
		VendorErrorCode: &invalidVendorErrorCode,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid vendorErrorCode")
	}

	if !strings.Contains(err.Error(), "vendorErrorCode") {
		t.Errorf(types.ErrorWantContains, err, "vendorErrorCode")
	}
}
