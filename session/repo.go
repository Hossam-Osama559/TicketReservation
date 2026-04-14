package session

import (
	"net/http"

	"github.com/aarondl/authboss/v3"
)

type SessionRepository interface {
	ReadState(*http.Request) (authboss.ClientState, error)

	WriteState(http.ResponseWriter, authboss.ClientState, []authboss.ClientStateEvent) error
}
