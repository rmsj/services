package web

// Middleware is a function designed to run some code before and/or after
// another Handler. It is designed to remove boilerplate or other concerns not
// direct to any given Handler.
type Middleware func(Handler) Handler

// We will need, thinking about the onion with the proper handler at the center, and come out of the onion from center to the out
// middleware in this order: authentication, panic handling, metrics, rror, logging

// wrapMiddleware creates a new handler by wrapping middleware around a final
// handler. The middlewares' Handlers will be executed by requests in the order
// they are provided.
func wrapMiddleware(mw []Middleware, handler Handler) Handler {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		// h is of type Middleware func(Handler) Handler - so it takes a handler and returns a handler
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}
