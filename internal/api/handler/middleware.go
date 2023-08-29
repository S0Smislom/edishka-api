package handler

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

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
