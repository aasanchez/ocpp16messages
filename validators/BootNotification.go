// Package validators provides functions to validate OCPP 1.6 messages
// against protocol-defined constraints. This includes checking required fields,
// optional field lengths, enum values, and semantic correctness.
package validators

import (
	"fmt"
	"reflect"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
)

// ValidateBootNotificationReq validates a BootNotificationReq message
// to ensure it conforms to OCPP 1.6J requirements.
//
// It checks:
//   - Required fields (e.g., ChargePointModel, ChargePointVendor) are present and valid
//   - Optional fields, when present, conform to max length constraints
func ValidateBootNotificationReq(req chargePoint.BootNotificationReq) error {
	if err := validateRequiredFields(req); err != nil {
		return err
	}
	if err := validateOptionalFields(req); err != nil {
		return err
	}
	return nil
}

// validateRequiredFields checks required fields in the BootNotificationReq.
// If a required field is missing or exceeds limits, an error is returned.
func validateRequiredFields(req chargePoint.BootNotificationReq) error {
	if err := req.ChargePointModel.Validate(); err != nil {
		return fmt.Errorf("ChargePointModel: %w", err)
	}
	if err := req.ChargePointVendor.Validate(); err != nil {
		return fmt.Errorf("ChargePointVendor: %w", err)
	}
	return nil
}

// validateOptionalFields checks optional fields in BootNotificationReq.
// If a field is present and invalid, it returns an error. Fields are skipped if nil.
func validateOptionalFields(req chargePoint.BootNotificationReq) error {
	validators := []struct {
		name string
		val  interface {
			Validate() error
		}
	}{
		{"ChargeBoxSerialNumber", req.ChargeBoxSerialNumber},
		{"ChargePointSerialNumber", req.ChargePointSerialNumber},
		{"FirmwareVersion", req.FirmwareVersion},
		{"Iccid", req.Iccid},
		{"Imsi", req.Imsi},
		{"MeterSerialNumber", req.MeterSerialNumber},
		{"MeterType", req.MeterType},
	}

	for _, v := range validators {
		if v.val != nil && !isNil(v.val) {
			if err := v.val.Validate(); err != nil {
				return fmt.Errorf("%s: %w", v.name, err)
			}
		}
	}
	return nil
}

// isNil checks whether an interface value is nil or contains a nil pointer.
//
// This accounts for a Go edge case where an interface value can be non-nil
// while still containing a nil pointer (e.g., typed nil).
func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

// ValidateBootNotificationConf validates a BootNotificationConf message
// to ensure that status is one of the allowed enum values and the interval is positive.
//
// Returns an error if the message is invalid.
func ValidateBootNotificationConf(conf chargePoint.BootNotificationConf) error {
	if !conf.Status.IsValid() {
		return fmt.Errorf("invalid BootNotification status: %s", conf.Status)
	}
	if conf.Interval <= 0 {
		return fmt.Errorf("interval must be > 0")
	}
	return nil
}
