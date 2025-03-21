package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
	"github.com/aasanchez/ocpp16_messages/models"
	"github.com/aasanchez/ocpp16_messages/validators"
)

func TestBootNotificationReq(t *testing.T) {
	tests := []struct {
		name    string
		req     chargePoint.BootNotificationReq
		wantErr bool
	}{
		{
			name: "Valid minimal request",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("ModelX"),
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: false,
		},
		{
			name: "Missing model",
			req: chargePoint.BootNotificationReq{
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: true,
		},
		{
			name: "Missing vendor",
			req: chargePoint.BootNotificationReq{
				ChargePointModel: models.CiString20Type("ModelX"),
			},
			wantErr: true,
		},
		{
			name: "Empty model",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type(""),
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: true,
		},
		{
			name: "Model too long",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("123456789012345678901"),
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: true,
		},
		{
			name: "Optional firmware version valid",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("ModelX"),
				ChargePointVendor: models.CiString20Type("VendorY"),
				FirmwareVersion:   ptr(models.CiString50Type("v1.2.3")),
			},
			wantErr: false,
		},
		{
			name: "Optional firmware version too long",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("ModelX"),
				ChargePointVendor: models.CiString20Type("VendorY"),
				FirmwareVersion:   ptr(models.CiString50Type("123456789012345678901234567890123456789012345678901")),
			},
			wantErr: true,
		},
		{
			name: "All optional fields valid",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:        models.CiString20Type("ModelX"),
				ChargePointVendor:       models.CiString20Type("VendorY"),
				ChargeBoxSerialNumber:   ptr(models.CiString25Type("CB123456")),
				ChargePointSerialNumber: ptr(models.CiString25Type("CPSN7890")),
				FirmwareVersion:         ptr(models.CiString50Type("v1.0.0")),
				Iccid:                   ptr(models.CiString20Type("ICCID123456")),
				Imsi:                    ptr(models.CiString20Type("IMSI123456")),
				MeterSerialNumber:       ptr(models.CiString25Type("MSN9999")),
				MeterType:               ptr(models.CiString25Type("MT1000")),
			},
			wantErr: false,
		},
		{
			name: "Optional serial number too long",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("ModelX"),
				ChargePointVendor: models.CiString20Type("VendorY"),
				MeterSerialNumber: ptr(models.CiString25Type("12345678901234567890123456")), // 26 chars
			},
			wantErr: true,
		},
		{
			name: "Optional ICCID empty",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("ModelX"),
				ChargePointVendor: models.CiString20Type("VendorY"),
				Iccid:             ptr(models.CiString20Type("")),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validators.ValidateBootNotificationReq(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBootNotificationReq() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
