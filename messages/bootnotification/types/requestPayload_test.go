package bootnotificationtypes

import (
	"testing"
)

func TestRequestPayload_ValidateValidPayload(t *testing.T) {
	t.Parallel()
	payload := RequestPayload{
		ChargePointModel:  "ACME Model X",
		ChargePointVendor: "ACME Inc.",
	}

	if err := payload.Validate(); err != nil {
		t.Errorf("expected no validation error, got: %v", err)
	}
}

func TestRequestPayload_ValidateInvalidPayload(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargePointModel:  "ACME Model X",
		ChargePointVendor: "", // Simulates "missing"
	}

	err := payload.Validate()
	if err == nil {
		t.Fatal("expected validation error for missing ChargePointVendor, got nil")
	}
}

func TestRequestPayload_ValidateInvalidPayload2(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{
		ChargePointModel:  "",
		ChargePointVendor: "Vendor",
	}

	err := payload.Validate()
	if err == nil {
		t.Fatal("expected validation error for missing ChargePointVendor, got nil")
	}
}
