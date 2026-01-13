// Package firmwareStatusNotification implements the OCPP 1.6
// FirmwareStatusNotification message.
//
// # Handling Rules
//
// A Charge Point SHALL inform the Central System about the progress of a
// firmware update by sending FirmwareStatusNotification.req messages.
//
// These notifications report the status of firmware download and
// installation initiated by the Central System via FirmwareUpdate.req.
//
// The Charge Point SHALL only send status Idle after receiving a
// TriggerMessage for FirmwareStatusNotification and when it is not
// actively downloading or installing firmware.
//
// Upon receipt of FirmwareStatusNotification.req, the Central System
// SHALL respond with FirmwareStatusNotification.conf.
package firmwareStatusNotification
