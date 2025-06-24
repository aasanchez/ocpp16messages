package bootnotification

import (
	"fmt"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestMessage struct {
	ChargeBoxSerialNumber   *sharedtypes.CiString25Type
	ChargePointModel        sharedtypes.CiString20Type
	ChargePointSerialNumber *sharedtypes.CiString25Type
	ChargePointVendor       sharedtypes.CiString20Type
	FirmwareVersion         *sharedtypes.CiString50Type
	Iccid                   *sharedtypes.CiString20Type
	Imsi                    *sharedtypes.CiString20Type
	MeterSerialNumber       *sharedtypes.CiString25Type
	MeterType               *sharedtypes.CiString25Type
}

func Request(input bootnotificationtypes.RequestPayload) (RequestMessage, error) {
	var err error

	chargeBoxSerialNumber, err := sharedtypes.CiString25Optional("ChargeBoxSerialNumber", input.ChargeBoxSerialNumber)
	if err != nil {
		return RequestMessage{}, wrapErr("chargeBoxSerialNumber", err)
	}

	chargePointModel, err := requiredCiString20("ChargePointModel", input.ChargePointModel)
	if err != nil {
		return RequestMessage{}, err
	}

	chargePointSerialNumber, err := sharedtypes.CiString25Optional("ChargePointSerialNumber", input.ChargePointSerialNumber)
	if err != nil {
		return RequestMessage{}, wrapErr("chargePointSerialNumber", err)
	}

	chargePointVendor, err := requiredCiString20("ChargePointVendor", input.ChargePointVendor)
	if err != nil {
		return RequestMessage{}, err
	}

	firmwareVersion, err := sharedtypes.CiString50Optional("FirmwareVersion", input.FirmwareVersion)
	if err != nil {
		return RequestMessage{}, wrapErr("firmwareVersion", err)
	}

	iccid, err := sharedtypes.CiString20Optional("Iccid", input.Iccid)
	if err != nil {
		return RequestMessage{}, wrapErr("iccid", err)
	}

	imsi, err := sharedtypes.CiString20Optional("Imsi", input.Imsi)
	if err != nil {
		return RequestMessage{}, wrapErr("imsi", err)
	}

	meterSerialNumber, err := sharedtypes.CiString25Optional("MeterSerialNumber", input.MeterSerialNumber)
	if err != nil {
		return RequestMessage{}, wrapErr("meterSerialNumber", err)
	}

	meterType, err := sharedtypes.CiString25Optional("MeterType", input.MeterType)
	if err != nil {
		return RequestMessage{}, wrapErr("meterType", err)
	}

	return RequestMessage{
		ChargeBoxSerialNumber:   chargeBoxSerialNumber,
		ChargePointModel:        chargePointModel,
		ChargePointSerialNumber: chargePointSerialNumber,
		ChargePointVendor:       chargePointVendor,
		FirmwareVersion:         firmwareVersion,
		Iccid:                   iccid,
		Imsi:                    imsi,
		MeterSerialNumber:       meterSerialNumber,
		MeterType:               meterType,
	}, nil
}

// --- Helpers ---

func wrapErr(field string, err error) error {
	return fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, field, err)
}

func requiredCiString20(field, value string) (sharedtypes.CiString20Type, error) {
	if value == "" {
		return sharedtypes.CiString20Type{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, field, sharedtypes.ErrEmptyValueNotAllowed)
	}

	cs, err := sharedtypes.CiString20(value)
	if err != nil {
		return sharedtypes.CiString20Type{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, field, err)
	}

	return cs, nil
}
