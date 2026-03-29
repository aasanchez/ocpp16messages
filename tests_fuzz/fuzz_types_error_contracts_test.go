//go:build fuzz

package fuzz

import (
	"errors"
	"strings"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzChargingProfileMultiFieldErrors(f *testing.F) {
	f.Add(-1, -1, "bad-purpose", "bad-kind", "not-a-time")
	f.Add(0, 0, "", "", "")
	f.Add(99999, 99999, "x", "y", "z")

	f.Fuzz(func(
		t *testing.T,
		profileId, stackLevel int,
		purpose, kind, validFrom string,
	) {
		if len(purpose) > maxFuzzLen ||
			len(kind) > maxFuzzLen ||
			len(validFrom) > maxFuzzLen {
			t.Skip()
		}

		validFromPtr := &validFrom

		_, err := types.NewChargingProfile(
			types.ChargingProfileInput{
				ChargingProfileId:      profileId,
				StackLevel:             stackLevel,
				ChargingProfilePurpose: purpose,
				ChargingProfileKind:    kind,
				ValidFrom:              validFromPtr,
				ChargingSchedule: types.ChargingScheduleInput{
					ChargingRateUnit:       "W",
					ChargingSchedulePeriod: nil,
				},
			},
		)
		if err == nil {
			return
		}

		// Every error must wrap a sentinel
		if !errors.Is(err, types.ErrInvalidValue) &&
			!errors.Is(err, types.ErrEmptyValue) {
			t.Fatalf(
				"error = %v, want wrapping sentinel",
				err,
			)
		}

		errStr := err.Error()

		// If purpose is invalid, error must mention it
		purposeValid := purpose == "ChargePointMaxProfile" ||
			purpose == "TxDefaultProfile" ||
			purpose == "TxProfile"
		if !purposeValid &&
			!strings.Contains(errStr, "chargingProfilePurpose") {
			t.Fatalf(
				"invalid purpose %q not in error: %v",
				purpose, err,
			)
		}

		// If kind is invalid, error must mention it
		kindValid := kind == "Absolute" ||
			kind == "Recurring" ||
			kind == "Relative"
		if !kindValid &&
			!strings.Contains(errStr, "chargingProfileKind") {
			t.Fatalf(
				"invalid kind %q not in error: %v",
				kind, err,
			)
		}
	})
}

func FuzzSampledValueMultiFieldErrors(f *testing.F) {
	f.Add("", "bad-ctx", "bad-fmt", "bad-meas")
	f.Add("ok", "bad", "bad", "bad")

	f.Fuzz(func(
		t *testing.T,
		value, context, format, measurand string,
	) {
		if len(value) > maxFuzzLen ||
			len(context) > maxFuzzLen ||
			len(format) > maxFuzzLen ||
			len(measurand) > maxFuzzLen {
			t.Skip()
		}

		ctxPtr := &context
		fmtPtr := &format
		measPtr := &measurand

		_, err := types.NewSampledValue(
			types.SampledValueInput{
				Value:     value,
				Context:   ctxPtr,
				Format:    fmtPtr,
				Measurand: measPtr,
			},
		)
		if err == nil {
			return
		}

		if !errors.Is(err, types.ErrInvalidValue) &&
			!errors.Is(err, types.ErrEmptyValue) {
			t.Fatalf(
				"error = %v, want wrapping sentinel",
				err,
			)
		}

		errStr := err.Error()

		// If value is empty, error must mention it
		if value == "" &&
			!strings.Contains(errStr, "value") {
			t.Fatalf(
				"empty value not in error: %v",
				err,
			)
		}
	})
}

func FuzzKeyValueMultiFieldErrors(f *testing.F) {
	f.Add("", "")
	f.Add("\x01", "\x01")
	f.Add("", "valid")

	f.Fuzz(func(t *testing.T, key, value string) {
		if len(key) > maxFuzzLen ||
			len(value) > maxFuzzLen {
			t.Skip()
		}

		valuePtr := &value

		_, err := types.NewKeyValue(types.KeyValueInput{
			Key:      key,
			Readonly: false,
			Value:    valuePtr,
		})
		if err == nil {
			return
		}

		if !errors.Is(err, types.ErrInvalidValue) &&
			!errors.Is(err, types.ErrEmptyValue) {
			t.Fatalf(
				"error = %v, want wrapping sentinel",
				err,
			)
		}

		errStr := err.Error()

		// If key is empty, error must mention key
		if key == "" &&
			!strings.Contains(errStr, "key") {
			t.Fatalf(
				"empty key not in error: %v",
				err,
			)
		}
	})
}

func FuzzChargingScheduleMultiFieldErrors(f *testing.F) {
	f.Add("bad-unit", -1, -1.0)
	f.Add("", 0, 0.0)

	f.Fuzz(func(
		t *testing.T,
		rateUnit string,
		startPeriod int,
		limit float64,
	) {
		if len(rateUnit) > maxFuzzLen {
			t.Skip()
		}

		_, err := types.NewChargingSchedule(
			types.ChargingScheduleInput{
				ChargingRateUnit: rateUnit,
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod: startPeriod,
						Limit:       limit,
					},
				},
			},
		)
		if err == nil {
			return
		}

		if !errors.Is(err, types.ErrInvalidValue) &&
			!errors.Is(err, types.ErrEmptyValue) {
			t.Fatalf(
				"error = %v, want wrapping sentinel",
				err,
			)
		}
	})
}
