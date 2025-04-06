package errorTypes

// ErrorCode represents predefined OCPP 1.6J error codes for CALLERROR messages.
type ErrorCode string

const (
	ErrorNotSupported                  ErrorCode = "NotSupported"
	ErrorInternalError                 ErrorCode = "InternalError"
	ErrorProtocolError                 ErrorCode = "ProtocolError"
	ErrorSecurityError                 ErrorCode = "SecurityError"
	ErrorFormationViolation            ErrorCode = "FormationViolation"
	ErrorPropertyConstraintViolation   ErrorCode = "PropertyConstraintViolation"
	ErrorOccurrenceConstraintViolation ErrorCode = "OccurrenceConstraintViolation"
	ErrorTypeConstraintViolation       ErrorCode = "TypeConstraintViolation"
	ErrorGenericError                  ErrorCode = "GenericError"

	// You can define library-specific error codes prefixed accordingly
	ErrorLibraryValidationFailed ErrorCode = "LibraryValidationFailed"
	ErrorLibraryTimeout          ErrorCode = "LibraryTimeout"
)

func (e ErrorCode) String() string {
	return string(e)
}

func (e ErrorCode) IsValid() bool {
	switch e {
	case ErrorNotSupported, ErrorInternalError, ErrorProtocolError,
		ErrorSecurityError, ErrorFormationViolation, ErrorPropertyConstraintViolation,
		ErrorOccurrenceConstraintViolation, ErrorTypeConstraintViolation,
		ErrorGenericError,
		ErrorLibraryValidationFailed, ErrorLibraryTimeout:
		return true
	default:
		return false
	}
}
