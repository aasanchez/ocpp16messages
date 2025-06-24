package bootnotification

import (
	"fmt"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestMessage struct {
	ChargeBoxSerialNumber   *sharedtypes.CiString25Type // Optional
	ChargePointModel        sharedtypes.CiString20Type  // Required
	ChargePointSerialNumber *sharedtypes.CiString25Type // Optional
	ChargePointVendor       sharedtypes.CiString20Type  // Required
	FirmwareVersion         *sharedtypes.CiString50Type // Optional
	Iccid                   *sharedtypes.CiString20Type // Optional
	Imsi                    *sharedtypes.CiString20Type // Optional
	MeterSerialNumber       *sharedtypes.CiString25Type // Optional
	MeterType               *sharedtypes.CiString25Type // Optional
}

func Request(input bootnotificationtypes.RequestPayload) (RequestMessage, error) {
	chargeBoxSerialNumber, err := sharedtypes.CiString25_Optional("ChargeBoxSerialNumber", input.ChargeBoxSerialNumber)
	if err != nil {
		return RequestMessage{}, err
	}

	chargePointModel, err := sharedtypes.CiString20(input.ChargePointModel)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("invalid ChargePointModel: %w", err)
	}

	chargePointSerialNumber, err := sharedtypes.CiString25_Optional("ChargePointSerialNumber", input.ChargePointSerialNumber)
	if err != nil {
		return RequestMessage{}, err
	}

	chargePointVendor, err := sharedtypes.CiString20(input.ChargePointVendor)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("invalid ChargePointVendor: %w", err)
	}

	firmwareVersion, err := sharedtypes.CiString50_Optional("FirmwareVersion", input.FirmwareVersion)
	if err != nil {
		return RequestMessage{}, err
	}

	iccid, err := sharedtypes.CiString20_Optional("Iccid", input.Iccid)
	if err != nil {
		return RequestMessage{}, err
	}

	imsi, err := sharedtypes.CiString20_Optional("Imsi", input.Imsi)
	if err != nil {
		return RequestMessage{}, err
	}

	meterSerialNumber, err := sharedtypes.CiString25_Optional("MeterSerialNumber", input.MeterSerialNumber)
	if err != nil {
		return RequestMessage{}, err
	}

	meterType, err := sharedtypes.CiString25_Optional("MeterType", input.MeterType)
	if err != nil {
		return RequestMessage{}, err
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

