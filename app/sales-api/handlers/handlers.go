// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"github.com/rmsj/services/business/mid"
	"github.com/rmsj/services/foundation/web"
	"log"
	"net/http"
	"os"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics(log))
	check := check{
		log: log,
	}
	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
