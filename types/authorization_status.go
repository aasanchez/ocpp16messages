package types

// AuthorizationStatus defines the set of possible outcomes from an authorization attempt
// of an EV driver's identifier (idTag) in an OCPP 1.6J-based charging infrastructure.
//
// This type corresponds to the `status` field in the `Authorize.conf` message defined
// in the Open Charge Point Protocol (OCPP) 1.6J specification. It is returned by the
// Central System (CSMS) in response to an `Authorize.req` sent by the Charge Point.
//
// Each value describes whether the identifier is currently authorized to initiate a
// charging session, and if not, why the request has been denied. This type is critical
// for ensuring secure and policy-compliant access control at public and private EV
// charging stations.
//
// Specification Reference:
// - OCPP 1.6J Specification, Section 5.2: Authorize.req / Authorize.conf
type AuthorizationStatus string

const (
	// Accepted indicates that the identifier has been successfully authorized by the CSMS.
	//
	// Charging may proceed. This is the only status that allows a charging transaction
	// to start under normal OCPP conditions.
	Accepted AuthorizationStatus = "Accepted"

	// Blocked indicates that the identifier has been blocked in the CSMS and is not allowed to charge.
	//
	// Reasons for blocking may include policy violations, fraud, overdue billing, or
	// explicit administrative actions.
	Blocked AuthorizationStatus = "Blocked"

	// Expired indicates that the identifier is known to the CSMS but its validity period has passed.
	//
	// This may occur if the idTag is associated with a subscription or membership
	// that has lapsed or was not renewed. Charging is denied.
	Expired AuthorizationStatus = "Expired"

	// Invalid indicates that the identifier is unknown to the CSMS and not recognized as a valid user.
	//
	// This may result from typing errors, unregistered tags, or using unauthorized tokens.
	// Charging is denied, and corrective action is required.
	Invalid AuthorizationStatus = "Invalid"

	// ConcurrentTx indicates that the identifier is already being used in an active transaction
	// and concurrent transactions are not permitted by the Charge Point or CSMS configuration.
	//
	// This status only applies to `StartTransaction.req` messages and not to general authorization
	// attempts. Charging is denied unless the existing transaction is completed first.
	ConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid returns true if the AuthorizationStatus is one of the valid predefined values
// accepted by the OCPP 1.6J specification.
//
// This method is typically used during decoding, validation, or integration testing
// to ensure that the incoming status value is recognized and supported by compliant
// software. Unknown or misspelled values are considered invalid.
//
// Example usage:
//
//	if !status.IsValid() {
//	    return fmt.Errorf("invalid authorization status: %s", status)
//	}
func (s AuthorizationStatus) IsValid() bool {
	switch s {
	case Accepted, Blocked, Expired, Invalid, ConcurrentTx:
		return true
	default:
		return false
	}
}
