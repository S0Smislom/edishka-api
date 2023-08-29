package middleware

import (
	"food/pkg/response"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			// "request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &response.ResponseWriter{ResponseWriter: w, Code: http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.Code >= 500:
			level = logrus.ErrorLevel
		case rw.Code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.Code,
			http.StatusText(rw.Code),
			time.Now().Sub(start),
		)
	})
}
