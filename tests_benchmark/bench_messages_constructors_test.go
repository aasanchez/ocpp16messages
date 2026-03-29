//go:build bench

package benchmark

import (
	"testing"

	ar "github.com/aasanchez/ocpp16messages/authorize"
	bn "github.com/aasanchez/ocpp16messages/bootnotification"
	dt "github.com/aasanchez/ocpp16messages/datatransfer"
	gcs "github.com/aasanchez/ocpp16messages/getcompositeschedule"
	gconf "github.com/aasanchez/ocpp16messages/getconfiguration"
	sll "github.com/aasanchez/ocpp16messages/sendlocallist"
	scp "github.com/aasanchez/ocpp16messages/setchargingprofile"
	tm "github.com/aasanchez/ocpp16messages/triggermessage"
	types "github.com/aasanchez/ocpp16types"
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
		Status:      types.AuthorizationStatusAccepted.String(),
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

	unit := types.ChargingRateUnitWatts.String()
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
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	chargingScheduleInput := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
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

	var configurationKeys []types.KeyValueInput
	for i := 0; i < 25; i++ {
		configurationKeys = append(configurationKeys, types.KeyValueInput{
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

	var entries []types.AuthorizationDataInput
	for i := 0; i < 25; i++ {
		entries = append(entries, types.AuthorizationDataInput{IdTag: "TAG-1"})
	}

	input := sll.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: entries,
		UpdateType:             types.UpdateTypeFull.String(),
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
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	profileInput := types.ChargingProfileInput{
		ChargingProfileId:      1,
		TransactionId:          nil,
		StackLevel:             0,
		ChargingProfilePurpose: types.TxProfile.String(),
		ChargingProfileKind:    types.ChargingProfileKindAbsolute.String(),
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

	requestedMessage := types.MessageTriggerHeartbeat.String()
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
