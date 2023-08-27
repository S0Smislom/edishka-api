package handler

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (s *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			// "request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, Code: http.StatusOK}
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

// func (s *Handler) authenticateUser(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		header := r.Header.Get(authorizationHeader)
// 		if header == "" {
// 			s.errorRespond(w, r, http.StatusInternalServerError, errors.New("empty auth header"))
// 			return
// 		}
// 		headerParts := strings.Split(header, " ")
// 		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
// 			s.errorRespond(w, r, http.StatusInternalServerError, errors.New("invalid auth header"))
// 			return
// 		}
// 		if len(headerParts[1]) == 0 {
// 			s.errorRespond(w, r, http.StatusInternalServerError, errors.New("token is empty"))
// 			return
// 		}
// 		userId, err := s.services.Authorization.ParseToken(headerParts[1])
// 		if err != nil {
// 			s.errorRespond(w, r, http.StatusInternalServerError, errors.New("invalid auth header"))
// 			return
// 		}
// 		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userCtx, userId)))
// 	})
// }
