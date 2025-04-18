package bootnotification

// import (
// 	"fmt"

// 	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
// )

// type RequestMessage struct {
// 	ChargeBoxSerialNumber   *sharedtypes.CiString25Type // Optional
// 	ChargePointModel        sharedtypes.CiString20Type  // Required
// 	ChargePointSerialNumber *sharedtypes.CiString25Type // Optional
// 	ChargePointVendor       sharedtypes.CiString20Type  // Required
// 	FirmwareVersion         *sharedtypes.CiString50Type // Optional
// 	Iccid                   *sharedtypes.CiString20Type // Optional
// 	Imsi                    *sharedtypes.CiString20Type // Optional
// 	MeterSerialNumber       *sharedtypes.CiString25Type // Optional
// 	MeterType               *sharedtypes.CiString25Type // Optional
// }

// func Request(input bootNotificationReqInput) (RequestMessage, error) {
// 	chargePointModel, err := setChargePointModel(input.ChargePointModel)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	chargePointVendor, err := setChargePointVendor(input.ChargePointVendor)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	chargeBoxSerialNumber, err := setOptional(input.ChargeBoxSerialNumber, setChargeBoxSerialNumber)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	chargePointSerialNumber, err := setOptional(input.ChargePointSerialNumber, setChargePointSerialNumber)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	firmwareVersion, err := setOptional(input.FirmwareVersion, setFirmwareVersion)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	iccid, err := setOptional(input.Iccid, setIccid)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	imsi, err := setOptional(input.Imsi, setImsi)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	meterSerialNumber, err := setOptional(input.MeterSerialNumber, setMeterSerialNumber)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	meterType, err := setOptional(input.MeterType, setMeterType)
// 	if err != nil {
// 		return RequestMessage{}, err
// 	}

// 	msg := RequestMessage{
// 		ChargeBoxSerialNumber:   chargeBoxSerialNumber,
// 		ChargePointModel:        chargePointModel,
// 		ChargePointSerialNumber: chargePointSerialNumber,
// 		ChargePointVendor:       chargePointVendor,
// 		FirmwareVersion:         firmwareVersion,
// 		Iccid:                   iccid,
// 		Imsi:                    imsi,
// 		MeterSerialNumber:       meterSerialNumber,
// 		MeterType:               meterType,
// 	}

// 	return msg, nil
// }

// func (m RequestMessage) Validate() error {
// 	if err := m.ChargePointModel.Validate(); err != nil {
// 		return fmt.Errorf("invalid ChargePointModel: %w", err)
// 	}

// 	if err := m.ChargePointVendor.Validate(); err != nil {
// 		return fmt.Errorf("invalid ChargePointVendor: %w", err)
// 	}

// 	return nil
// }

// func (m RequestMessage) String() string {
// 	str := fmt.Sprintf(
// 		"BootNotification.req(chargePointModel:%s, chargePointVendor:%s",
// 		m.ChargePointModel, m.ChargePointVendor,
// 	)

// 	if m.ChargeBoxSerialNumber != nil {
// 		str += fmt.Sprintf(", chargeBoxSerialNumber:%s", *m.ChargeBoxSerialNumber)
// 	}

// 	if m.ChargePointSerialNumber != nil {
// 		str += fmt.Sprintf(", chargePointSerialNumber:%s", *m.ChargePointSerialNumber)
// 	}

// 	if m.FirmwareVersion != nil {
// 		str += fmt.Sprintf(", firmwareVersion:%s", *m.FirmwareVersion)
// 	}

// 	if m.Iccid != nil {
// 		str += fmt.Sprintf(", iccid:%s", *m.Iccid)
// 	}

// 	if m.Imsi != nil {
// 		str += fmt.Sprintf(", imsi:%s", *m.Imsi)
// 	}

// 	if m.MeterSerialNumber != nil {
// 		str += fmt.Sprintf(", meterSerialNumber:%s", *m.MeterSerialNumber)
// 	}

// 	if m.MeterType != nil {
// 		str += fmt.Sprintf(", meterType:%s", *m.MeterType)
// 	}

// 	str += ")"

// 	return str
// }

// type bootNotificationReqInput struct {
// 	ChargeBoxSerialNumber   string // Optional
// 	ChargePointModel        string // Required
// 	ChargePointSerialNumber string // Optional
// 	ChargePointVendor       string // Required
// 	FirmwareVersion         string // Optional
// 	Iccid                   string // Optional
// 	Imsi                    string // Optional
// 	MeterSerialNumber       string // Optional
// 	MeterType               string // Optional
// }

// func setChargeBoxSerialNumber(raw string) (sharedtypes.CiString25Type, error) {
// 	chargeBoxSerialNumber, err := sharedtypes.CiString25(raw)
// 	if err != nil {
// 		return sharedtypes.CiString25Type{}, fmt.Errorf("setChargeBoxSerialNumber: %w", err)
// 	}

// 	return chargeBoxSerialNumber, nil
// }

// func setChargePointModel(raw string) (sharedtypes.CiString20Type, error) {
// 	chargePointModel, err := sharedtypes.CiString20(raw)
// 	if err != nil {
// 		return sharedtypes.CiString20Type{}, fmt.Errorf("setChargePointModel: %w", err)
// 	}

// 	return chargePointModel, nil
// }

// func setChargePointSerialNumber(raw string) (sharedtypes.CiString25Type, error) {
// 	chargePointSerialNumber, err := sharedtypes.CiString25(raw)
// 	if err != nil {
// 		return sharedtypes.CiString25Type{}, fmt.Errorf("setChargePointSerialNumber: %w", err)
// 	}

// 	return chargePointSerialNumber, nil
// }

// func setChargePointVendor(raw string) (sharedtypes.CiString20Type, error) {
// 	chargePointVendor, err := sharedtypes.CiString20(raw)
// 	if err != nil {
// 		return sharedtypes.CiString20Type{}, fmt.Errorf("setChargePointVendor: %w", err)
// 	}

// 	return chargePointVendor, nil
// }

// func setFirmwareVersion(raw string) (sharedtypes.CiString50Type, error) {
// 	firmwareVersion, err := sharedtypes.CiString50(raw)
// 	if err != nil {
// 		return sharedtypes.CiString50Type{}, fmt.Errorf("setFirmwareVersion: %w", err)
// 	}

// 	return firmwareVersion, nil
// }

// func setIccid(raw string) (sharedtypes.CiString20Type, error) {
// 	iccid, err := sharedtypes.CiString20(raw)
// 	if err != nil {
// 		return sharedtypes.CiString20Type{}, fmt.Errorf("setIccid: %w", err)
// 	}

// 	return iccid, nil
// }

// func setImsi(raw string) (sharedtypes.CiString20Type, error) {
// 	imsi, err := sharedtypes.CiString20(raw)
// 	if err != nil {
// 		return sharedtypes.CiString20Type{}, fmt.Errorf("setImsi: %w", err)
// 	}

// 	return imsi, nil
// }

// func setMeterSerialNumber(raw string) (sharedtypes.CiString25Type, error) {
// 	meterSerialNumber, err := sharedtypes.CiString25(raw)
// 	if err != nil {
// 		return sharedtypes.CiString25Type{}, fmt.Errorf("setMeterSerialNumber: %w", err)
// 	}

// 	return meterSerialNumber, nil
// }

// func setMeterType(raw string) (sharedtypes.CiString25Type, error) {
// 	meterType, err := sharedtypes.CiString25(raw)
// 	if err != nil {
// 		return sharedtypes.CiString25Type{}, fmt.Errorf("setMeterType: %w", err)
// 	}

// 	return meterType, nil
// }

// func setOptional[T any](raw string, setter func(string) (T, error)) (*T, error) {
// 	if raw == "" {
// 		return (*T)(nil), nil
// 	}

// 	val, err := setter(raw)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &val, nil
// }
