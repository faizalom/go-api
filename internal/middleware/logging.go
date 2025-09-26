package middleware

import (
	"net/http"
	"time"

	"github.com/faizalom/go-api/pkg/logger" // Assumes 'workout-api' is your module name
)

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start the timer
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request details
		logger.Info.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
