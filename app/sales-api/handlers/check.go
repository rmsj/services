package handlers

import (
	"context"
	"errors"
	"github.com/rmsj/services/foundation/web"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// mimic error
	if n := rand.Intn(100); n%2 == 0 {
		//return errors.New("untrusted error")
		return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
		//return web.NewShutdownError("forcing shutdown for testing")
		//panic("forcing panic")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
