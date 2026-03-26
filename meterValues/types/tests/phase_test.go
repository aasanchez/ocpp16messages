package types_test

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/metervalues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	phL1Str   = "L1"
	phL2Str   = "L2"
	phL3Str   = "L3"
	phNStr    = "N"
	phL1NStr  = "L1-N"
	phL2NStr  = "L2-N"
	phL3NStr  = "L3-N"
	phL1L2Str = "L1-L2"
	phL2L3Str = "L2-L3"
	phL3L1Str = "L3-L1"
	phTypeStr = "Phase.String()"
)

func TestPhase_IsValid_L1(t *testing.T) {
	t.Parallel()

	if !mt.L1.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L1")
	}
}

func TestPhase_IsValid_L2(t *testing.T) {
	t.Parallel()

	if !mt.L2.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L2")
	}
}

func TestPhase_IsValid_L3(t *testing.T) {
	t.Parallel()

	if !mt.L3.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L3")
	}
}

func TestPhase_IsValid_N(t *testing.T) {
	t.Parallel()

	if !mt.N.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "N")
	}
}

func TestPhase_IsValid_L1N(t *testing.T) {
	t.Parallel()

	if !mt.L1N.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L1N")
	}
}

func TestPhase_IsValid_L2N(t *testing.T) {
	t.Parallel()

	if !mt.L2N.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L2N")
	}
}

func TestPhase_IsValid_L3N(t *testing.T) {
	t.Parallel()

	if !mt.L3N.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L3N")
	}
}

func TestPhase_IsValid_L1L2(t *testing.T) {
	t.Parallel()

	if !mt.L1L2.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L1L2")
	}
}

func TestPhase_IsValid_L2L3(t *testing.T) {
	t.Parallel()

	if !mt.L2L3.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L2L3")
	}
}

func TestPhase_IsValid_L3L1(t *testing.T) {
	t.Parallel()

	if !mt.L3L1.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "L3L1")
	}
}

func TestPhase_IsValid_Empty(t *testing.T) {
	t.Parallel()

	ph := mt.Phase("")
	if ph.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Phase(\"\")")
	}
}

func TestPhase_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	ph := mt.Phase("Invalid")
	if ph.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "Phase(\"Invalid\")")
	}
}

func TestPhase_String_L1(t *testing.T) {
	t.Parallel()

	got := mt.L1.String()
	if got != phL1Str {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL1Str)
	}
}

func TestPhase_String_L2(t *testing.T) {
	t.Parallel()

	got := mt.L2.String()
	if got != phL2Str {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL2Str)
	}
}

func TestPhase_String_L3(t *testing.T) {
	t.Parallel()

	got := mt.L3.String()
	if got != phL3Str {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL3Str)
	}
}

func TestPhase_String_N(t *testing.T) {
	t.Parallel()

	got := mt.N.String()
	if got != phNStr {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phNStr)
	}
}

func TestPhase_String_L1N(t *testing.T) {
	t.Parallel()

	got := mt.L1N.String()
	if got != phL1NStr {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL1NStr)
	}
}

func TestPhase_String_L2N(t *testing.T) {
	t.Parallel()

	got := mt.L2N.String()
	if got != phL2NStr {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL2NStr)
	}
}

func TestPhase_String_L3N(t *testing.T) {
	t.Parallel()

	got := mt.L3N.String()
	if got != phL3NStr {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL3NStr)
	}
}

func TestPhase_String_L1L2(t *testing.T) {
	t.Parallel()

	got := mt.L1L2.String()
	if got != phL1L2Str {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL1L2Str)
	}
}

func TestPhase_String_L2L3(t *testing.T) {
	t.Parallel()

	got := mt.L2L3.String()
	if got != phL2L3Str {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL2L3Str)
	}
}

func TestPhase_String_L3L1(t *testing.T) {
	t.Parallel()

	got := mt.L3L1.String()
	if got != phL3L1Str {
		t.Errorf(st.ErrorMethodMismatch, phTypeStr, got, phL3L1Str)
	}
}
