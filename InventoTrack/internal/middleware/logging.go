package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs the details of each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}