package mid

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rmsj/services/foundation/web"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"runtime/debug"
)

// Panics recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panics(log *log.Logger) web.Middleware {

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

			ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.mid.panic")
			defer span.End()

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			// Defer a function to recover from a panic and set the err return
			// variable after the fact.
			// recover only works inside a defer
			// * The defer function execute after the return call, if a panic happens,
			// so we are using a named return and the defer function resets the return to be the panic
			defer func() {
				if r := recover(); r != nil {
					err = errors.Errorf("panic: %v", r)

					// Log the Go stack trace for this panic'd goroutine.
					log.Printf("%s: PANIC:\n%s", v.TraceID, debug.Stack())
				}
			}()

			// Call the next handler and set its return value in the err variable.
			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
