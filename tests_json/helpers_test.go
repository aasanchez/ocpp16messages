package tests_json_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/aasanchez/ocpp16messages/types"
)

var (
	dateTimeType = reflect.TypeOf(types.DateTime{})
)

func roundTripJSON[T any](t *testing.T, value T) T {
	t.Helper()

	originalJSON, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("json.Marshal(original): %v", err)
	}

	var decoded T
	if err := json.Unmarshal(originalJSON, &decoded); err != nil {
		t.Fatalf("json.Unmarshal: %v (json=%s)", err, string(originalJSON))
	}

	decodedJSON, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("json.Marshal(decoded): %v", err)
	}

	assertJSONSemanticallyEqual(t, originalJSON, decodedJSON)

	return decoded
}

func assertJSONSemanticallyEqual(t *testing.T, left, right []byte) {
	t.Helper()

	var leftValue any
	if err := json.Unmarshal(left, &leftValue); err != nil {
		t.Fatalf("json.Unmarshal(left): %v", err)
	}

	var rightValue any
	if err := json.Unmarshal(right, &rightValue); err != nil {
		t.Fatalf("json.Unmarshal(right): %v", err)
	}

	if !reflect.DeepEqual(leftValue, rightValue) {
		t.Fatalf(
			"json mismatch:\nleft:  %s\nright: %s",
			string(left),
			string(right),
		)
	}
}

func assertAllFieldsValid(t *testing.T, value any) {
	t.Helper()

	visitedPointers := map[uintptr]bool{}
	assertAllFieldsValidValue(t, reflect.ValueOf(value), visitedPointers)
}

func assertAllFieldsValidValue(
	t *testing.T,
	value reflect.Value,
	visitedPointers map[uintptr]bool,
) {
	t.Helper()

	if !value.IsValid() {
		return
	}

	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			return
		}
		assertAllFieldsValidValue(t, value.Elem(), visitedPointers)
		return
	}

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return
		}
		ptr := value.Pointer()
		if ptr != 0 {
			if visitedPointers[ptr] {
				return
			}
			visitedPointers[ptr] = true
		}
		assertAllFieldsValidValue(t, value.Elem(), visitedPointers)
		return
	}

	if value.Type() == dateTimeType {
		dateTime := value.Interface().(types.DateTime)
		dateTimeString := dateTime.String()
		if !strings.HasSuffix(dateTimeString, "Z") {
			t.Fatalf("DateTime not UTC: %q", dateTimeString)
		}
		parsed, err := time.Parse(time.RFC3339Nano, dateTimeString)
		if err != nil {
			t.Fatalf("DateTime not parseable: %q: %v", dateTimeString, err)
		}
		if parsed.Location() != time.UTC {
			t.Fatalf("DateTime not UTC location: %q", dateTimeString)
		}
		return
	}

	if value.CanInterface() {
		if validator, ok := value.Interface().(interface{ IsValid() bool }); ok {
			if !validator.IsValid() {
				t.Fatalf("IsValid() = false for %T", value.Interface())
			}
		}
	}

	switch value.Kind() {
	case reflect.Struct:
		for fieldIndex := 0; fieldIndex < value.NumField(); fieldIndex++ {
			field := value.Field(fieldIndex)
			assertAllFieldsValidValue(t, field, visitedPointers)
		}
	case reflect.Slice, reflect.Array:
		for index := 0; index < value.Len(); index++ {
			assertAllFieldsValidValue(t, value.Index(index), visitedPointers)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			assertAllFieldsValidValue(t, value.MapIndex(key), visitedPointers)
		}
	default:
	}
}
