package getDiagnostics_test

import (
	"strings"
	"testing"

	gd "github.com/aasanchez/ocpp16messages/getDiagnostics"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	errLocation      = "location"
	errRetries       = "retries"
	errRetryInterval = "retryInterval"
	errStartTime     = "startTime"
	errStopTime      = "stopTime"

	valueZero       = 0
	valueThree      = 3
	valueSixty      = 60
	valueNegative   = -1
	valueExceedsMax = 65536

	validLocationValue  = "https://example.com/diagnostics"
	validStartTimeValue = "2025-01-01T00:00:00Z"
	validStopTimeValue  = "2025-01-02T00:00:00Z"
	invalidTimeValue    = "invalid-datetime"

	retriesNotNil       = "Retries should not be nil"
	retryIntervalNotNil = "RetryInterval should not be nil"
	startTimeNotNil     = "StartTime should not be nil"
	stopTimeNotNil      = "StopTime should not be nil"
)

func strPtr(v string) *string {
	return &v
}

func intPtr(v int) *int {
	return &v
}

func TestReq_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Location.Value() != validLocationValue {
		t.Errorf(st.ErrorMismatch, validLocationValue, req.Location.Value())
	}

	if req.Retries != nil {
		t.Errorf("Retries should be nil, got %v", req.Retries)
	}

	if req.RetryInterval != nil {
		t.Errorf("RetryInterval should be nil, got %v", req.RetryInterval)
	}

	if req.StartTime != nil {
		t.Errorf("StartTime should be nil, got %v", req.StartTime)
	}

	if req.StopTime != nil {
		t.Errorf("StopTime should be nil, got %v", req.StopTime)
	}
}

func TestReq_Valid_WithRetries(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       intPtr(valueThree),
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Retries == nil {
		t.Fatal(retriesNotNil)
	}

	if req.Retries.Value() != valueThree {
		t.Errorf(st.ErrorMismatchValue, valueThree, req.Retries.Value())
	}
}

func TestReq_Valid_WithRetriesZero(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       intPtr(valueZero),
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Retries == nil {
		t.Fatal(retriesNotNil)
	}

	if req.Retries.Value() != valueZero {
		t.Errorf(st.ErrorMismatchValue, valueZero, req.Retries.Value())
	}
}

func TestReq_Valid_WithRetryInterval(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: intPtr(valueSixty),
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.RetryInterval == nil {
		t.Fatal(retryIntervalNotNil)
	}

	if req.RetryInterval.Value() != valueSixty {
		t.Errorf(st.ErrorMismatchValue, valueSixty, req.RetryInterval.Value())
	}
}

func TestReq_Valid_WithStartTime(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     strPtr(validStartTimeValue),
		StopTime:      nil,
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.StartTime == nil {
		t.Fatal(startTimeNotNil)
	}
}

func TestReq_Valid_WithStopTime(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      strPtr(validStopTimeValue),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.StopTime == nil {
		t.Fatal(stopTimeNotNil)
	}
}

func TestReq_Valid_WithAllFields(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       intPtr(valueThree),
		RetryInterval: intPtr(valueSixty),
		StartTime:     strPtr(validStartTimeValue),
		StopTime:      strPtr(validStopTimeValue),
	})
	if err != nil {
		t.Errorf(st.ErrorUnexpectedError, err)
	}

	if req.Location.Value() != validLocationValue {
		t.Errorf(st.ErrorMismatch, validLocationValue, req.Location.Value())
	}

	if req.Retries == nil {
		t.Fatal(retriesNotNil)
	}

	if req.RetryInterval == nil {
		t.Fatal(retryIntervalNotNil)
	}

	if req.StartTime == nil {
		t.Fatal(startTimeNotNil)
	}

	if req.StopTime == nil {
		t.Fatal(stopTimeNotNil)
	}
}

func TestReq_Invalid_EmptyLocation(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      "",
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "empty Location")
	}

	if !strings.Contains(err.Error(), errLocation) {
		t.Errorf(st.ErrorWantContains, err, errLocation)
	}
}

func TestReq_Invalid_NegativeRetries(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       intPtr(valueNegative),
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative Retries")
	}

	if !strings.Contains(err.Error(), errRetries) {
		t.Errorf(st.ErrorWantContains, err, errRetries)
	}
}

func TestReq_Invalid_RetriesExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       intPtr(valueExceedsMax),
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "Retries exceeds max")
	}

	if !strings.Contains(err.Error(), errRetries) {
		t.Errorf(st.ErrorWantContains, err, errRetries)
	}
}

func TestReq_Invalid_NegativeRetryInterval(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: intPtr(valueNegative),
		StartTime:     nil,
		StopTime:      nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "negative RetryInterval")
	}

	if !strings.Contains(err.Error(), errRetryInterval) {
		t.Errorf(st.ErrorWantContains, err, errRetryInterval)
	}
}

func TestReq_Invalid_RetryIntervalExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: intPtr(valueExceedsMax),
		StartTime:     nil,
		StopTime:      nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "RetryInterval exceeds max")
	}

	if !strings.Contains(err.Error(), errRetryInterval) {
		t.Errorf(st.ErrorWantContains, err, errRetryInterval)
	}
}

func TestReq_Invalid_InvalidStartTime(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     strPtr(invalidTimeValue),
		StopTime:      nil,
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid StartTime")
	}

	if !strings.Contains(err.Error(), errStartTime) {
		t.Errorf(st.ErrorWantContains, err, errStartTime)
	}
}

func TestReq_Invalid_InvalidStopTime(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      validLocationValue,
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      strPtr(invalidTimeValue),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "invalid StopTime")
	}

	if !strings.Contains(err.Error(), errStopTime) {
		t.Errorf(st.ErrorWantContains, err, errStopTime)
	}
}

func TestReq_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := gd.Req(gd.ReqInput{
		Location:      "",
		Retries:       intPtr(valueNegative),
		RetryInterval: intPtr(valueNegative),
		StartTime:     strPtr(invalidTimeValue),
		StopTime:      strPtr(invalidTimeValue),
	})
	if err == nil {
		t.Errorf(st.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errLocation) {
		t.Errorf(st.ErrorWantContains, err, errLocation)
	}

	if !strings.Contains(err.Error(), errRetries) {
		t.Errorf(st.ErrorWantContains, err, errRetries)
	}

	if !strings.Contains(err.Error(), errRetryInterval) {
		t.Errorf(st.ErrorWantContains, err, errRetryInterval)
	}

	if !strings.Contains(err.Error(), errStartTime) {
		t.Errorf(st.ErrorWantContains, err, errStartTime)
	}

	if !strings.Contains(err.Error(), errStopTime) {
		t.Errorf(st.ErrorWantContains, err, errStopTime)
	}
}
