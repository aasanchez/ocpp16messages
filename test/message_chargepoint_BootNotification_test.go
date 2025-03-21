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
			name: "Empty ChargePointModel",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  models.CiString20Type(""),
				ChargePointVendor: models.CiString20Type("VendorY"),
			},
			wantErr: true,
		},
		{
			name: "ChargePointModel exceeds max length",
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
				ChargeBoxSerialNumber:   ptr(models.CiString25Type("CB123")),
				ChargePointSerialNumber: ptr(models.CiString25Type("CPSN-001")),
				FirmwareVersion:         ptr(models.CiString50Type("1.2.3")),
				Iccid:                   ptr(models.CiString20Type("ICCID9876")),
				Imsi:                    ptr(models.CiString20Type("IMSI5555")),
				MeterSerialNumber:       ptr(models.CiString25Type("MSN-007")),
				MeterType:               ptr(models.CiString25Type("TYPE-A")),
			},
		},
		{
			name: "FirmwareVersion too long",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "VendorY",
				FirmwareVersion:   ptr(models.CiString50Type("123456789012345678901234567890123456789012345678901")), // 51 chars
			},
			wantErr: true,
		},
		{
			name: "MeterSerialNumber too long",
			req: chargePoint.BootNotificationReq{
				ChargePointModel:  "ModelX",
				ChargePointVendor: "VendorY",
				MeterSerialNumber: ptr(models.CiString25Type("12345678901234567890123456")), // 26 chars
			},
			wantErr: true,
		},
		{
			name: "Iccid is empty string",
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
				t.Errorf("ValidateBootNotificationReq() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
		},
		{
			name: "Valid Pending",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationPending,
				CurrentTime: time.Now(),
				Interval:    10,
			},
		},
		{
			name: "Valid Rejected",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationRejected,
				CurrentTime: time.Now(),
				Interval:    15,
			},
		},
		{
			name: "Invalid status",
			conf: chargePoint.BootNotificationConf{
				Status:      enums.RegistrationStatus("InvalidStatus"),
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
				Interval:    -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validators.ValidateBootNotificationConf(tt.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateBootNotificationConf() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// ptr is a generic helper to create pointers of any type.
func ptr[T any](v T) *T {
	return &v
}
