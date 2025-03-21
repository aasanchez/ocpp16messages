package authorize

import (
	"fmt"

	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	st "github.com/aasanchez/ocpp16messages/shared/types"
)

type RequestMessage struct {
	idTag mat.IdToken
}

func Request(input mat.RequestPayload) (RequestMessage, error) {
	str, err := st.SetCiString20Type(input.IdTag())
	if err != nil {
		wrapped := fmt.Errorf(st.ErrorInvalid, "idTag", err)

		return RequestMessage{}, wrapped
	}

	idToken, _ := mat.SetIdToken(str)

	return RequestMessage{idTag: idToken}, nil
}
