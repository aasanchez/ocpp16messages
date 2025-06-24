package bootnotificationtypes

import (
	"errors"
	"fmt"
)

var (
	errMissingChargePointModel  = errors.New("missing required field: ChargePointModel")
	errMissingChargePointVendor = errors.New("missing required field: ChargePointVendor")
	errInvalidOptionalField     = errors.New("field must be of type *string or nil")
)

type RequestPayload struct {
	ChargeBoxSerialNumber   *string
	ChargePointModel        string
	ChargePointSerialNumber *string
	ChargePointVendor       string
	FirmwareVersion         *string
	Iccid                   *string
	Imsi                    *string
	MeterSerialNumber       *string
	MeterType               *string
}

func (r RequestPayload) Validate() error {
	if r.ChargePointModel == "" {
		return fmt.Errorf("request payload: %w", errMissingChargePointModel)
	}

	if r.ChargePointVendor == "" {
		return fmt.Errorf("request payload: %w", errMissingChargePointVendor)
	}

	return nil
}
