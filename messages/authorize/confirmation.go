package authorize

import (
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

// Response represents an OCPP 1.6 Authorize response message.
type Response struct {
	IdTagInfo mat.IdTagInfo
}

// NewResponse creates a new Authorize response with the given IdTagInfo.
func NewResponse(idTagInfo mat.IdTagInfo) Response {
	return Response{IdTagInfo: idTagInfo}
}
