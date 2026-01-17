package types_test

import (
	"testing"

	mt "github.com/aasanchez/ocpp16messages/meterValues/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

const (
	uomWhStr         = "Wh"
	uomKWhStr        = "kWh"
	uomVarhStr       = "varh"
	uomKvarhStr      = "kvarh"
	uomWStr          = "W"
	uomKWStr         = "kW"
	uomVAStr         = "VA"
	uomKVAStr        = "kVA"
	uomVarStr        = "var"
	uomKvarStr       = "kvar"
	uomAStr          = "A"
	uomVStr          = "V"
	uomCelsiusStr    = "Celsius"
	uomFahrenheitStr = "Fahrenheit"
	uomKStr          = "K"
	uomPercentStr    = "Percent"
	uomTypeStr       = "UnitOfMeasure.String()"
)

func TestUnitOfMeasure_IsValid_Wh(t *testing.T) {
	t.Parallel()

	if !mt.Wh.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Wh")
	}
}

func TestUnitOfMeasure_IsValid_KWh(t *testing.T) {
	t.Parallel()

	if !mt.KWh.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "KWh")
	}
}

func TestUnitOfMeasure_IsValid_Varh(t *testing.T) {
	t.Parallel()

	if !mt.Varh.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Varh")
	}
}

func TestUnitOfMeasure_IsValid_Kvarh(t *testing.T) {
	t.Parallel()

	if !mt.Kvarh.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Kvarh")
	}
}

func TestUnitOfMeasure_IsValid_W(t *testing.T) {
	t.Parallel()

	if !mt.W.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "W")
	}
}

func TestUnitOfMeasure_IsValid_KW(t *testing.T) {
	t.Parallel()

	if !mt.KW.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "KW")
	}
}

func TestUnitOfMeasure_IsValid_VA(t *testing.T) {
	t.Parallel()

	if !mt.VA.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "VA")
	}
}

func TestUnitOfMeasure_IsValid_KVA(t *testing.T) {
	t.Parallel()

	if !mt.KVA.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "KVA")
	}
}

func TestUnitOfMeasure_IsValid_Var(t *testing.T) {
	t.Parallel()

	if !mt.Var.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Var")
	}
}

func TestUnitOfMeasure_IsValid_Kvar(t *testing.T) {
	t.Parallel()

	if !mt.Kvar.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Kvar")
	}
}

func TestUnitOfMeasure_IsValid_A(t *testing.T) {
	t.Parallel()

	if !mt.A.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "A")
	}
}

func TestUnitOfMeasure_IsValid_V(t *testing.T) {
	t.Parallel()

	if !mt.V.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "V")
	}
}

func TestUnitOfMeasure_IsValid_Celsius(t *testing.T) {
	t.Parallel()

	if !mt.Celsius.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Celsius")
	}
}

func TestUnitOfMeasure_IsValid_Fahrenheit(t *testing.T) {
	t.Parallel()

	if !mt.Fahrenheit.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Fahrenheit")
	}
}

func TestUnitOfMeasure_IsValid_K(t *testing.T) {
	t.Parallel()

	if !mt.K.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "K")
	}
}

func TestUnitOfMeasure_IsValid_Percent(t *testing.T) {
	t.Parallel()

	if !mt.Percent.IsValid() {
		t.Errorf(st.ErrorIsValidFalse, "Percent")
	}
}

func TestUnitOfMeasure_IsValid_Empty(t *testing.T) {
	t.Parallel()

	uom := mt.UnitOfMeasure("")
	if uom.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "UnitOfMeasure(\"\")")
	}
}

func TestUnitOfMeasure_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	uom := mt.UnitOfMeasure("Invalid")
	if uom.IsValid() {
		t.Errorf(st.ErrorIsValidTrue, "UnitOfMeasure(\"Invalid\")")
	}
}

func TestUnitOfMeasure_String_Wh(t *testing.T) {
	t.Parallel()

	got := mt.Wh.String()
	if got != uomWhStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomWhStr)
	}
}

func TestUnitOfMeasure_String_KWh(t *testing.T) {
	t.Parallel()

	got := mt.KWh.String()
	if got != uomKWhStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomKWhStr)
	}
}

func TestUnitOfMeasure_String_Varh(t *testing.T) {
	t.Parallel()

	got := mt.Varh.String()
	if got != uomVarhStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomVarhStr)
	}
}

func TestUnitOfMeasure_String_Kvarh(t *testing.T) {
	t.Parallel()

	got := mt.Kvarh.String()
	if got != uomKvarhStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomKvarhStr)
	}
}

func TestUnitOfMeasure_String_W(t *testing.T) {
	t.Parallel()

	got := mt.W.String()
	if got != uomWStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomWStr)
	}
}

func TestUnitOfMeasure_String_KW(t *testing.T) {
	t.Parallel()

	got := mt.KW.String()
	if got != uomKWStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomKWStr)
	}
}

func TestUnitOfMeasure_String_VA(t *testing.T) {
	t.Parallel()

	got := mt.VA.String()
	if got != uomVAStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomVAStr)
	}
}

func TestUnitOfMeasure_String_KVA(t *testing.T) {
	t.Parallel()

	got := mt.KVA.String()
	if got != uomKVAStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomKVAStr)
	}
}

func TestUnitOfMeasure_String_Var(t *testing.T) {
	t.Parallel()

	got := mt.Var.String()
	if got != uomVarStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomVarStr)
	}
}

func TestUnitOfMeasure_String_Kvar(t *testing.T) {
	t.Parallel()

	got := mt.Kvar.String()
	if got != uomKvarStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomKvarStr)
	}
}

func TestUnitOfMeasure_String_A(t *testing.T) {
	t.Parallel()

	got := mt.A.String()
	if got != uomAStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomAStr)
	}
}

func TestUnitOfMeasure_String_V(t *testing.T) {
	t.Parallel()

	got := mt.V.String()
	if got != uomVStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomVStr)
	}
}

func TestUnitOfMeasure_String_Celsius(t *testing.T) {
	t.Parallel()

	got := mt.Celsius.String()
	if got != uomCelsiusStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomCelsiusStr)
	}
}

func TestUnitOfMeasure_String_Fahrenheit(t *testing.T) {
	t.Parallel()

	got := mt.Fahrenheit.String()
	if got != uomFahrenheitStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomFahrenheitStr)
	}
}

func TestUnitOfMeasure_String_K(t *testing.T) {
	t.Parallel()

	got := mt.K.String()
	if got != uomKStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomKStr)
	}
}

func TestUnitOfMeasure_String_Percent(t *testing.T) {
	t.Parallel()

	got := mt.Percent.String()
	if got != uomPercentStr {
		t.Errorf(st.ErrorMethodMismatch, uomTypeStr, got, uomPercentStr)
	}
}
