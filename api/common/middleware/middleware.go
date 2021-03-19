package middleware

import "net/http"

// Handler used as middleware for http.Handler
type Handler func(http.Handler) http.Handler

// Apply middlewares on top of the base handler
// FILO - first in, last out
func Apply(h http.Handler, mws ...Handler) http.Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h
}
