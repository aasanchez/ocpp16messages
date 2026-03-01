//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	gt "github.com/aasanchez/ocpp16messages/getConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	maxFuzzInputLen = 1024
)

func FuzzNewKeyValue(f *testing.F) {
	f.Add("", false, "", false)
	f.Add("TestKey", true, "TestValue", true)
	f.Add("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false, "", false)  // 50 chars
	f.Add("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", false, "", false) // 51 chars
	f.Add("\n", false, "", false)
	f.Add("é", false, "", false)

	f.Fuzz(func(
		t *testing.T,
		key string,
		readonly bool,
		value string,
		hasValue bool,
	) {
		if len(key) > maxFuzzInputLen || len(value) > maxFuzzInputLen {
			t.Skip()
		}

		var valuePtr *string
		if hasValue {
			valuePtr = &value
		}

		keyValue, err := gt.NewKeyValue(
			gt.KeyValueInput{
				Key:      key,
				Readonly: readonly,
				Value:    valuePtr,
			},
		)
		if err != nil {
			if !errors.Is(err, st.ErrEmptyValue) &&
				!errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if key == "" {
			t.Fatal("error = nil, want error for empty key")
		}

		if len(key) > st.CiString50Max {
			t.Fatalf(
				"len(key) = %d, want <= %d",
				len(key),
				st.CiString50Max,
			)
		}

		for _, r := range key {
			if r < 32 || r > 126 {
				t.Fatalf("key contains non-printable ASCII rune %q", r)
			}
		}

		if got := keyValue.Key().String(); got != key {
			t.Fatalf("Key().String() = %q, want %q", got, key)
		}

		if keyValue.Readonly() != readonly {
			t.Fatalf("Readonly() = %v, want %v", keyValue.Readonly(), readonly)
		}

		if !hasValue {
			if keyValue.Value() != nil {
				t.Fatal("Value() != nil, want nil for omitted value")
			}

			return
		}

		if value == "" {
			t.Fatal("error = nil, want error for empty value when provided")
		}

		if keyValue.Value() == nil {
			t.Fatal("Value() = nil, want non-nil when value is provided")
		}

		if got := keyValue.Value().String(); got != value {
			t.Fatalf("Value().String() = %q, want %q", got, value)
		}
	})
}
