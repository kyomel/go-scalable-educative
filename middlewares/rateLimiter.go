package middlewares

import (
	"net/http"
	"time"
)

const (
	MAX_REQUESTS      = 10
	MAX_WAIT          = 2 * time.Second
	TOO_MANY_REQUESTS = "Too many requests!"
	REQUEST_CANCELED  = "Request canceled"
)

var semaphore = make(chan bool, MAX_REQUESTS)

func WithRateLimit(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := time.NewTimer(MAX_WAIT)
		ctx := r.Context()
		select {
		case <-timer.C:
			http.Error(w, TOO_MANY_REQUESTS, http.StatusTooManyRequests)
			return
		case <-ctx.Done():
			timer.Stop()
			http.Error(w, REQUEST_CANCELED, http.StatusTooManyRequests)
			return
		case semaphore <- true:
			timer.Stop()
			defer func() {
				timer.Stop()
				<-semaphore
			}()
			handler.ServeHTTP(w, r)
			return
		}
	})
}
