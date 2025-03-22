package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

var rawAuthorizeReq = []byte(`[2,"01221201194032","Authorize",{"idTag":"D0431F35"}]`)

func BenchmarkValidateMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := core.ValidateRawMessage(rawAuthorizeReq)
		if err != nil {
			b.Fatalf("validation failed: %v", err)
		}
	}
}

func BenchmarkUnmarshalAuthorizeReq(b *testing.B) {
	var req authorize.Req
	payload := []byte(`{"idTag":"D0431F35"}`)
	for i := 0; i < b.N; i++ {
		if err := json.Unmarshal(payload, &req); err != nil {
			b.Fatalf("unmarshal failed: %v", err)
		}
	}
}

func BenchmarkFieldLevelValidation(b *testing.B) {
	idTag := authorize.IdToken{IdTag: "D0431F35"}
	for i := 0; i < b.N; i++ {
		if err := idTag.Validate(); err != nil {
			b.Fatalf("field validation failed: %v", err)
		}
	}
}
