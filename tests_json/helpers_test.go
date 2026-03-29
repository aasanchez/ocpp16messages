package testsjson_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

const (
	startIndex = 0
)

func roundTripJSON(t *testing.T, value any) {
	t.Helper()

	originalJSON, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("json.Marshal(original): %v", err)
	}

	decoded, err := unmarshalSameType(value, originalJSON)
	if err != nil {
		t.Fatalf("json.Unmarshal: %v (json=%s)", err, string(originalJSON))
	}

	assertAllFieldsValid(t, decoded)

	decodedJSON, err := json.Marshal(decoded)
	if err != nil {
		t.Fatalf("json.Marshal(decoded): %v", err)
	}

	assertJSONSemanticallyEqual(t, originalJSON, decodedJSON)
}

func assertJSONSemanticallyEqual(t *testing.T, left, right []byte) {
	t.Helper()

	var leftValue any

	err := json.Unmarshal(left, &leftValue)
	if err != nil {
		t.Fatalf("json.Unmarshal(left): %v", err)
	}

	var rightValue any

	err = json.Unmarshal(right, &rightValue)
	if err != nil {
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

	visitedPointers := map[uintptr]struct{}{}
	dateTimeType := reflect.TypeFor[types.DateTime]()

	assertAllFieldsValidValue(
		t,
		reflect.ValueOf(value),
		visitedPointers,
		dateTimeType,
	)
}

func assertAllFieldsValidValue(
	t *testing.T,
	value reflect.Value,
	visitedPointers map[uintptr]struct{},
	dateTimeType reflect.Type,
) {
	t.Helper()

	value, ok := derefValue(value, visitedPointers)
	if !ok {
		return
	}

	if validateDateTime(t, value, dateTimeType) {
		return
	}

	validateIsValid(t, value)

	assertAllChildrenValid(t, value, visitedPointers, dateTimeType)
}

func unmarshalSameType(source any, rawJSON []byte) (any, error) {
	sourceValue := reflect.ValueOf(source)
	sourceType := sourceValue.Type()

	decodeTarget := reflect.New(sourceType)

	err := json.Unmarshal(rawJSON, decodeTarget.Interface())
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return decodeTarget.Elem().Interface(), nil
}

//nolint:revive // Complexity is acceptable for a test helper.
func derefValue(
	value reflect.Value,
	visitedPointers map[uintptr]struct{},
) (reflect.Value, bool) {
	for {
		dereferencedValue, didDereference := derefInterfaceValue(value)
		if !didDereference {
			return reflect.Value{}, false
		}

		value = dereferencedValue

		dereferencedValue, didDereference = derefPointerValue(
			value,
			visitedPointers,
		)
		if !didDereference {
			return reflect.Value{}, false
		}

		value = dereferencedValue

		if value.Kind() != reflect.Interface && value.Kind() != reflect.Ptr {
			return value, true
		}
	}
}

func derefInterfaceValue(value reflect.Value) (reflect.Value, bool) {
	if !value.IsValid() {
		return reflect.Value{}, false
	}

	if value.Kind() != reflect.Interface {
		return value, true
	}

	if value.IsNil() {
		return reflect.Value{}, false
	}

	return value.Elem(), true
}

func derefPointerValue(
	value reflect.Value,
	visitedPointers map[uintptr]struct{},
) (reflect.Value, bool) {
	if !value.IsValid() {
		return reflect.Value{}, false
	}

	if value.Kind() != reflect.Ptr {
		return value, true
	}

	if value.IsNil() {
		return reflect.Value{}, false
	}

	ptr := value.Pointer()
	if ptr != uintptr(startIndex) {
		if _, ok := visitedPointers[ptr]; ok {
			return reflect.Value{}, false
		}

		visitedPointers[ptr] = struct{}{}
	}

	return value.Elem(), true
}

func validateDateTime(
	t *testing.T,
	value reflect.Value,
	dateTimeType reflect.Type,
) bool {
	t.Helper()

	if value.Type() != dateTimeType {
		return false
	}

	dateTime, ok := value.Interface().(types.DateTime)
	if !ok {
		t.Fatalf("expected types.DateTime, got %T", value.Interface())
	}

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

	return true
}

func validateIsValid(t *testing.T, value reflect.Value) {
	t.Helper()

	if !value.CanInterface() {
		return
	}

	type isValidInterface interface {
		IsValid() bool
	}

	validator, ok := value.Interface().(isValidInterface)
	if !ok {
		return
	}

	if !validator.IsValid() {
		t.Fatalf("IsValid() = false for %T", value.Interface())
	}
}

func assertAllChildrenValid(
	t *testing.T,
	value reflect.Value,
	visitedPointers map[uintptr]struct{},
	dateTimeType reflect.Type,
) {
	t.Helper()

	//nolint:exhaustive
	switch value.Kind() {
	case reflect.Struct:
		fieldCount := value.NumField()

		for fieldIndex := range fieldCount {
			field := value.Field(fieldIndex)
			assertAllFieldsValidValue(t, field, visitedPointers, dateTimeType)
		}

	case reflect.Slice, reflect.Array:
		for index := startIndex; index < value.Len(); index++ {
			assertAllFieldsValidValue(
				t,
				value.Index(index),
				visitedPointers,
				dateTimeType,
			)
		}

	case reflect.Map:
		for _, key := range value.MapKeys() {
			assertAllFieldsValidValue(
				t,
				value.MapIndex(key),
				visitedPointers,
				dateTimeType,
			)
		}

	default:
	}
}
