//go:build race

package race

import (
	"fmt"
	"testing"

	bn "github.com/aasanchez/ocpp16messages/bootnotification"
	ccp "github.com/aasanchez/ocpp16messages/clearchargingprofile"
	dt "github.com/aasanchez/ocpp16messages/datatransfer"
	csc "github.com/aasanchez/ocpp16messages/getcompositeschedule"
	gconf "github.com/aasanchez/ocpp16messages/getconfiguration"
	gct "github.com/aasanchez/ocpp16messages/getconfiguration/types"
	gd "github.com/aasanchez/ocpp16messages/getdiagnostics"
	mv "github.com/aasanchez/ocpp16messages/metervalues"
	mvt "github.com/aasanchez/ocpp16messages/metervalues/types"
	rst "github.com/aasanchez/ocpp16messages/remotestarttransaction"
	rn "github.com/aasanchez/ocpp16messages/reservenow"
	sl "github.com/aasanchez/ocpp16messages/sendlocallist"
	slt "github.com/aasanchez/ocpp16messages/sendlocallist/types"
	scp "github.com/aasanchez/ocpp16messages/setchargingprofile"
	scpt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	stt "github.com/aasanchez/ocpp16messages/starttransaction"
	sn "github.com/aasanchez/ocpp16messages/statusnotification"
	snt "github.com/aasanchez/ocpp16messages/statusnotification/types"
	stp "github.com/aasanchez/ocpp16messages/stoptransaction"
	tm "github.com/aasanchez/ocpp16messages/triggermessage"
	tmt "github.com/aasanchez/ocpp16messages/triggermessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
	uf "github.com/aasanchez/ocpp16messages/updatefirmware"
)

func TestRace_BootNotificationReq(t *testing.T) {
	t.Parallel()

	firmwareVersion := "1.0.0"
	input := bn.ReqInput{
		ChargePointVendor: "Vendor",
		ChargePointModel:  "Model",
		FirmwareVersion:   &firmwareVersion,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := bn.Req(input)
		if err != nil {
			return fmt.Errorf("bootnotification.Req: %w", err)
		}
		return nil
	})
}

func TestRace_BootNotificationConf(t *testing.T) {
	t.Parallel()

	input := bn.ConfInput{
		Status:      "Accepted",
		CurrentTime: "2025-01-02T15:00:00Z",
		Interval:    60,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := bn.Conf(input)
		if err != nil {
			return fmt.Errorf("bootnotification.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_StartTransactionReq(t *testing.T) {
	t.Parallel()

	reservationId := 42
	input := stt.ReqInput{
		ConnectorId:   1,
		IdTag:         "TAG-1",
		MeterStart:    100,
		Timestamp:     "2025-01-02T15:00:00Z",
		ReservationId: &reservationId,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := stt.Req(input)
		if err != nil {
			return fmt.Errorf("starttransaction.Req: %w", err)
		}
		return nil
	})
}

func TestRace_StartTransactionConf(t *testing.T) {
	t.Parallel()

	expiry := "2025-01-03T15:00:00Z"
	parent := "PARENT-1"
	input := stt.ConfInput{
		TransactionId: 1,
		Status:        "Accepted",
		ExpiryDate:    &expiry,
		ParentIdTag:   &parent,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := stt.Conf(input)
		if err != nil {
			return fmt.Errorf("starttransaction.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_StopTransactionReq(t *testing.T) {
	t.Parallel()

	idTag := "TAG-1"
	reason := "Local"

	sampledValues := []mvt.SampledValueInput{{Value: "100"}}
	transactionData := []mvt.MeterValueInput{
		{
			Timestamp:    "2025-01-02T15:00:00Z",
			SampledValue: sampledValues,
		},
	}

	input := stp.ReqInput{
		TransactionId:   1,
		IdTag:           &idTag,
		MeterStop:       200,
		Timestamp:       "2025-01-02T15:01:00Z",
		Reason:          &reason,
		TransactionData: transactionData,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := stp.Req(input)
		if err != nil {
			return fmt.Errorf("stoptransaction.Req: %w", err)
		}
		return nil
	})
}

func TestRace_StopTransactionConf(t *testing.T) {
	t.Parallel()

	status := "Accepted"
	expiry := "2025-01-03T15:00:00Z"
	parent := "PARENT-1"
	input := stp.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := stp.Conf(input)
		if err != nil {
			return fmt.Errorf("stoptransaction.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_MeterValuesConf(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := mv.Conf(mv.ConfInput{})
		if err != nil {
			return fmt.Errorf("metervalues.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_TriggerMessageReq(t *testing.T) {
	t.Parallel()

	requestedMessage := tmt.MessageTriggerHeartbeat.String()
	connectorId := 1
	input := tm.ReqInput{
		RequestedMessage: requestedMessage,
		ConnectorId:      &connectorId,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := tm.Req(input)
		if err != nil {
			return fmt.Errorf("triggermessage.Req: %w", err)
		}
		return nil
	})
}

func TestRace_DataTransferReq(t *testing.T) {
	t.Parallel()

	messageId := "Message-1"
	data := "payload"
	input := dt.ReqInput{
		VendorId:  "Vendor-1",
		MessageId: &messageId,
		Data:      &data,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := dt.Req(input)
		if err != nil {
			return fmt.Errorf("datatransfer.Req: %w", err)
		}
		return nil
	})
}

func TestRace_StatusNotificationReq(t *testing.T) {
	t.Parallel()

	timestamp := "2025-01-02T15:00:00Z"
	info := "Info"
	vendorId := "Vendor-1"
	vendorErrorCode := "E-1"

	input := sn.ReqInput{
		ConnectorId:     0,
		ErrorCode:       snt.ErrCodeNoError.String(),
		Status:          snt.ChargePointStatusAvailable.String(),
		Info:            &info,
		Timestamp:       &timestamp,
		VendorId:        &vendorId,
		VendorErrorCode: &vendorErrorCode,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := sn.Req(input)
		if err != nil {
			return fmt.Errorf("statusnotification.Req: %w", err)
		}
		return nil
	})
}

func TestRace_RemoteStartTransactionReq(t *testing.T) {
	t.Parallel()

	connectorId := 1
	input := rst.ReqInput{
		IdTag:       "TAG-1",
		ConnectorId: &connectorId,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rst.Req(input)
		if err != nil {
			return fmt.Errorf("remotestarttransaction.Req: %w", err)
		}
		return nil
	})
}

func TestRace_ReserveNowReq(t *testing.T) {
	t.Parallel()

	parentIdTag := "PARENT-1"
	input := rn.ReqInput{
		ReservationId: 1,
		ConnectorId:   1,
		IdTag:         "TAG-1",
		ExpiryDate:    "2025-01-02T16:00:00Z",
		ParentIdTag:   &parentIdTag,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rn.Req(input)
		if err != nil {
			return fmt.Errorf("reservenow.Req: %w", err)
		}
		return nil
	})
}

func TestRace_UpdateFirmwareReq(t *testing.T) {
	t.Parallel()

	retries := 3
	retryInterval := 60
	input := uf.ReqInput{
		Location:      "https://example.com/firmware.bin",
		RetrieveDate:  "2025-01-02T15:00:00Z",
		Retries:       &retries,
		RetryInterval: &retryInterval,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := uf.Req(input)
		if err != nil {
			return fmt.Errorf("updatefirmware.Req: %w", err)
		}
		return nil
	})
}

func TestRace_GetDiagnosticsReq(t *testing.T) {
	t.Parallel()

	retries := 3
	retryInterval := 60
	startTime := "2025-01-02T15:00:00Z"
	stopTime := "2025-01-02T16:00:00Z"

	input := gd.ReqInput{
		Location:      "https://example.com/upload",
		Retries:       &retries,
		RetryInterval: &retryInterval,
		StartTime:     &startTime,
		StopTime:      &stopTime,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gd.Req(input)
		if err != nil {
			return fmt.Errorf("getdiagnostics.Req: %w", err)
		}
		return nil
	})
}

func TestRace_GetDiagnosticsConf(t *testing.T) {
	t.Parallel()

	fileName := "diagnostics.log"
	input := gd.ConfInput{FileName: &fileName}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gd.Conf(input)
		if err != nil {
			return fmt.Errorf("getdiagnostics.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_GetConfigurationReq(t *testing.T) {
	t.Parallel()

	keys := []string{"HeartbeatInterval"}
	input := gconf.ReqInput{Key: keys}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gconf.Req(input)
		if err != nil {
			return fmt.Errorf("getconfiguration.Req: %w", err)
		}
		return nil
	})
}

func TestRace_GetConfigurationConf(t *testing.T) {
	t.Parallel()

	value := "60"
	configKeys := []gct.KeyValueInput{
		{Key: "HeartbeatInterval", Readonly: false, Value: &value},
	}
	unknownKeys := []string{"UnknownKey"}

	input := gconf.ConfInput{
		ConfigurationKey: configKeys,
		UnknownKey:       unknownKeys,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gconf.Conf(input)
		if err != nil {
			return fmt.Errorf("getconfiguration.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_GetCompositeScheduleReq(t *testing.T) {
	t.Parallel()

	unit := st.ChargingRateUnitWatts.String()
	input := csc.ReqInput{
		ConnectorId:      0,
		Duration:         60,
		ChargingRateUnit: &unit,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := csc.Req(input)
		if err != nil {
			return fmt.Errorf("getcompositeschedule.Req: %w", err)
		}
		return nil
	})
}

func TestRace_GetCompositeScheduleConf(t *testing.T) {
	t.Parallel()

	connectorId := 0
	scheduleStart := "2025-01-02T15:00:00Z"

	duration := 60
	minChargingRate := 0.0
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}
	chargingScheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := csc.ConfInput{
		Status:           "Accepted",
		ConnectorId:      &connectorId,
		ScheduleStart:    &scheduleStart,
		ChargingSchedule: &chargingScheduleInput,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := csc.Conf(input)
		if err != nil {
			return fmt.Errorf("getcompositeschedule.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_ClearChargingProfileReq(t *testing.T) {
	t.Parallel()

	id := 1
	connectorId := 0
	stackLevel := 0
	purpose := st.TxProfile.String()

	input := ccp.ReqInput{
		Id:                     &id,
		ConnectorId:            &connectorId,
		ChargingProfilePurpose: &purpose,
		StackLevel:             &stackLevel,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := ccp.Req(input)
		if err != nil {
			return fmt.Errorf("clearchargingprofile.Req: %w", err)
		}
		return nil
	})
}

func TestRace_SendLocalListReq(t *testing.T) {
	t.Parallel()

	authList := []slt.AuthorizationDataInput{
		{IdTag: "TAG-1"},
	}

	input := sl.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: authList,
		UpdateType:             slt.UpdateTypeFull.String(),
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := sl.Req(input)
		if err != nil {
			return fmt.Errorf("sendlocallist.Req: %w", err)
		}
		return nil
	})
}

func TestRace_SendLocalListConf(t *testing.T) {
	t.Parallel()

	input := sl.ConfInput{Status: slt.UpdateStatusAccepted.String()}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := sl.Conf(input)
		if err != nil {
			return fmt.Errorf("sendlocallist.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_SetChargingProfileReq(t *testing.T) {
	t.Parallel()

	duration := 60
	scheduleStart := "2025-01-02T15:00:00Z"
	minChargingRate := 0.0
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	profileInput := scpt.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          nil,
		StackLevel:             0,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind:    scpt.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	}

	input := scp.ReqInput{
		ConnectorId:        0,
		CsChargingProfiles: profileInput,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := scp.Req(input)
		if err != nil {
			return fmt.Errorf("setchargingprofile.Req: %w", err)
		}
		return nil
	})
}

func TestRace_SetChargingProfileConf(t *testing.T) {
	t.Parallel()

	input := scp.ConfInput{
		Status: scpt.ChargingProfileStatusAccepted.String(),
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := scp.Conf(input)
		if err != nil {
			return fmt.Errorf("setchargingprofile.Conf: %w", err)
		}
		return nil
	})
}
