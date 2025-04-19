package bootnotification

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/types"
)

// RequestMessage represents the OCPP 1.6J BootNotification.req message.
//
// This message is sent by a Charge Point to the Central System after it boots or reboots.
// It contains identifying information about the Charge Point and its configuration.
//
// Specification Reference:
// - OCPP 1.6J, Section 4.2: BootNotification.req
type RequestMessage struct {
	ChargeBoxSerialNumber   *types.CiString25Type // Optional
	ChargePointModel        types.CiString20Type  // Required
	ChargePointSerialNumber *types.CiString25Type // Optional
	ChargePointVendor       types.CiString20Type  // Required
	FirmwareVersion         *types.CiString50Type // Optional
	Iccid                   *types.CiString20Type // Optional
	Imsi                    *types.CiString20Type // Optional
	MeterSerialNumber       *types.CiString25Type // Optional
	MeterType               *types.CiString25Type // Optional
}

// BootNotificationInput defines the raw string input for creating a BootNotification RequestMessage.
// All values are plain strings and will be validated and converted to proper types during construction.
//
// Fields marked as "Optional" may be left as an empty string.
type BootNotificationInput struct {
	ChargePointModel        string // Required
	ChargePointVendor       string // Required
	ChargeBoxSerialNumber   string // Optional
	ChargePointSerialNumber string // Optional
	FirmwareVersion         string // Optional
	Iccid                   string // Optional
	Imsi                    string // Optional
	MeterSerialNumber       string // Optional
	MeterType               string // Optional
}

func Request(input BootNotificationInput) (RequestMessage, error) {
	model, err := types.CiString20(input.ChargePointModel)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("invalid ChargePointModel: %w", err)
	}
	vendor, err := types.CiString20(input.ChargePointVendor)
	if err != nil {
		return RequestMessage{}, fmt.Errorf("invalid ChargePointVendor: %w", err)
	}

	msg := RequestMessage{
		ChargePointModel:  model,
		ChargePointVendor: vendor,
	}

	if input.ChargeBoxSerialNumber != "" {
		v, err := types.CiString25(input.ChargeBoxSerialNumber)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid ChargeBoxSerialNumber: %w", err)
		}
		msg.ChargeBoxSerialNumber = &v
	}
	if input.ChargePointSerialNumber != "" {
		v, err := types.CiString25(input.ChargePointSerialNumber)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid ChargePointSerialNumber: %w", err)
		}
		msg.ChargePointSerialNumber = &v
	}
	if input.FirmwareVersion != "" {
		v, err := types.CiString50(input.FirmwareVersion)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid FirmwareVersion: %w", err)
		}
		msg.FirmwareVersion = &v
	}
	if input.Iccid != "" {
		v, err := types.CiString20(input.Iccid)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid Iccid: %w", err)
		}
		msg.Iccid = &v
	}
	if input.Imsi != "" {
		v, err := types.CiString20(input.Imsi)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid Imsi: %w", err)
		}
		msg.Imsi = &v
	}
	if input.MeterSerialNumber != "" {
		v, err := types.CiString25(input.MeterSerialNumber)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid MeterSerialNumber: %w", err)
		}
		msg.MeterSerialNumber = &v
	}
	if input.MeterType != "" {
		v, err := types.CiString25(input.MeterType)
		if err != nil {
			return RequestMessage{}, fmt.Errorf("invalid MeterType: %w", err)
		}
		msg.MeterType = &v
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
	s := fmt.Sprintf(
		"BootNotification.req(chargePointModel:%s, chargePointVendor:%s",
		m.ChargePointModel, m.ChargePointVendor,
	)

	if m.ChargeBoxSerialNumber != nil {
		s += fmt.Sprintf(", chargeBoxSerialNumber:%s", *m.ChargeBoxSerialNumber)
	}
	if m.ChargePointSerialNumber != nil {
		s += fmt.Sprintf(", chargePointSerialNumber:%s", *m.ChargePointSerialNumber)
	}
	if m.FirmwareVersion != nil {
		s += fmt.Sprintf(", firmwareVersion:%s", *m.FirmwareVersion)
	}
	if m.Iccid != nil {
		s += fmt.Sprintf(", iccid:%s", *m.Iccid)
	}
	if m.Imsi != nil {
		s += fmt.Sprintf(", imsi:%s", *m.Imsi)
	}
	if m.MeterSerialNumber != nil {
		s += fmt.Sprintf(", meterSerialNumber:%s", *m.MeterSerialNumber)
	}
	if m.MeterType != nil {
		s += fmt.Sprintf(", meterType:%s", *m.MeterType)
	}

	s += ")"
	return s
}
