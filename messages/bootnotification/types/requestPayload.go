package bootnotificationtypes

import (
	"errors"
	"fmt"
	"reflect"
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

	optionalFields := map[string]interface{}{
		"ChargeBoxSerialNumber":   r.ChargeBoxSerialNumber,
		"ChargePointSerialNumber": r.ChargePointSerialNumber,
		"FirmwareVersion":         r.FirmwareVersion,
		"Iccid":                   r.Iccid,
		"Imsi":                    r.Imsi,
		"MeterSerialNumber":       r.MeterSerialNumber,
		"MeterType":               r.MeterType,
	}

	for name, val := range optionalFields {
		if val == nil {
			continue
		}

		typ := reflect.TypeOf(val)
		if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.String {
			return fmt.Errorf("request payload: field %s: %w (%s)", name, errInvalidOptionalField, typ.String())
		}
	}

	return nil
}
