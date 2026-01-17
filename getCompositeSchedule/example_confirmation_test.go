package getCompositeSchedule_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getCompositeSchedule"
	gt "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
)

const (
	exampleConnectorIdConf   = 1
	exampleScheduleStartConf = "2025-01-15T10:00:00Z"
	exampleDurationConf      = 3600
	exampleStartPeriodConf   = 0
	exampleLimitConf         = 32.0
	exampleStartPeriodSecond = 1800
	exampleLimitSecondPeriod = 16.0
	statusLabel              = "Status:"
)

// ExampleConf demonstrates creating a GetCompositeSchedule.conf message
// with status Accepted and no optional fields.
func ExampleConf() {
	conf, err := getCompositeSchedule.Conf(getCompositeSchedule.ConfInput{
		Status:           "Accepted",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: Accepted
}

// ExampleConf_rejected demonstrates creating a GetCompositeSchedule.conf
// message with status Rejected.
func ExampleConf_rejected() {
	conf, err := getCompositeSchedule.Conf(getCompositeSchedule.ConfInput{
		Status:           "Rejected",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	// Output:
	// Status: Rejected
}

// ExampleConf_withAllFields demonstrates creating a GetCompositeSchedule.conf
// message with all optional fields including a composite charging schedule.
func ExampleConf_withAllFields() {
	connectorId := exampleConnectorIdConf
	scheduleStart := exampleScheduleStartConf
	duration := exampleDurationConf

	conf, err := getCompositeSchedule.Conf(getCompositeSchedule.ConfInput{
		Status:        "Accepted",
		ConnectorId:   &connectorId,
		ScheduleStart: &scheduleStart,
		ChargingSchedule: &gt.ChargingScheduleInput{
			Duration:         &duration,
			ChargingRateUnit: "A",
			ChargingSchedulePeriod: []gt.ChargingSchedulePeriodInput{
				{
					StartPeriod:  exampleStartPeriodConf,
					Limit:        exampleLimitConf,
					NumberPhases: nil,
				},
				{
					StartPeriod:  exampleStartPeriodSecond,
					Limit:        exampleLimitSecondPeriod,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(statusLabel, conf.Status.String())
	fmt.Println("ConnectorId:", conf.ConnectorId.Value())
	fmt.Println("Periods:", len(conf.ChargingSchedule.ChargingSchedulePeriod()))
	// Output:
	// Status: Accepted
	// ConnectorId: 1
	// Periods: 2
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := getCompositeSchedule.Conf(getCompositeSchedule.ConfInput{
		Status:           "Invalid",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		fmt.Println("Error: invalid status")
	}
	// Output:
	// Error: invalid status
}
