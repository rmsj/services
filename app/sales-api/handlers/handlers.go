// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rmsj/services/business/auth"
	"github.com/rmsj/services/business/mid"
	"github.com/rmsj/services/foundation/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, a *auth.Auth, db *sqlx.DB) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics(log))
	check := checkGroup{
		build: build,
		db:    db,
	}
	app.Handle(http.MethodGet, "/debug/readiness", check.readiness)
	app.Handle(http.MethodGet, "/debug/liveness", check.liveness)

	return app
}
