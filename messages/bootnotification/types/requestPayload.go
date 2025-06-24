package bootnotificationtypes

import (
	"fmt"
	"reflect"
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
		return fmt.Errorf("missing required field: ChargePointModel")
	}
	if r.ChargePointVendor == "" {
		return fmt.Errorf("missing required field: ChargePointVendor")
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
			continue // OK
		}
		typ := reflect.TypeOf(val)
		if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.String {
			return fmt.Errorf("field %s must be of type *string or nil, got %s", name, typ.String())
		}
	}

	return nil
}
