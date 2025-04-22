package bootnotification

import (
	"fmt"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

const errorbase = "ocpp16messages/bootnotification/request"

// RequestMessage represents the OCPP 1.6J BootNotification.req message.
//
// This message is sent by a Charge Point to the Central System after it boots or reboots.
// It contains identifying information about the Charge Point and its configuration.
//
// Specification Reference:
// - OCPP 1.6J, Section 4.2: BootNotification.req
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

// BootNotificationInput defines the raw string input for creating a BootNotification RequestMessage.
// All values are plain strings and will be validated and converted to proper types during construction.
//
// Fields marked as "Optional" may be left as an empty string.
type BootNotificationInput struct {
	ChargeBoxSerialNumber   string // Optional
	ChargePointModel        string // Required
	ChargePointSerialNumber string // Optional
	ChargePointVendor       string // Required
	FirmwareVersion         string // Optional
	Iccid                   string // Optional
	Imsi                    string // Optional
	MeterSerialNumber       string // Optional
	MeterType               string // Optional
}

func setChargePointModel(raw string) (sharedtypes.CiString20Type, error) {
	chargePointModel, err := sharedtypes.CiString20(raw)
	if err != nil {
		return sharedtypes.CiString20Type{}, fmt.Errorf(
			"%s.setChargePointModel: invalid ChargePointModel: %w",
			errorbase,
			err,
		)
	}

	return chargePointModel, nil
}

func setChargePointVendor(raw string) (sharedtypes.CiString20Type, error) {
	chargePointVendor, err := sharedtypes.CiString20(raw)
	if err != nil {
		return sharedtypes.CiString20Type{}, fmt.Errorf(
			"%s.setChargePointVendor ChargePointModel: %w",
			errorbase,
			err,
		)
	}

	return chargePointVendor, nil
}

func setChargeBoxSerialNumber(raw string) (sharedtypes.CiString25Type, error) {
	chargeBoxSerialNumber, err := sharedtypes.CiString25(raw)
	if err != nil {
		return sharedtypes.CiString25Type{}, fmt.Errorf(
			"%s.setChargeBoxSerialNumber ChargePointModel: %w",
			errorbase,
			err,
		)
	}

	return chargeBoxSerialNumber, nil
}

func Request(input BootNotificationInput) (RequestMessage, error) {
	chargePointModel, err := setChargePointModel(input.ChargePointModel)
	if err != nil {
		return RequestMessage{}, err
	}

	chargePointVendor, err := setChargePointVendor(input.ChargePointVendor)
	if err != nil {
		return RequestMessage{}, err
	}

	msg := RequestMessage{
		ChargeBoxSerialNumber:   nil,
		ChargePointModel:        chargePointModel,
		ChargePointSerialNumber: nil,
		ChargePointVendor:       chargePointVendor,
		FirmwareVersion:         nil,
		Iccid:                   nil,
		Imsi:                    nil,
		MeterSerialNumber:       nil,
		MeterType:               nil,
	}

	if input.ChargeBoxSerialNumber != "" {
		chargeBoxSerialNumber, err := setChargeBoxSerialNumber(input.ChargeBoxSerialNumber)
		if err != nil {
			return RequestMessage{}, err
		}

		msg.ChargeBoxSerialNumber = &chargeBoxSerialNumber
	}

	return msg, nil
}

// Validate checks if the required fields in the RequestMessage are valid.
func (m RequestMessage) Validate() error {
	if err := m.ChargePointModel.Validate(); err != nil {
		return fmt.Errorf("invalid ChargePointModel: %w", err)
	}

	if err := m.ChargePointVendor.Validate(); err != nil {
		return fmt.Errorf("invalid ChargePointVendor: %w", err)
	}

	return nil
}

// String returns a human-readable representation of the BootNotification RequestMessage,
// including all required and optional fields (if present).
func (m RequestMessage) String() string {
	str := fmt.Sprintf(
		"BootNotification.req(chargePointModel:%s, chargePointVendor:%s",
		m.ChargePointModel, m.ChargePointVendor,
	)

	if m.ChargeBoxSerialNumber != nil {
		str += fmt.Sprintf(", chargeBoxSerialNumber:%s", *m.ChargeBoxSerialNumber)
	}

	if m.ChargePointSerialNumber != nil {
		str += fmt.Sprintf(", chargePointSerialNumber:%s", *m.ChargePointSerialNumber)
	}

	if m.FirmwareVersion != nil {
		str += fmt.Sprintf(", firmwareVersion:%s", *m.FirmwareVersion)
	}

	if m.Iccid != nil {
		str += fmt.Sprintf(", iccid:%s", *m.Iccid)
	}

	if m.Imsi != nil {
		str += fmt.Sprintf(", imsi:%s", *m.Imsi)
	}

	if m.MeterSerialNumber != nil {
		str += fmt.Sprintf(", meterSerialNumber:%s", *m.MeterSerialNumber)
	}

	if m.MeterType != nil {
		str += fmt.Sprintf(", meterType:%s", *m.MeterType)
	}

	str += ")"

	return str
}
