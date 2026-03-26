//go:build race

package race

import (
	"fmt"
	"testing"

	ar "github.com/aasanchez/ocpp16messages/authorize"
	cc "github.com/aasanchez/ocpp16messages/cancelreservation"
	cac "github.com/aasanchez/ocpp16messages/changeavailability"
	cconf "github.com/aasanchez/ocpp16messages/changeconfiguration"
	clc "github.com/aasanchez/ocpp16messages/clearcache"
	cs "github.com/aasanchez/ocpp16messages/clearcache/types"
	cpt "github.com/aasanchez/ocpp16messages/clearchargingprofile"
	cchp "github.com/aasanchez/ocpp16messages/clearchargingprofile/types"
	dt "github.com/aasanchez/ocpp16messages/datatransfer"
	dsn "github.com/aasanchez/ocpp16messages/diagnosticsstatusnotification"
	fn "github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
	gllv "github.com/aasanchez/ocpp16messages/getlocallistversion"
	hb "github.com/aasanchez/ocpp16messages/heartbeat"
	rstt "github.com/aasanchez/ocpp16messages/remotestarttransaction"
	rstp "github.com/aasanchez/ocpp16messages/remotestoptransaction"
	rn "github.com/aasanchez/ocpp16messages/reservenow"
	rst "github.com/aasanchez/ocpp16messages/reset"
	sn "github.com/aasanchez/ocpp16messages/statusnotification"
	tm "github.com/aasanchez/ocpp16messages/triggermessage"
	uc "github.com/aasanchez/ocpp16messages/unlockconnector"
	uf "github.com/aasanchez/ocpp16messages/updatefirmware"
)

func TestRace_ClearCacheReq(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := clc.Req(clc.ReqInput{})
		if err != nil {
			return fmt.Errorf("clearcache.Req: %w", err)
		}
		return nil
	})
}

func TestRace_ClearCacheConf(t *testing.T) {
	t.Parallel()

	input := clc.ConfInput{Status: cs.ClearCacheStatusAccepted.String()}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := clc.Conf(input)
		if err != nil {
			return fmt.Errorf("clearcache.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_HeartbeatReq(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := hb.Req(hb.ReqInput{})
		if err != nil {
			return fmt.Errorf("heartbeat.Req: %w", err)
		}
		return nil
	})
}

func TestRace_HeartbeatConf(t *testing.T) {
	t.Parallel()

	input := hb.ConfInput{CurrentTime: "2025-01-02T15:00:00Z"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := hb.Conf(input)
		if err != nil {
			return fmt.Errorf("heartbeat.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_CancelReservationReq(t *testing.T) {
	t.Parallel()

	input := cc.ReqInput{ReservationId: 1}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cc.Req(input)
		if err != nil {
			return fmt.Errorf("cancelreservation.Req: %w", err)
		}
		return nil
	})
}

func TestRace_CancelReservationConf(t *testing.T) {
	t.Parallel()

	input := cc.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cc.Conf(input)
		if err != nil {
			return fmt.Errorf("cancelreservation.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_RemoteStopTransactionReq(t *testing.T) {
	t.Parallel()

	input := rstp.ReqInput{TransactionId: 1}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rstp.Req(input)
		if err != nil {
			return fmt.Errorf("remotestoptransaction.Req: %w", err)
		}
		return nil
	})
}

func TestRace_RemoteStopTransactionConf(t *testing.T) {
	t.Parallel()

	input := rstp.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rstp.Conf(input)
		if err != nil {
			return fmt.Errorf("remotestoptransaction.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_GetLocalListVersionReq(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gllv.Req(gllv.ReqInput{})
		if err != nil {
			return fmt.Errorf("getlocallistversion.Req: %w", err)
		}
		return nil
	})
}

func TestRace_GetLocalListVersionConf(t *testing.T) {
	t.Parallel()

	input := gllv.ConfInput{ListVersion: 0}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := gllv.Conf(input)
		if err != nil {
			return fmt.Errorf("getlocallistversion.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_DiagnosticsStatusNotificationReq(t *testing.T) {
	t.Parallel()

	input := dsn.ReqInput{Status: "Uploaded"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := dsn.Req(input)
		if err != nil {
			return fmt.Errorf("diagnosticsstatusnotification.Req: %w", err)
		}
		return nil
	})
}

func TestRace_DiagnosticsStatusNotificationConf(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := dsn.Conf(dsn.ConfInput{})
		if err != nil {
			return fmt.Errorf("diagnosticsstatusnotification.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_FirmwareStatusNotificationReq(t *testing.T) {
	t.Parallel()

	input := fn.ReqInput{Status: "Downloaded"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := fn.Req(input)
		if err != nil {
			return fmt.Errorf("firmwarestatusnotification.Req: %w", err)
		}
		return nil
	})
}

func TestRace_FirmwareStatusNotificationConf(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := fn.Conf(fn.ConfInput{})
		if err != nil {
			return fmt.Errorf("firmwarestatusnotification.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_UnlockConnectorReq(t *testing.T) {
	t.Parallel()

	input := uc.ReqInput{ConnectorId: 1}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := uc.Req(input)
		if err != nil {
			return fmt.Errorf("unlockconnector.Req: %w", err)
		}
		return nil
	})
}

func TestRace_UnlockConnectorConf(t *testing.T) {
	t.Parallel()

	input := uc.ConfInput{Status: "Unlocked"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := uc.Conf(input)
		if err != nil {
			return fmt.Errorf("unlockconnector.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_UpdateFirmwareConf(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := uf.Conf(uf.ConfInput{})
		if err != nil {
			return fmt.Errorf("updatefirmware.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_StatusNotificationConf(t *testing.T) {
	t.Parallel()

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := sn.Conf(sn.ConfInput{})
		if err != nil {
			return fmt.Errorf("statusnotification.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_ResetReq(t *testing.T) {
	t.Parallel()

	input := rst.ReqInput{Type: "Soft"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rst.Req(input)
		if err != nil {
			return fmt.Errorf("reset.Req: %w", err)
		}
		return nil
	})
}

func TestRace_ResetConf(t *testing.T) {
	t.Parallel()

	input := rst.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rst.Conf(input)
		if err != nil {
			return fmt.Errorf("reset.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_ClearChargingProfileConf(t *testing.T) {
	t.Parallel()

	input := cpt.ConfInput{Status: cchp.ClearChargingProfileStatusAccepted.String()}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cpt.Conf(input)
		if err != nil {
			return fmt.Errorf("clearchargingprofile.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_DataTransferConf(t *testing.T) {
	t.Parallel()

	data := "payload"
	input := dt.ConfInput{Status: "Accepted", Data: &data}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := dt.Conf(input)
		if err != nil {
			return fmt.Errorf("datatransfer.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_TriggerMessageConf(t *testing.T) {
	t.Parallel()

	input := tm.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := tm.Conf(input)
		if err != nil {
			return fmt.Errorf("triggermessage.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_ReserveNowConf(t *testing.T) {
	t.Parallel()

	input := rn.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rn.Conf(input)
		if err != nil {
			return fmt.Errorf("reservenow.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_AuthorizeReq(t *testing.T) {
	t.Parallel()

	input := ar.ReqInput{IdTag: "TAG-1"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := ar.Req(input)
		if err != nil {
			return fmt.Errorf("authorize.Req: %w", err)
		}
		return nil
	})
}

func TestRace_AuthorizeConf(t *testing.T) {
	t.Parallel()

	expiry := "2025-01-02T15:00:00Z"
	parent := "PARENT-1"
	input := ar.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := ar.Conf(input)
		if err != nil {
			return fmt.Errorf("authorize.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_ChangeAvailabilityReq(t *testing.T) {
	t.Parallel()

	input := cac.ReqInput{ConnectorId: 0, Type: "Operative"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cac.Req(input)
		if err != nil {
			return fmt.Errorf("changeavailability.Req: %w", err)
		}
		return nil
	})
}

func TestRace_ChangeAvailabilityConf(t *testing.T) {
	t.Parallel()

	input := cac.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cac.Conf(input)
		if err != nil {
			return fmt.Errorf("changeavailability.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_ChangeConfigurationReq(t *testing.T) {
	t.Parallel()

	input := cconf.ReqInput{Key: "HeartbeatInterval", Value: "60"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cconf.Req(input)
		if err != nil {
			return fmt.Errorf("changeconfiguration.Req: %w", err)
		}
		return nil
	})
}

func TestRace_ChangeConfigurationConf(t *testing.T) {
	t.Parallel()

	input := cconf.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := cconf.Conf(input)
		if err != nil {
			return fmt.Errorf("changeconfiguration.Conf: %w", err)
		}
		return nil
	})
}

func TestRace_RemoteStartTransactionConf(t *testing.T) {
	t.Parallel()

	input := rstt.ConfInput{Status: "Accepted"}

	runConcurrent(t, raceWorkers, raceIterations, func(_, _ int) error {
		_, err := rstt.Conf(input)
		if err != nil {
			return fmt.Errorf("remotestarttransaction.Conf: %w", err)
		}
		return nil
	})
}
