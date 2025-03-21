// Package chargePoint defines messages initiated by the Charge Point (CP)
// and sent to the Central System (CSMS) as part of the OCPP 1.6J protocol.
package chargePoint

import (
	"time"

	"github.com/aasanchez/ocpp16_messages/enums"
	"github.com/aasanchez/ocpp16_messages/models"
)

// BootNotificationReq represents a BootNotification.req message,
// sent by a Charge Point to the Central System (CSMS) when it boots up.
//
// This message announces the CP's identity, vendor, and optional metadata such as firmware,
// serial numbers, and connectivity identifiers (ICCID, IMSI).
type BootNotificationReq struct {
	// ChargePointModel is a case-insensitive identifier of the model of the Charge Point.
	ChargePointModel models.CiString20Type `json:"chargePointModel"`

	// ChargePointVendor is a case-insensitive name of the Charge Point vendor.
	ChargePointVendor models.CiString20Type `json:"chargePointVendor"`

	// ChargeBoxSerialNumber is an optional case-insensitive serial number of the Charge Box.
	ChargeBoxSerialNumber *models.CiString25Type `json:"chargeBoxSerialNumber,omitempty"`

	// ChargePointSerialNumber is an optional serial number of the main Charge Point component.
	ChargePointSerialNumber *models.CiString25Type `json:"chargePointSerialNumber,omitempty"`

	// FirmwareVersion is the firmware version running on the Charge Point.
	FirmwareVersion *models.CiString50Type `json:"firmwareVersion,omitempty"`

	// Iccid is the integrated circuit card identifier (ICCID) of the modem’s SIM card.
	Iccid *models.CiString20Type `json:"iccid,omitempty"`

	// Imsi is the international mobile subscriber identity (IMSI) of the modem’s SIM card.
	Imsi *models.CiString20Type `json:"imsi,omitempty"`

	// MeterSerialNumber is an optional serial number of the main power meter.
	MeterSerialNumber *models.CiString25Type `json:"meterSerialNumber,omitempty"`

	// MeterType describes the type of meter (e.g., analog, digital).
	MeterType *models.CiString25Type `json:"meterType,omitempty"`
}

// BootNotificationConf represents a BootNotification.conf message,
// sent by the Central System (CSMS) in response to a BootNotification.req.
//
// This response tells the Charge Point whether it is accepted, pending,
// or rejected by the CSMS, and includes the current server time and
// heartbeat interval.
type BootNotificationConf struct {
	// CurrentTime is the server's current time at the moment of response.
	CurrentTime time.Time `json:"currentTime"`

	// Interval is the heartbeat interval in seconds that the CP should follow.
	Interval int `json:"interval"`

	// Status indicates whether the CP was accepted, rejected, or is pending.
	Status enums.RegistrationStatus `json:"status"`
}
