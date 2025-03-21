package test

import (
	"testing"
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
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
			name: "Missing ChargePointModel",
			req: chargePoint.BootNotificationReq{
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: true,
		},
		{
			name: "Missing ChargePointVendor",
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
			name: "Model exceeds 20 characters",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type("123456789012345678901"),
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: true,
		},
		{
			name: "All optional fields valid",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:        "ModelX",
				ChargePointVendor:       "VendorY",
				ChargeBoxSerialNumber:   ptr(models.CiString25Type("CB123456789")),
				ChargePointSerialNumber: ptr(models.CiString25Type("CP123456789")),
				FirmwareVersion:         ptr(models.CiString50Type("FW1.2.3")),
				Iccid:                   ptr(models.CiString20Type("ICCID123456")),
				Imsi:                    ptr(models.CiString20Type("IMSI987654")),
				MeterSerialNumber:       ptr(models.CiString25Type("MSN0001")),
				MeterType:               ptr(models.CiString25Type("MTYPE-A")),
			},
			wantErr: false,
		},
		{
			name: "FirmwareVersion too long",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "VendorY",
				FirmwareVersion:   ptr(models.CiString50Type("123456789012345678901234567890123456789012345678901")),
			},
			wantErr: true,
		},
		{
			name: "Optional field too long: MeterSerialNumber",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "VendorY",
				MeterSerialNumber: ptr(models.CiString25Type("12345678901234567890123456")), // 26 chars
			},
			wantErr: true,
		},
		{
			name: "Optional field empty: Iccid",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "VendorY",
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

func TestBootNotificationConf(t *testing.T) {
	tests := []struct {
		name    string
		conf    chargePoint.BootNotificationConf
		wantErr bool
	}{
		{
			name: "Valid Accepted",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationAccepted,
				CurrentTime: time.Now(),
				Interval:    30,
			},
			wantErr: false,
		},
		{
			name: "Valid Pending",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationPending,
				CurrentTime: time.Now(),
				Interval:    10,
			},
			wantErr: false,
		},
		{
			name: "Valid Rejected",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationRejected,
				CurrentTime: time.Now(),
				Interval:    15,
			},
			wantErr: false,
		},
		{
			name: "Invalid status",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationStatus("Unknown"),
				CurrentTime: time.Now(),
				Interval:    10,
			},
			wantErr: true,
		},
		{
			name: "Zero interval",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationAccepted,
				CurrentTime: time.Now(),
				Interval:    0,
			},
			wantErr: true,
		},
		{
			name: "Negative interval",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationAccepted,
				CurrentTime: time.Now(),
				Interval:    -10,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validators.ValidateBootNotificationConf(tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBootNotificationConf() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
