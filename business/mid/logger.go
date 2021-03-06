package mid

import (
	"context"
	"github.com/rmsj/services/foundation/web"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"
)

// Logger writes some information about the request to the logs in the
// format: TraceID : (200) GET /foo -> IP ADDR (latency)
func Logger(log *log.Logger) web.Middleware {

	// we create this closure to be able to return what a Middleware is - receiving a handler, return a handler
	// and still be able, through the main function signature, receive a logger  which we need

	// This is the actual middleware function to be executed.
	m := func(handler web.Handler) web.Handler {

		// the idea of calling the parameter before/after/handler helps with readability - we know
		// the handler executes before of the logic present in this logger middleware

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.mid.logger")
			defer span.End()

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return web.NewShutdownError("web value missing from context")
			}

			log.Printf("%s: started   : %s %s -> %s",
				v.TraceID,
				r.Method, r.URL.Path, r.RemoteAddr,
			)

			// Call the next handler.
			err := handler(ctx, w, r)

			log.Printf("%s: completed : %s %s -> %s (%d) (%s)",
				v.TraceID,
				r.Method, r.URL.Path, r.RemoteAddr,
				v.StatusCode, time.Since(v.Now),
			)

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}

	return m
}
