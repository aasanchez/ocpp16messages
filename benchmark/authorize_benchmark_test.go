package benchmark

import (
	"encoding/json"
	"testing"

	"github.com/aasanchez/ocpp16_messages/authorize"
	"github.com/aasanchez/ocpp16_messages/core"
)

var rawAuthorizeReq = []byte(`[2,"01221201194032","Authorize",{"idTag":"D0431F35"}]`)

func BenchmarkUnmarshalAuthorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var v any
		if err := json.Unmarshal(rawAuthorizeReq, &v); err != nil {
			b.Fatalf("unmarshal failed: %v", err)
		}
	}
}

func BenchmarkValidateMessage(b *testing.B) {
	msg, err := core.ValidateRawMessage(rawAuthorizeReq)
	if err != nil {
		b.Fatalf("failed to parse message: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := core.ValidateMessage(msg.Action, msg.Payload); err != nil {
			b.Fatalf("validation failed: %v", err)
		}
	}
}

func BenchmarkFieldLevelValidation(b *testing.B) {
	idToken, err := core.NewIdToken("D0431F35")
	if err != nil {
		b.Fatalf("failed to create idToken: %v", err)
	}
	validReq := authorize.AuthorizeReq{IdTag: idToken}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		payload, err := json.Marshal(validReq)
		if err != nil {
			b.Fatalf("marshal failed: %v", err)
		}
		if err := core.ValidateMessage("Authorize", payload); err != nil {
			b.Fatalf("validation failed: %v", err)
		}
	}
}
