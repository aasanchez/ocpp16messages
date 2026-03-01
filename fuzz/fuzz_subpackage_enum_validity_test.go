//go:build fuzz

package fuzz

import (
	"testing"

	bn "github.com/aasanchez/ocpp16messages/bootNotification/types"
	cn "github.com/aasanchez/ocpp16messages/cancelReservation/types"
	av "github.com/aasanchez/ocpp16messages/changeAvailability/types"
	cg "github.com/aasanchez/ocpp16messages/changeConfiguration/types"
	cc "github.com/aasanchez/ocpp16messages/clearCache/types"
	cp "github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
	dt "github.com/aasanchez/ocpp16messages/dataTransfer/types"
	ds "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification/types"
	fs "github.com/aasanchez/ocpp16messages/firmwareStatusNotification/types"
	gc "github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
	mv "github.com/aasanchez/ocpp16messages/meterValues/types"
	rs "github.com/aasanchez/ocpp16messages/remoteStartTransaction/types"
	rt "github.com/aasanchez/ocpp16messages/remoteStopTransaction/types"
	rn "github.com/aasanchez/ocpp16messages/reserveNow/types"
	re "github.com/aasanchez/ocpp16messages/reset/types"
	sl "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	sc "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	sn "github.com/aasanchez/ocpp16messages/statusNotification/types"
	sr "github.com/aasanchez/ocpp16messages/stopTransaction/types"
	tm "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	uc "github.com/aasanchez/ocpp16messages/unlockConnector/types"
)

func assertIsValidMatchesSet(
	t *testing.T,
	input string,
	allowed map[string]struct{},
	got bool,
) {
	t.Helper()

	if got {
		if _, ok := allowed[input]; !ok {
			t.Fatalf("IsValid() = true for %q, want false", input)
		}
	} else {
		if _, ok := allowed[input]; ok {
			t.Fatalf("IsValid() = false for %q, want true", input)
		}
	}
}

func FuzzBootNotificationRegistrationStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Pending":  {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Accepted ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := bn.RegistrationStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzCancelReservationStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Rejected ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := cn.CancelReservationStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzChangeAvailabilityStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":  {},
		"Rejected":  {},
		"Scheduled": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Scheduled ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := av.AvailabilityStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzChangeAvailabilityTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Inoperative": {},
		"Operative":   {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("operative")
	f.Add("Operative ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		typ := av.AvailabilityType(input)
		assertIsValidMatchesSet(t, input, allowed, typ.IsValid())
	})
}

func FuzzChangeConfigurationStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":       {},
		"Rejected":       {},
		"RebootRequired": {},
		"NotSupported":   {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("NotSupported ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := cg.ConfigurationStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzClearCacheStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := cc.ClearCacheStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzClearChargingProfileStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Unknown":  {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Unknown ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := cp.ClearChargingProfileStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzDataTransferStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":         {},
		"Rejected":         {},
		"UnknownMessageId": {},
		"UnknownVendor":    {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("UnknownVendor ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := dt.DataTransferStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzDiagnosticsStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Idle":         {},
		"Uploaded":     {},
		"UploadFailed": {},
		"Uploading":    {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("idle")
	f.Add("Uploaded ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := ds.DiagnosticsStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzFirmwareStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Downloaded":         {},
		"DownloadFailed":     {},
		"Downloading":        {},
		"Idle":               {},
		"InstallationFailed": {},
		"Installing":         {},
		"Installed":          {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("idle")
	f.Add("Installed ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := fs.FirmwareStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzGetCompositeScheduleStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Rejected ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := gc.GetCompositeScheduleStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzMeterValuesLocationIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Body":   {},
		"Cable":  {},
		"EV":     {},
		"Inlet":  {},
		"Outlet": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("body")
	f.Add("Outlet ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		loc := mv.Location(input)
		assertIsValidMatchesSet(t, input, allowed, loc.IsValid())
	})
}

func FuzzMeterValuesMeasurandIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Current.Export":                  {},
		"Current.Import":                  {},
		"Current.Offered":                 {},
		"Energy.Active.Export.Register":   {},
		"Energy.Active.Import.Register":   {},
		"Energy.Reactive.Export.Register": {},
		"Energy.Reactive.Import.Register": {},
		"Energy.Active.Export.Interval":   {},
		"Energy.Active.Import.Interval":   {},
		"Energy.Reactive.Export.Interval": {},
		"Energy.Reactive.Import.Interval": {},
		"Frequency":                       {},
		"Power.Active.Export":             {},
		"Power.Active.Import":             {},
		"Power.Factor":                    {},
		"Power.Offered":                   {},
		"Power.Reactive.Export":           {},
		"Power.Reactive.Import":           {},
		"RPM":                             {},
		"SoC":                             {},
		"Temperature":                     {},
		"Voltage":                         {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("energy.active.import.register")
	f.Add("Energy.Active.Import.Register ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		measurand := mv.Measurand(input)
		assertIsValidMatchesSet(t, input, allowed, measurand.IsValid())
	})
}

func FuzzMeterValuesPhaseIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"L1":    {},
		"L2":    {},
		"L3":    {},
		"N":     {},
		"L1-N":  {},
		"L2-N":  {},
		"L3-N":  {},
		"L1-L2": {},
		"L2-L3": {},
		"L3-L1": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("L1N")
	f.Add("L1-N ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		phase := mv.Phase(input)
		assertIsValidMatchesSet(t, input, allowed, phase.IsValid())
	})
}

func FuzzMeterValuesReadingContextIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Interruption.Begin": {},
		"Interruption.End":   {},
		"Other":              {},
		"Sample.Clock":       {},
		"Sample.Periodic":    {},
		"Transaction.Begin":  {},
		"Transaction.End":    {},
		"Trigger":            {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("OtherContext")
	f.Add("Transaction.Begin ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		ctx := mv.ReadingContext(input)
		assertIsValidMatchesSet(t, input, allowed, ctx.IsValid())
	})
}

func FuzzMeterValuesUnitOfMeasureIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Wh":         {},
		"kWh":        {},
		"varh":       {},
		"kvarh":      {},
		"W":          {},
		"kW":         {},
		"VA":         {},
		"kVA":        {},
		"var":        {},
		"kvar":       {},
		"A":          {},
		"V":          {},
		"Celsius":    {},
		"Fahrenheit": {},
		"K":          {},
		"Percent":    {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("wh")
	f.Add("kwh")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		unit := mv.UnitOfMeasure(input)
		assertIsValidMatchesSet(t, input, allowed, unit.IsValid())
	})
}

func FuzzMeterValuesValueFormatIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Raw":        {},
		"SignedData": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("raw")
	f.Add("SignedData ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		format := mv.ValueFormat(input)
		assertIsValidMatchesSet(t, input, allowed, format.IsValid())
	})
}

func FuzzRemoteStartTransactionStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := rs.RemoteStartTransactionStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzRemoteStopTransactionStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := rt.RemoteStopTransactionStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzReservationStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":    {},
		"Faulted":     {},
		"Occupied":    {},
		"Rejected":    {},
		"Unavailable": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Unavailable ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := rn.ReservationStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzResetStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted": {},
		"Rejected": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := re.ResetStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzResetTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Hard": {},
		"Soft": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("hard")
	f.Add("Soft ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		typ := re.ResetType(input)
		assertIsValidMatchesSet(t, input, allowed, typ.IsValid())
	})
}

func FuzzSendLocalListUpdateStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":        {},
		"Failed":          {},
		"NotSupported":    {},
		"VersionMismatch": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("Failed ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := sl.UpdateStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzSendLocalListUpdateTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Full":         {},
		"Differential": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("full")
	f.Add("Full ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		typ := sl.UpdateType(input)
		assertIsValidMatchesSet(t, input, allowed, typ.IsValid())
	})
}

func FuzzChargingProfileKindTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Absolute":  {},
		"Recurring": {},
		"Relative":  {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("absolute")
	f.Add("Relative ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		kind := sc.ChargingProfileKindType(input)
		assertIsValidMatchesSet(t, input, allowed, kind.IsValid())
	})
}

func FuzzChargingProfileStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":     {},
		"Rejected":     {},
		"NotSupported": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("NotSupported ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := sc.ChargingProfileStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzRecurrencyKindTypeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Daily":  {},
		"Weekly": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("daily")
	f.Add("Weekly ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		kind := sc.RecurrencyKindType(input)
		assertIsValidMatchesSet(t, input, allowed, kind.IsValid())
	})
}

func FuzzChargePointErrorCodeIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"ConnectorLockFailure": {},
		"EVCommunicationError": {},
		"GroundFailure":        {},
		"HighTemperature":      {},
		"InternalError":        {},
		"LocalListConflict":    {},
		"NoError":              {},
		"OtherError":           {},
		"OverCurrentFailure":   {},
		"OverVoltage":          {},
		"PowerMeterFailure":    {},
		"PowerSwitchFailure":   {},
		"ReaderFailure":        {},
		"ResetFailure":         {},
		"UnderVoltage":         {},
		"WeakSignal":           {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("NoError ")
	f.Add("noerror")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		code := sn.ChargePointErrorCode(input)
		assertIsValidMatchesSet(t, input, allowed, code.IsValid())
	})
}

func FuzzChargePointStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Available":     {},
		"Preparing":     {},
		"Charging":      {},
		"SuspendedEV":   {},
		"SuspendedEVSE": {},
		"Finishing":     {},
		"Reserved":      {},
		"Unavailable":   {},
		"Faulted":       {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("available")
	f.Add("Available ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := sn.ChargePointStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzStopReasonIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"DeAuthorized":   {},
		"EmergencyStop":  {},
		"EVDisconnected": {},
		"HardReset":      {},
		"Local":          {},
		"Other":          {},
		"PowerLoss":      {},
		"Reboot":         {},
		"Remote":         {},
		"SoftReset":      {},
		"UnlockCommand":  {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("local")
	f.Add("Local ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		reason := sr.Reason(input)
		assertIsValidMatchesSet(t, input, allowed, reason.IsValid())
	})
}

func FuzzMessageTriggerIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"BootNotification":              {},
		"DiagnosticsStatusNotification": {},
		"FirmwareStatusNotification":    {},
		"Heartbeat":                     {},
		"MeterValues":                   {},
		"StatusNotification":            {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("heartbeat")
	f.Add("Heartbeat ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		trigger := tm.MessageTrigger(input)
		assertIsValidMatchesSet(t, input, allowed, trigger.IsValid())
	})
}

func FuzzTriggerMessageStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Accepted":       {},
		"Rejected":       {},
		"NotImplemented": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("accepted")
	f.Add("NotImplemented ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := tm.TriggerMessageStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}

func FuzzUnlockStatusIsValid(f *testing.F) {
	allowed := map[string]struct{}{
		"Unlocked":     {},
		"UnlockFailed": {},
		"NotSupported": {},
	}

	for value := range allowed {
		f.Add(value)
	}
	f.Add("")
	f.Add("unlocked")
	f.Add("Unlocked ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		status := uc.UnlockStatus(input)
		assertIsValidMatchesSet(t, input, allowed, status.IsValid())
	})
}
