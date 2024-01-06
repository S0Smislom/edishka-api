package handler_test

// import (
// 	"encoding/json"
// 	"food/internal/api/handler"
// 	"food/internal/api/model"
// 	"food/internal/api/repository/postgres"
// 	"food/internal/api/service"
// 	"food/pkg/config"
// 	"food/pkg/response"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// func TestAuthMiddleware(t *testing.T) {
// 	config, _ := config.InitTestConfig()
// 	repo := postgres.NewRepository(nil)
// 	ser := service.NewService(repo, nil, config)

// 	h := handler.NewHandler(config, ser)
// 	s := h.AuthenticateUser(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))

// 	wrongAccessToken, _ := service.GenerateToken(model.AccessTokenType, &model.User{ID: 1, Phone: "+71111111111"}, 60, "testsignatureururur")

// 	headerOnlyBearer := "Bearer"
// 	headerIncorrectToken := "Bearer 123"
// 	headerWrongSignature := "Bearer " + wrongAccessToken
// 	headerEmpty := ""
// 	headerWithoutBearer := wrongAccessToken
// 	testCases := []struct {
// 		name         string
// 		headers      *string
// 		expectedData interface{}
// 		expectedCode int
// 	}{
// 		{
// 			name:    "Without token",
// 			headers: &headerOnlyBearer,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Unauthorized",
// 				Detail: "invalid auth header",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:    "Wrong access token",
// 			headers: &headerIncorrectToken,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Unauthorized",
// 				Detail: "invalid auth header",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:    "With wrong signature",
// 			headers: &headerWrongSignature,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Unauthorized",
// 				Detail: "invalid auth header",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:    "Without header",
// 			headers: nil,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Unauthorized",
// 				Detail: "invalid auth header",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:    "Empty header",
// 			headers: &headerEmpty,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Unauthorized",
// 				Detail: "invalid auth header",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:    "Without bearer",
// 			headers: &headerWithoutBearer,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Unauthorized",
// 				Detail: "invalid auth header",
// 			},
// 			expectedCode: http.StatusUnauthorized,
// 		},

// 		// {
// 		// 	name:    "With right token",
// 		// 	headers: "Bearer " + accessToken,
// 		// 	expectedData: &response.ErrorResponse{
// 		// 		Type:   "Permission Error",
// 		// 		Title:  "Unauthorized",
// 		// 		Detail: "invalid auth header",
// 		// 	},
// 		// 	expectedCode: http.StatusUnauthorized,
// 		// },
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
// 			if tc.headers != nil {
// 				req.Header.Add("Authorization", *tc.headers)
// 			}
// 			s.ServeHTTP(rec, req)
// 			if tc.expectedCode != rec.Code {
// 				t.Error("Wrong status code", rec.Code, tc.expectedCode)
// 			}
// 			response := strings.TrimRight(rec.Body.String(), "\n")
// 			expectedBytes, _ := json.Marshal(tc.expectedData)
// 			expectedString := string(expectedBytes)
// 			if response != expectedString {
// 				t.Error("Wrong data", response, expectedString, len(response), len(expectedString))
// 			}
// 		})
// 	}
// }
