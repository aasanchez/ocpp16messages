package models

import (
	"testing"
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
)

func TestValidateBootNotificationReq_Valid(t *testing.T) {
	req := chargePoint.BootNotificationReq{
		ChargePointModel:  CiString20Type("Model123"),
		ChargePointVendor: CiString20Type("VendorX"),
	}

	if err := ValidateBootNotificationReq(req); err != nil {
		t.Errorf("expected valid BootNotificationReq, got error: %v", err)
	}
}

func TestValidateBootNotificationReq_InvalidModel(t *testing.T) {
	req := chargePoint.BootNotificationReq{
		ChargePointModel:  CiString20Type(""), // invalid: empty
		ChargePointVendor: CiString20Type("VendorX"),
	}

	if err := ValidateBootNotificationReq(req); err == nil {
		t.Error("expected error for empty ChargePointModel, got nil")
	}
}

func TestValidateBootNotificationReq_TooLongVendor(t *testing.T) {
	req := chargePoint.BootNotificationReq{
		ChargePointModel:  CiString20Type("Model123"),
		ChargePointVendor: CiString20Type("ThisVendorNameIsWayTooLongToBeAccepted"),
	}

	if err := ValidateBootNotificationReq(req); err == nil {
		t.Error("expected error for too long ChargePointVendor, got nil")
	}
}

func TestValidateBootNotificationConf_Valid(t *testing.T) {
	conf := chargePoint.BootNotificationConf{
		Status:      enums.RegistrationAccepted,
		CurrentTime: time.Now(),
		Interval:    60,
	}

	if err := ValidateBootNotificationConf(conf); err != nil {
		t.Errorf("expected valid BootNotificationConf, got error: %v", err)
	}
}

func TestValidateBootNotificationConf_InvalidStatus(t *testing.T) {
	conf := chargePoint.BootNotificationConf{
		Status:      enums.RegistrationStatus("Unknown"),
		CurrentTime: time.Now(),
		Interval:    60,
	}

	if err := ValidateBootNotificationConf(conf); err == nil {
		t.Error("expected error for invalid status, got nil")
	}
}

func TestValidateBootNotificationConf_InvalidInterval(t *testing.T) {
	conf := chargePoint.BootNotificationConf{
		Status:      enums.RegistrationAccepted,
		CurrentTime: time.Now(),
		Interval:    0, // invalid
	}

	if err := ValidateBootNotificationConf(conf); err == nil {
		t.Error("expected error for interval <= 0, got nil")
	}
}
