package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/aasanchez/ocpp16_messages"
	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

var rawAuthorize = []byte(`[2, "uid123", "Authorize", {"idTag": "D0431F35"}]`)
var payload = authorize.AuthorizeReq{
	IdTag: "D0431F35",
}

func BenchmarkUnmarshalAuthorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var frame []json.RawMessage
		if err := json.Unmarshal(rawAuthorize, &frame); err != nil {
			b.Fatalf("frame unmarshal failed: %v", err)
		}

		if len(frame) < 4 {
			b.Fatalf("expected 4 elements in CALL message")
		}

		var out authorize.AuthorizeReq
		if err := json.Unmarshal(frame[3], &out); err != nil {
			b.Fatalf("payload unmarshal failed: %v", err)
		}
	}
}

func BenchmarkValidateMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ocpp16_messages.ValidateMessage(rawAuthorize); err != nil {
			b.Fatalf("validation failed: %v", err)
		}
	}
}

func BenchmarkFieldLevelValidation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := core.CiString20Type(payload.IdTag).Validate(); err != nil {
			b.Fatalf("field validation failed: %v", err)
		}
	}
}
