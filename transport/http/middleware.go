package http

import "net/http"

// Middleware is a function type alias representing a single item in a
// middleware stack.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain returns a function taking a HandlerFunc, likely to be the actual
// handler for a server request, that itself also returns an HandlerFunc. This
// function iterates through a varaidc list of HandlerFunc, representing the
// middleware stack, passing to each a reference to the previous function.
func Chain(fns ...Middleware) Middleware {
	return func(route http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handler := route
			for _, fn := range fns {
				handler = fn(handler)
			}

			handler(w, r)
		}
	}
}
