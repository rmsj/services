// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/rmsj/services/business/auth"
	"github.com/rmsj/services/business/mid"
	"github.com/rmsj/services/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics(log))
	check := check{
		log: log,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness, mid.Authenticate(a))

	return app
}
