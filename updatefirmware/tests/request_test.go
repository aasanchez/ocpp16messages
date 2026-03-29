package updatefirmware_test

import (
	"strings"
	"testing"

	uf "github.com/aasanchez/ocpp16messages/updatefirmware"
	types "github.com/aasanchez/ocpp16types"
)

const (
	errLocation      = "location"
	errRetrieveDate  = "retrieveDate"
	errRetries       = "retries"
	errRetryInterval = "retryInterval"

	valueZero       = 0
	valueThree      = 3
	valueSixty      = 60
	valueNegative   = -1
	valueExceedsMax = 65536

	validLocationValue     = "https://example.com/firmware/v1.2.3.bin"
	validRetrieveDateValue = "2025-01-15T10:00:00Z"
	invalidDateTimeValue   = "invalid-datetime"

	retriesNotNil       = "Retries should not be nil"
	retryIntervalNotNil = "RetryInterval should not be nil"
)

func intPtr(v int) *int {
	return &v
}

func TestReq_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	req, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       nil,
		RetryInterval: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Location.Value() != validLocationValue {
		t.Errorf(types.ErrorMismatch, validLocationValue, req.Location.Value())
	}

	if req.Retries != nil {
		t.Errorf("Retries should be nil, got %v", req.Retries)
	}

	if req.RetryInterval != nil {
		t.Errorf("RetryInterval should be nil, got %v", req.RetryInterval)
	}
}

func TestReq_Valid_WithRetries(t *testing.T) {
	t.Parallel()

	req, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       intPtr(valueThree),
		RetryInterval: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Retries == nil {
		t.Fatal(retriesNotNil)
	}

	if req.Retries.Value() != valueThree {
		t.Errorf(types.ErrorMismatchValue, valueThree, req.Retries.Value())
	}
}

func TestReq_Valid_WithRetriesZero(t *testing.T) {
	t.Parallel()

	req, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       intPtr(valueZero),
		RetryInterval: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Retries == nil {
		t.Fatal(retriesNotNil)
	}

	if req.Retries.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.Retries.Value())
	}
}

func TestReq_Valid_WithRetryInterval(t *testing.T) {
	t.Parallel()

	req, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       nil,
		RetryInterval: intPtr(valueSixty),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RetryInterval == nil {
		t.Fatal(retryIntervalNotNil)
	}

	if req.RetryInterval.Value() != valueSixty {
		t.Errorf(
			types.ErrorMismatchValue, valueSixty, req.RetryInterval.Value(),
		)
	}
}

func TestReq_Valid_WithAllFields(t *testing.T) {
	t.Parallel()

	req, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       intPtr(valueThree),
		RetryInterval: intPtr(valueSixty),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Location.Value() != validLocationValue {
		t.Errorf(types.ErrorMismatch, validLocationValue, req.Location.Value())
	}

	if req.Retries == nil {
		t.Fatal(retriesNotNil)
	}

	if req.RetryInterval == nil {
		t.Fatal(retryIntervalNotNil)
	}
}

func TestReq_Invalid_EmptyLocation(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      "",
		RetrieveDate:  validRetrieveDateValue,
		Retries:       nil,
		RetryInterval: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty Location")
	}

	if !strings.Contains(err.Error(), errLocation) {
		t.Errorf(types.ErrorWantContains, err, errLocation)
	}
}

func TestReq_Invalid_EmptyRetrieveDate(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  "",
		Retries:       nil,
		RetryInterval: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty RetrieveDate")
	}

	if !strings.Contains(err.Error(), errRetrieveDate) {
		t.Errorf(types.ErrorWantContains, err, errRetrieveDate)
	}
}

func TestReq_Invalid_InvalidRetrieveDate(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  invalidDateTimeValue,
		Retries:       nil,
		RetryInterval: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid RetrieveDate")
	}

	if !strings.Contains(err.Error(), errRetrieveDate) {
		t.Errorf(types.ErrorWantContains, err, errRetrieveDate)
	}
}

func TestReq_Invalid_NegativeRetries(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       intPtr(valueNegative),
		RetryInterval: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative Retries")
	}

	if !strings.Contains(err.Error(), errRetries) {
		t.Errorf(types.ErrorWantContains, err, errRetries)
	}
}

func TestReq_Invalid_RetriesExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       intPtr(valueExceedsMax),
		RetryInterval: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Retries exceeds max")
	}

	if !strings.Contains(err.Error(), errRetries) {
		t.Errorf(types.ErrorWantContains, err, errRetries)
	}
}

func TestReq_Invalid_NegativeRetryInterval(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       nil,
		RetryInterval: intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative RetryInterval")
	}

	if !strings.Contains(err.Error(), errRetryInterval) {
		t.Errorf(types.ErrorWantContains, err, errRetryInterval)
	}
}

func TestReq_Invalid_RetryIntervalExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      validLocationValue,
		RetrieveDate:  validRetrieveDateValue,
		Retries:       nil,
		RetryInterval: intPtr(valueExceedsMax),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "RetryInterval exceeds max")
	}

	if !strings.Contains(err.Error(), errRetryInterval) {
		t.Errorf(types.ErrorWantContains, err, errRetryInterval)
	}
}

func TestReq_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := uf.Req(uf.ReqInput{
		Location:      "",
		RetrieveDate:  invalidDateTimeValue,
		Retries:       intPtr(valueNegative),
		RetryInterval: intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errLocation) {
		t.Errorf(types.ErrorWantContains, err, errLocation)
	}

	if !strings.Contains(err.Error(), errRetrieveDate) {
		t.Errorf(types.ErrorWantContains, err, errRetrieveDate)
	}

	if !strings.Contains(err.Error(), errRetries) {
		t.Errorf(types.ErrorWantContains, err, errRetries)
	}

	if !strings.Contains(err.Error(), errRetryInterval) {
		t.Errorf(types.ErrorWantContains, err, errRetryInterval)
	}
}
