package sharedtypes_test

import (
	"math"
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func assertOracleConsistency(t *testing.T, s string, err error, oracleErr error) {
	t.Helper()

	if (err == nil) != (oracleErr == nil) {
		t.Fatalf("mismatch vs oracle for input %q: err=%v, oracle err=%v", s, err, oracleErr)
	}
}

func assertZeroOnError(t *testing.T, got st.Integer, s string) {
	t.Helper()

	if got.Value() != 0 {
		t.Fatalf("non-zero value on error for %q: %d", s, got.Value())
	}
}

func assertEqualsOracle(t *testing.T, s string, v uint16, oracle uint64) {
	t.Helper()

	if oracle > uint64(math.MaxUint16) {
		t.Fatalf("oracle value out of range for uint16: %d", oracle)
	}

	if uint64(v) != oracle {
		t.Fatalf("value mismatch for %q: got=%d, want=%d", s, v, oracle)
	}
}

func assertRoundTrip(t *testing.T, v uint16) {
	t.Helper()

	rt := strconv.FormatUint(uint64(v), 10)

	back, berr := st.SetInteger(rt)
	if berr != nil {
		t.Fatalf("round-trip parse failed for %q: %v", rt, berr)
	}

	if back.Value() != v {
		t.Fatalf("round-trip value mismatch for %q: got=%d, want=%d", rt, back.Value(), v)
	}
}

func assertSuccessLengthInvariant(t *testing.T, s string, v uint16) {
	t.Helper()

	if len(s) > 5 && s[0] != '0' {
		t.Fatalf("parsed success from >5-digit non-zero-leading input: %q -> %d", s, v)
	}
}

func assertDigitsOnlyOnSuccess(t *testing.T, s string) {
	t.Helper()

	for i := range len(s) {
		if s[i] < '0' || s[i] > '9' {
			t.Fatalf("non-digit character accepted in successful parse: %q at %d in %q", s[i], i, s)
		}
	}
}

func assertNonEmptyOnSuccess(t *testing.T, s string) {
	t.Helper()

	if len(s) == 0 {
		t.Fatal("empty string parsed successfully")
	}
}

func FuzzSetInteger(f *testing.F) {
	for _, seed := range []string{
		"0", "1", "9", "10",
		"65535",      // max uint16
		"65536",      // overflow
		"00042",      // leading zeros
		"",           // empty
		"-1",         // negative
		"+1",         // plus sign
		" 1",         // space prefix
		"1 ",         // space suffix
		"foo",        // non-digits
		"4294967295", // large
		"100000",     // 6 digits, overflow (no leading zeros)
	} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 64 {
			return
		}

		got, err := st.SetInteger(s)
		oracle, oerr := strconv.ParseUint(s, 10, 16)

		assertOracleConsistency(t, s, err, oerr)

		if err != nil {
			assertZeroOnError(t, got, s)

			return
		}

		v := got.Value()
		assertEqualsOracle(t, s, v, oracle)
		assertRoundTrip(t, v)
		assertSuccessLengthInvariant(t, s, v)
		assertDigitsOnlyOnSuccess(t, s)
		assertNonEmptyOnSuccess(t, s)
	})
}
