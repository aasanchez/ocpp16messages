package test

import (
	"testing"
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/messages/chargePoint"
	"github.com/aasanchez/ocpp16_messages/validators"
)

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
