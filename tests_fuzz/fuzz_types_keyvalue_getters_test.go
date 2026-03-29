//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzKeyValueGetterConsistency(f *testing.F) {
	f.Add("TestKey", true, true, "TestValue")
	f.Add("Key", false, false, "")
	f.Add("A", true, true, "V")

	f.Fuzz(func(
		t *testing.T,
		key string, readonly bool,
		hasValue bool, value string,
	) {
		if len(key) > maxFuzzLen ||
			len(value) > maxFuzzLen {
			t.Skip()
		}

		var valuePtr *string
		if hasValue {
			valuePtr = &value
		}

		kv, err := types.NewKeyValue(types.KeyValueInput{
			Key:      key,
			Readonly: readonly,
			Value:    valuePtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrEmptyValue) &&
				!errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want sentinel",
					err,
				)
			}

			return
		}

		// Key round-trip
		if kv.Key().String() != key {
			t.Fatalf(
				"Key() = %q, want %q",
				kv.Key().String(), key,
			)
		}

		// Key String() determinism
		k1 := kv.Key().String()
		k2 := kv.Key().String()

		if k1 != k2 {
			t.Fatal("Key().String() not deterministic")
		}

		// Readonly
		if kv.Readonly() != readonly {
			t.Fatalf(
				"Readonly() = %v, want %v",
				kv.Readonly(), readonly,
			)
		}

		// Value
		if hasValue {
			if kv.Value() == nil {
				t.Fatal("Value() = nil, want non-nil")
			}

			if kv.Value().String() != value {
				t.Fatalf(
					"Value() = %q, want %q",
					kv.Value().String(), value,
				)
			}
		} else if kv.Value() != nil {
			t.Fatal("Value() != nil, want nil")
		}
	})
}

func FuzzKeyValuePointerIsolation(f *testing.F) {
	f.Add("TestKey", "TestValue")

	f.Fuzz(func(t *testing.T, key, value string) {
		if len(key) > maxFuzzLen ||
			len(value) > maxFuzzLen {
			t.Skip()
		}

		valuePtr := &value

		kv, err := types.NewKeyValue(types.KeyValueInput{
			Key:      key,
			Readonly: true,
			Value:    valuePtr,
		})
		if err != nil {
			return
		}

		if kv.Value() == nil {
			return
		}

		// Get Value() twice and verify consistency
		v1 := kv.Value().String()
		v2 := kv.Value().String()

		if v1 != v2 {
			t.Fatal("Value() changed between calls")
		}
	})
}
