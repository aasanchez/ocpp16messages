package types

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
)

func (e ErrorCode) String() string {
	return string(e)
}

func (e ErrorCode) IsValid() bool {
	switch e {
	case ErrorNotSupported, ErrorInternalError, ErrorProtocolError,
		ErrorSecurityError, ErrorFormationViolation, ErrorPropertyConstraintViolation,
		ErrorOccurrenceConstraintViolation, ErrorTypeConstraintViolation, ErrorGenericError:
		return true
	default:
		return false
	}
}
