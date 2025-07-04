package api

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/kyomel/go-scalable-educative/logger"
	"github.com/kyomel/go-scalable-educative/network"
	"go.uber.org/zap"
)

func getTimedContext(timeout int) *context.Context {
	ctx, cancel := context.WithCancel(context.TODO())
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		cancel()
	})
	return &ctx
}

func apiTimingLogger(method, url string) func() {
	startTime := time.Now()
	return func() {
		duration := time.Since(startTime).Milliseconds()
		logger.GetLoggerInstance().Info(
			"API timing",
			zap.String("Method", method),
			zap.String("URL", url),
			zap.Int64("API duration", duration),
		)
	}
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	endTiming := apiTimingLogger("GET", "/users")
	res, err := network.NewClient().
		Name("fakerapi").
		Timeout(10).
		WithContext(getTimedContext(1)).
		Get("https://fakerapi.it/api/v1/persons")
	endTiming()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resData, _ := io.ReadAll(res.Body)
	w.Write(resData)
}
