package validators

import (
	"fmt"
	"reflect"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
)

func ValidateBootNotificationReq(req chargePoint.BootNotificationReq) error {
	if err := validateRequiredFields(req); err != nil {
		return err
	}
	if err := validateOptionalFields(req); err != nil {
		return err
	}
	return nil
}

func validateRequiredFields(req chargePoint.BootNotificationReq) error {
	if err := req.ChargePointModel.Validate(); err != nil {
		return fmt.Errorf("ChargePointModel: %w", err)
	}
	if err := req.ChargePointVendor.Validate(); err != nil {
		return fmt.Errorf("ChargePointVendor: %w", err)
	}
	return nil
}

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

func isNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func ValidateBootNotificationConf(conf chargePoint.BootNotificationConf) error {
	if !conf.Status.IsValid() {
		return fmt.Errorf("invalid BootNotification status: %s", conf.Status)
	}
	if conf.Interval <= 0 {
		return fmt.Errorf("interval must be > 0")
	}
	return nil
}
