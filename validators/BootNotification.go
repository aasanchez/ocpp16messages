package validators

import (
	"fmt"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
)

func ValidateBootNotificationReq(req chargePoint.BootNotificationReq) error {
	if err := req.ChargePointModel.Validate(); err != nil {
		return err
	}
	if err := req.ChargePointVendor.Validate(); err != nil {
		return err
	}
	if req.ChargeBoxSerialNumber != nil {
		if err := req.ChargeBoxSerialNumber.Validate(); err != nil {
			return err
		}
	}
	if req.ChargePointSerialNumber != nil {
		if err := req.ChargePointSerialNumber.Validate(); err != nil {
			return err
		}
	}
	if req.FirmwareVersion != nil {
		if err := req.FirmwareVersion.Validate(); err != nil {
			return err
		}
	}
	if req.Iccid != nil {
		if err := req.Iccid.Validate(); err != nil {
			return err
		}
	}
	if req.Imsi != nil {
		if err := req.Imsi.Validate(); err != nil {
			return err
		}
	}
	if req.MeterSerialNumber != nil {
		if err := req.MeterSerialNumber.Validate(); err != nil {
			return err
		}
	}
	if req.MeterType != nil {
		if err := req.MeterType.Validate(); err != nil {
			return err
		}
	}
	return nil
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
