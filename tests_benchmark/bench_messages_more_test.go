//go:build bench

package benchmark

import (
	"testing"

	ccp "github.com/aasanchez/ocpp16messages/clearChargingProfile"
	gconf "github.com/aasanchez/ocpp16messages/getConfiguration"
	gd "github.com/aasanchez/ocpp16messages/getDiagnostics"
	rn "github.com/aasanchez/ocpp16messages/reserveNow"
	sll "github.com/aasanchez/ocpp16messages/sendLocalList"
	slt "github.com/aasanchez/ocpp16messages/sendLocalList/types"
	scp "github.com/aasanchez/ocpp16messages/setChargingProfile"
	scpt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	stt "github.com/aasanchez/ocpp16messages/startTransaction"
	sn "github.com/aasanchez/ocpp16messages/statusNotification"
	stp "github.com/aasanchez/ocpp16messages/stopTransaction"
	st "github.com/aasanchez/ocpp16messages/types"
	uf "github.com/aasanchez/ocpp16messages/updateFirmware"
)

func BenchmarkGetConfigurationReq_SingleKey(b *testing.B) {
	b.ReportAllocs()

	keys := []string{"HeartbeatInterval"}
	input := gconf.ReqInput{Key: keys}

	for i := 0; i < b.N; i++ {
		if _, err := gconf.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetConfigurationReq_25Keys(b *testing.B) {
	b.ReportAllocs()

	var keys []string
	for i := 0; i < 25; i++ {
		keys = append(keys, "HeartbeatInterval")
	}

	input := gconf.ReqInput{Key: keys}

	for i := 0; i < b.N; i++ {
		if _, err := gconf.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClearChargingProfileReq_AllFields(b *testing.B) {
	b.ReportAllocs()

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

	for i := 0; i < b.N; i++ {
		if _, err := ccp.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetDiagnosticsReq_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	retries := 3
	retryInterval := 60
	startTime := sampleTimestamp
	stopTime := "2025-01-02T16:00:00Z"

	input := gd.ReqInput{
		Location:      "https://example.com/upload",
		Retries:       &retries,
		RetryInterval: &retryInterval,
		StartTime:     &startTime,
		StopTime:      &stopTime,
	}

	for i := 0; i < b.N; i++ {
		if _, err := gd.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetDiagnosticsConf_WithFileName(b *testing.B) {
	b.ReportAllocs()

	fileName := "diagnostics.log"
	input := gd.ConfInput{FileName: &fileName}

	for i := 0; i < b.N; i++ {
		if _, err := gd.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUpdateFirmwareReq_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	retries := 3
	retryInterval := 60

	input := uf.ReqInput{
		Location:      "https://example.com/firmware.bin",
		RetrieveDate:  sampleTimestamp,
		Retries:       &retries,
		RetryInterval: &retryInterval,
	}

	for i := 0; i < b.N; i++ {
		if _, err := uf.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReserveNowReq_AllFields(b *testing.B) {
	b.ReportAllocs()

	parentIdTag := "PARENT-1"
	input := rn.ReqInput{
		ReservationId: 1,
		ConnectorId:   1,
		IdTag:         "TAG-1",
		ExpiryDate:    "2025-01-02T16:00:00Z",
		ParentIdTag:   &parentIdTag,
	}

	for i := 0; i < b.N; i++ {
		if _, err := rn.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStatusNotificationReq_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	info := "Info"
	timestamp := sampleTimestamp
	vendorId := "Vendor-1"
	vendorErrorCode := "E-1"

	input := sn.ReqInput{
		ConnectorId:     0,
		ErrorCode:       "NoError",
		Status:          "Available",
		Info:            &info,
		Timestamp:       &timestamp,
		VendorId:        &vendorId,
		VendorErrorCode: &vendorErrorCode,
	}

	for i := 0; i < b.N; i++ {
		if _, err := sn.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStartTransactionReq_WithReservationId(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := stt.ReqInput{
		ConnectorId:   1,
		IdTag:         "TAG-1",
		MeterStart:    100,
		Timestamp:     sampleTimestamp,
		ReservationId: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		if _, err := stt.Req(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStartTransactionConf_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	expiry := "2025-01-03T15:00:00Z"
	parent := "PARENT-1"

	input := stt.ConfInput{
		TransactionId: 1,
		Status:        "Accepted",
		ExpiryDate:    &expiry,
		ParentIdTag:   &parent,
	}

	for i := 0; i < b.N; i++ {
		if _, err := stt.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStopTransactionConf_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	status := "Accepted"
	expiry := "2025-01-03T15:00:00Z"
	parent := "PARENT-1"

	input := stp.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}

	for i := 0; i < b.N; i++ {
		if _, err := stp.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSendLocalListConf(b *testing.B) {
	b.ReportAllocs()

	input := sll.ConfInput{Status: slt.UpdateStatusAccepted.String()}

	for i := 0; i < b.N; i++ {
		if _, err := sll.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetChargingProfileConf(b *testing.B) {
	b.ReportAllocs()

	input := scp.ConfInput{Status: scpt.ChargingProfileStatusAccepted.String()}

	for i := 0; i < b.N; i++ {
		if _, err := scp.Conf(input); err != nil {
			b.Fatal(err)
		}
	}
}
