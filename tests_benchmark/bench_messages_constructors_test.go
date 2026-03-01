//go:build bench

package benchmark

import (
	"testing"

	ar "github.com/aasanchez/ocpp16messages/authorize"
	bn "github.com/aasanchez/ocpp16messages/bootNotification"
	dt "github.com/aasanchez/ocpp16messages/dataTransfer"
	gcs "github.com/aasanchez/ocpp16messages/getCompositeSchedule"
	gconf "github.com/aasanchez/ocpp16messages/getConfiguration"
	gct "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	sll "github.com/aasanchez/ocpp16messages/sendLocalList"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	scp "github.com/aasanchez/ocpp16messages/setChargingProfile"
	scpt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	tm "github.com/aasanchez/ocpp16messages/triggerMessage"
	tmt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func BenchmarkAuthorizeReq(b *testing.B) {
	b.ReportAllocs()

	input := ar.ReqInput{IdTag: "TAG-1"}

	for i := 0; i < b.N; i++ {
		if _, err := ar.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAuthorizeConf(b *testing.B) {
	b.ReportAllocs()

	expiry := sampleTimestamp
	parent := "PARENT-1"
	input := ar.ConfInput{
		Status:      st.AuthorizationStatusAccepted.String(),
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	for i := 0; i < b.N; i++ {
		if _, err := ar.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBootNotificationReq(b *testing.B) {
	b.ReportAllocs()

	firmwareVersion := "1.0.0"
	input := bn.ReqInput{
		ChargePointVendor: "Vendor",
		ChargePointModel:  "Model",
		FirmwareVersion:   &firmwareVersion,
	}

	for i := 0; i < b.N; i++ {
		if _, err := bn.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBootNotificationConf(b *testing.B) {
	b.ReportAllocs()

	input := bn.ConfInput{
		Status:      "Accepted",
		CurrentTime: sampleTimestamp,
		Interval:    60,
	}

	for i := 0; i < b.N; i++ {
		if _, err := bn.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTransferReq(b *testing.B) {
	b.ReportAllocs()

	messageId := "Message-1"
	data := "payload"

	input := dt.ReqInput{
		VendorId:  "Vendor-1",
		MessageId: &messageId,
		Data:      &data,
	}

	for i := 0; i < b.N; i++ {
		if _, err := dt.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTransferConf(b *testing.B) {
	b.ReportAllocs()

	data := "payload"
	input := dt.ConfInput{
		Status: "Accepted",
		Data:   &data,
	}

	for i := 0; i < b.N; i++ {
		if _, err := dt.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetCompositeScheduleReq(b *testing.B) {
	b.ReportAllocs()

	unit := st.ChargingRateUnitWatts.String()
	input := gcs.ReqInput{
		ConnectorId:      0,
		Duration:         60,
		ChargingRateUnit: &unit,
	}

	for i := 0; i < b.N; i++ {
		if _, err := gcs.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetCompositeScheduleConf_WithSchedule(b *testing.B) {
	b.ReportAllocs()

	connectorId := 0
	scheduleStart := sampleTimestamp

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

	input := gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      &connectorId,
		ScheduleStart:    &scheduleStart,
		ChargingSchedule: &chargingScheduleInput,
	}

	for i := 0; i < b.N; i++ {
		if _, err := gcs.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetConfigurationConf_ManyKeys(b *testing.B) {
	b.ReportAllocs()

	value := "60"

	var configurationKeys []gct.KeyValueInput
	for i := 0; i < 25; i++ {
		configurationKeys = append(configurationKeys, gct.KeyValueInput{
			Key:      "HeartbeatInterval",
			Readonly: false,
			Value:    &value,
		})
	}

	var unknownKeys []string
	for i := 0; i < 25; i++ {
		unknownKeys = append(unknownKeys, "UnknownKey")
	}

	input := gconf.ConfInput{
		ConfigurationKey: configurationKeys,
		UnknownKey:       unknownKeys,
	}

	for i := 0; i < b.N; i++ {
		if _, err := gconf.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSendLocalListReq_ManyEntries(b *testing.B) {
	b.ReportAllocs()

	var entries []slt.AuthorizationDataInput
	for i := 0; i < 25; i++ {
		entries = append(entries, slt.AuthorizationDataInput{IdTag: "TAG-1"})
	}

	input := sll.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: entries,
		UpdateType:             slt.UpdateTypeFull.String(),
	}

	for i := 0; i < b.N; i++ {
		if _, err := sll.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetChargingProfileReq(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		if _, err := scp.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTriggerMessageReq(b *testing.B) {
	b.ReportAllocs()

	requestedMessage := tmt.MessageTriggerHeartbeat.String()
	connectorId := 1

	input := tm.ReqInput{
		RequestedMessage: requestedMessage,
		ConnectorId:      &connectorId,
	}

	for i := 0; i < b.N; i++ {
		if _, err := tm.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}
