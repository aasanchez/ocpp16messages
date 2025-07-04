package bootnotification

import (
	"fmt"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestMessage struct {
	ChargeBoxSerialNumber   sharedtypes.CiString25
	ChargePointModel        sharedtypes.CiString20
	ChargePointSerialNumber sharedtypes.CiString25
	ChargePointVendor       sharedtypes.CiString20
	FirmwareVersion         sharedtypes.CiString50
	Iccid                   sharedtypes.CiString20
	Imsi                    sharedtypes.CiString20
	MeterSerialNumber       sharedtypes.CiString25
	MeterType               sharedtypes.CiString25
}

func Request(input bootnotificationtypes.RequestPayload) (RequestMessage, error) {
	var err error

	chargeBoxSerialNumber, err := sharedtypes.SetCiString25(input.ChargeBoxSerialNumber)
	if err != nil {
		return RequestMessage{}, wrapErr("chargeBoxSerialNumber", err)
	}

	chargePointModel, err := requiredCiString20("ChargePointModel", input.ChargePointModel)
	if err != nil {
		return RequestMessage{}, err
	}

	chargePointSerialNumber, err := sharedtypes.SetCiString25(input.ChargePointSerialNumber)
	if err != nil {
		return RequestMessage{}, wrapErr("chargePointSerialNumber", err)
	}

	chargePointVendor, err := requiredCiString20("ChargePointVendor", input.ChargePointVendor)
	if err != nil {
		return RequestMessage{}, err
	}

	firmwareVersion, err := sharedtypes.SetCiString50(input.FirmwareVersion)
	if err != nil {
		return RequestMessage{}, wrapErr("firmwareVersion", err)
	}

	iccid, err := sharedtypes.SetCiString20(input.Iccid)
	if err != nil {
		return RequestMessage{}, wrapErr("iccid", err)
	}

	imsi, err := sharedtypes.SetCiString20(input.Imsi)
	if err != nil {
		return RequestMessage{}, wrapErr("imsi", err)
	}

	meterSerialNumber, err := sharedtypes.SetCiString25(input.MeterSerialNumber)
	if err != nil {
		return RequestMessage{}, wrapErr("meterSerialNumber", err)
	}

	meterType, err := sharedtypes.SetCiString25(input.MeterType)
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

func requiredCiString20(field, value string) (sharedtypes.CiString20, error) {
	if value == "" {
		return sharedtypes.CiString20{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, field, sharedtypes.ErrEmptyValueNotAllowed)
	}

	cs, err := sharedtypes.SetCiString20(value)
	if err != nil {
		return sharedtypes.CiString20{}, fmt.Errorf(sharedtypes.ErrFmtFieldWrapped, field, err)
	}

	return cs, nil
}
