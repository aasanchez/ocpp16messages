//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzIdTagInfoWithExpiryDate(f *testing.F) {
	f.Add("Accepted", "2025-01-15T10:30:00Z")
	f.Add("Blocked", "2025-12-31T23:59:59Z")
	f.Add("Expired", "invalid-time")

	f.Fuzz(func(t *testing.T, status, expiryDate string) {
		if len(status) > maxFuzzLen ||
			len(expiryDate) > maxFuzzLen {
			t.Skip()
		}

		info, err := types.NewIdTagInfo(
			types.AuthorizationStatus(status),
		)
		if err != nil {
			return
		}

		dt, err := types.NewDateTime(expiryDate)
		if err != nil {
			return
		}

		result := info.WithExpiryDate(dt)

		if result.Status != info.Status {
			t.Fatalf(
				"Status changed: got %q, want %q",
				result.Status, info.Status,
			)
		}

		if result.ExpiryDate == nil {
			t.Fatal("ExpiryDate = nil after WithExpiryDate")
		}

		if result.ExpiryDate.String() != dt.String() {
			t.Fatalf(
				"ExpiryDate = %q, want %q",
				result.ExpiryDate.String(), dt.String(),
			)
		}

		if result.ParentIdTag != nil {
			t.Fatal(
				"ParentIdTag != nil after WithExpiryDate only",
			)
		}

		first := result.String()
		second := result.String()

		if first != second {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				first, second,
			)
		}
	})
}

func FuzzIdTagInfoWithParentIdTag(f *testing.F) {
	f.Add("Accepted", "RFID-ABC123")
	f.Add("Invalid", "tag")
	f.Add("Blocked", "")

	f.Fuzz(func(t *testing.T, status, idTag string) {
		if len(status) > maxFuzzLen ||
			len(idTag) > maxFuzzLen {
			t.Skip()
		}

		info, err := types.NewIdTagInfo(
			types.AuthorizationStatus(status),
		)
		if err != nil {
			return
		}

		ciStr, err := types.NewCiString20Type(idTag)
		if err != nil {
			return
		}

		token := types.NewIdToken(ciStr)
		result := info.WithParentIdTag(token)

		if result.Status != info.Status {
			t.Fatalf(
				"Status changed: got %q, want %q",
				result.Status, info.Status,
			)
		}

		if result.ParentIdTag == nil {
			t.Fatal("ParentIdTag = nil after WithParentIdTag")
		}

		if result.ParentIdTag.String() != idTag {
			t.Fatalf(
				"ParentIdTag = %q, want %q",
				result.ParentIdTag.String(), idTag,
			)
		}

		if result.ExpiryDate != nil {
			t.Fatal(
				"ExpiryDate != nil after WithParentIdTag only",
			)
		}
	})
}

func FuzzIdTagInfoWithBothBuilders(f *testing.F) {
	f.Add("Accepted", "2025-01-15T10:30:00Z", "TAG1")
	f.Add("Blocked", "2025-06-01T00:00:00Z", "RFID-XYZ")

	f.Fuzz(func(
		t *testing.T,
		status, expiryDate, idTag string,
	) {
		if len(status) > maxFuzzLen ||
			len(expiryDate) > maxFuzzLen ||
			len(idTag) > maxFuzzLen {
			t.Skip()
		}

		info, err := types.NewIdTagInfo(
			types.AuthorizationStatus(status),
		)
		if err != nil {
			return
		}

		dt, err := types.NewDateTime(expiryDate)
		if err != nil {
			return
		}

		ciStr, err := types.NewCiString20Type(idTag)
		if err != nil {
			return
		}

		token := types.NewIdToken(ciStr)

		orderA := info.WithExpiryDate(dt).WithParentIdTag(token)
		orderB := info.WithParentIdTag(token).WithExpiryDate(dt)

		if orderA.Status != orderB.Status {
			t.Fatal("Status differs between builder orders")
		}

		if orderA.ExpiryDate == nil || orderB.ExpiryDate == nil {
			t.Fatal("ExpiryDate = nil in one order")
		}

		if orderA.ExpiryDate.String() !=
			orderB.ExpiryDate.String() {
			t.Fatal("ExpiryDate differs between orders")
		}

		if orderA.ParentIdTag == nil ||
			orderB.ParentIdTag == nil {
			t.Fatal("ParentIdTag = nil in one order")
		}

		if orderA.ParentIdTag.String() !=
			orderB.ParentIdTag.String() {
			t.Fatal("ParentIdTag differs between orders")
		}

		if orderA.String() != orderB.String() {
			t.Fatalf(
				"String() differs between orders: %q vs %q",
				orderA.String(), orderB.String(),
			)
		}
	})
}

func FuzzIdTagInfoStringDeterminism(f *testing.F) {
	f.Add("Accepted", false, "", false, "")
	f.Add("Blocked", true, "2025-01-15T10:30:00Z", true, "TAG1")
	f.Add("Expired", true, "2025-06-01T00:00:00Z", false, "")

	f.Fuzz(func(
		t *testing.T,
		status string,
		hasExpiry bool, expiryDate string,
		hasTag bool, idTag string,
	) {
		if len(status) > maxFuzzLen ||
			len(expiryDate) > maxFuzzLen ||
			len(idTag) > maxFuzzLen {
			t.Skip()
		}

		info, err := types.NewIdTagInfo(
			types.AuthorizationStatus(status),
		)
		if err != nil {
			return
		}

		if hasExpiry {
			dt, dtErr := types.NewDateTime(expiryDate)
			if dtErr != nil {
				if !errors.Is(
					dtErr, types.ErrInvalidValue,
				) && !errors.Is(
					dtErr, types.ErrEmptyValue,
				) {
					t.Fatalf(
						"DateTime error = %v, want sentinel",
						dtErr,
					)
				}

				return
			}

			info = info.WithExpiryDate(dt)
		}

		if hasTag {
			ciStr, ciErr := types.NewCiString20Type(idTag)
			if ciErr != nil {
				return
			}

			info = info.WithParentIdTag(types.NewIdToken(ciStr))
		}

		first := info.String()
		second := info.String()

		if first != second {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				first, second,
			)
		}
	})
}
