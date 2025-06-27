package middlewares

import (
	"net/http"

	"github.com/kyomel/go-scalable-educative/logger"
	"go.uber.org/zap"
)

func WithSimpleLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.GetLoggerInstance().Info("Incoming traffic", zap.String("path", r.URL.Path))
		handler.ServeHTTP(w, r)
	})
}
