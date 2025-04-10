package types

// IdTagType represents a validated authorization token (idTag) used in the
// Open Charge Point Protocol (OCPP) 1.6J specification.
//
// The `idTag` is a core element in OCPP used to identify and authorize users
// for charging sessions. It appears in multiple messages such as `Authorize.req`,
// `StartTransaction.req`, and `StopTransaction.req`, and is typically associated
// with a physical RFID card, mobile app identifier, or user account.
//
// This type wraps a validated `IdTokenType` and exposes methods to safely
// retrieve the underlying string or revalidate it against protocol constraints.
//
// Specification Reference:
// - OCPP 1.6J, Section 5.2: Authorize.req
// - OCPP 1.6J, Section 5.3: StartTransaction.req
type IdTagType struct {
	value IdTokenType
}

// IdTag constructs a new IdTagType instance from a raw string,
// validating it against the rules defined for `idTag` in OCPP 1.6J.
//
// The string must:
//   - Be non-empty.
//   - Be composed entirely of printable ASCII characters.
//   - Not exceed 20 characters in length.
//
// Returns an IdTagType instance on success, or an error on validation failure.
//
// Example:
//
//	tag, err := types.IdTag("ABC123XYZ")
//	if err != nil {
//	    log.Fatalf("Invalid idTag: %v", err)
//	}
func IdTag(s string) (IdTagType, error) {
	idToken, err := IdToken(s)
	if err != nil {
		return IdTagType{}, err
	}
	return IdTagType{value: idToken}, nil
}

// String returns the underlying string representation of the idTag.
//
// This method provides read-only access to the stored idTag value and
// is useful when constructing outbound OCPP messages or logs.
func (id IdTagType) String() string {
	return id.value.String()
}

// Validate revalidates the internal idTag value against all OCPP 1.6J constraints.
//
// Although construction already ensures correctness, this method can be used
// to verify the value after deserialization, transformation, or data import.
func (id IdTagType) Validate() error {
	return id.value.Validate()
}
