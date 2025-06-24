package bootnotification

import (
	"errors"
	"fmt"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

var ErrEmptyValueNotAllowed = errors.New("value must not be empty")

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
	chargeBoxSerialNumber, err := sharedtypes.CiString25Optional("ChargeBoxSerialNumber", input.ChargeBoxSerialNumber)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("chargeBoxSerialNumber: %w", err)
	}

	if input.ChargePointModel == "" {
		return RequestMessage{}, fmt.Errorf("chargePointModel: %w", ErrEmptyValueNotAllowed)
	}

	chargePointModel, err := sharedtypes.CiString20(input.ChargePointModel)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("chargePointModel: %w", err)
	}

	chargePointSerialNumber, err := sharedtypes.CiString25Optional("ChargePointSerialNumber", input.ChargePointSerialNumber)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("chargePointSerialNumber: %w", err)
	}

	if input.ChargePointVendor == "" {
		return RequestMessage{}, fmt.Errorf("ChargePointVendor: %w", ErrEmptyValueNotAllowed)
	}

	chargePointVendor, err := sharedtypes.CiString20(input.ChargePointVendor)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("chargePointVendor: %w", err)
	}

	firmwareVersion, err := sharedtypes.CiString50Optional("FirmwareVersion", input.FirmwareVersion)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("firmwareVersion: %w", err)
	}

	iccid, err := sharedtypes.CiString20Optional("Iccid", input.Iccid)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("iccid: %w", err)
	}

	imsi, err := sharedtypes.CiString20Optional("Imsi", input.Imsi)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("imsi: %w", err)
	}

	meterSerialNumber, err := sharedtypes.CiString25Optional("MeterSerialNumber", input.MeterSerialNumber)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("meterSerialNumber: %w", err)
	}

	meterType, err := sharedtypes.CiString25Optional("MeterType", input.MeterType)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("meterType: %w", err)
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
