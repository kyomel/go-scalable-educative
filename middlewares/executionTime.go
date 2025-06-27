package middlewares

import (
	"net/http"
	"time"

	"github.com/kyomel/go-scalable-educative/logger"
	"go.uber.org/zap"
)

func WithExecutionTime(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(w, r)
		defer logger.GetLoggerInstance().Info("Execution time", zap.Int64("microseconds", time.Since(startTime).Microseconds()))
	})
}
