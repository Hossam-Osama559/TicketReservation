package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aarondl/authboss/v3"
	_ "github.com/aarondl/authboss/v3/auth"
	"github.com/aarondl/authboss/v3/defaults"
	_ "github.com/aarondl/authboss/v3/register"
)

// this package is responsible for handling the setup of authboss

var (
	AuthbossInstance *authboss.Authboss = authboss.New()
)

func Setup(storer authboss.ServerStorer, sessionhandler authboss.ClientStateReadWriter, renderer authboss.Renderer) {
	AuthbossInstance.Config.Paths.RootURL = "http://localhost:8080"
	AuthbossInstance.Config.Paths.Mount = "/"

	AuthbossInstance.Config.Storage.Server = storer

	AuthbossInstance.Config.Storage.SessionState = sessionhandler

	AuthbossInstance.Config.Core.ViewRenderer = renderer

	AuthbossInstance.Config.Core.BodyReader = defaults.NewHTTPBodyReader(false, false)

	AuthbossInstance.Config.Core.Responder = defaults.NewResponder(renderer)
	AuthbossInstance.Config.Core.Redirector = defaults.NewRedirector(renderer, "/auth/login")

	defaults.SetCore(&AuthbossInstance.Config, false, false)
	AuthbossInstance.Config.Core.ViewRenderer = renderer

	// 4. Initialize
	if err := AuthbossInstance.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Authboss Init Error: %v\n", err)
		os.Exit(1)
	}

}

func SetAuthMiddleWare(mux *http.ServeMux) http.Handler {

	return AuthbossInstance.LoadClientStateMiddleware(mux)
}
