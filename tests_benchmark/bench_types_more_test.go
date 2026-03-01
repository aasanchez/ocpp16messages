//go:build bench

package benchmark

import (
	"testing"

	gct "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	scpt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func BenchmarkNewChargingSchedulePeriod_WithNumberPhases(b *testing.B) {
	b.ReportAllocs()

	phases := 3
	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  0,
		Limit:        16,
		NumberPhases: &phases,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewChargingSchedulePeriod(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetConfigurationNewKeyValue_ValueNil(b *testing.B) {
	b.ReportAllocs()

	input := gct.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    nil,
	}

	for i := 0; i < b.N; i++ {
		if _, err := gct.NewKeyValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetConfigurationNewKeyValue_ValueSet(b *testing.B) {
	b.ReportAllocs()

	value := "60"
	input := gct.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	}

	for i := 0; i < b.N; i++ {
		if _, err := gct.NewKeyValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSendLocalListNewAuthorizationData_WithIdTagInfo(b *testing.B) {
	b.ReportAllocs()

	expiry := sampleTimestamp
	parentIdTag := "PARENT-1"

	input := slt.AuthorizationDataInput{
		IdTag: "TAG-1",
		IdTagInfo: &slt.IdTagInfoInput{
			Status:      st.AuthorizationStatusAccepted.String(),
			ExpiryDate:  &expiry,
			ParentIdTag: &parentIdTag,
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := slt.NewAuthorizationData(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetChargingProfileNewChargingProfile(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	scheduleStart := sampleTimestamp
	minChargingRate := 0.0
	periods := []st.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := scpt.ChargingProfileInput{
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

	for i := 0; i < b.N; i++ {
		if _, err := scpt.NewChargingProfile(input); err != nil {
			b.Fatal(err)
		}
	}
}
