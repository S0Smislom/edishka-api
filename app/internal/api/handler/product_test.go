package handler_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"food/internal/api/model"
// 	"food/internal/api/service"
// 	"food/internal/test"
// 	"food/pkg/database"
// 	"food/pkg/response"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// )

// func TestProductGetById(t *testing.T) {
// 	s, _, repo, _, _, db, err := test.InitTestServer()
// 	defer db.Close()
// 	defer database.TeardownTestDB(db, "\"user\"", "product")
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	// Create test data in db
// 	id, _ := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})
// 	description := "test"
// 	repo.Product().Create(&model.CreateProduct{
// 		Title:         "test",
// 		Slug:          "test",
// 		CreatedById:   id,
// 		Description:   &description,
// 		Calories:      123,
// 		Squirrels:     123,
// 		Fats:          123,
// 		Carbohydrates: 123,
// 	})
// 	testCases := []struct {
// 		name         string
// 		endpoint     string
// 		expectedData interface{}
// 		expectedCode int
// 	}{
// 		{
// 			name:     "Valide",
// 			endpoint: "/v1/product/1",
// 			expectedData: &model.Product{
// 				Base:          model.Base{Id: 1},
// 				Title:         "test",
// 				Slug:          "test",
// 				Description:   &description,
// 				Calories:      123,
// 				Squirrels:     123,
// 				Fats:          123,
// 				Carbohydrates: 123,
// 			},
// 			expectedCode: 200,
// 		},
// 		{
// 			name:         "Not found",
// 			endpoint:     "/v1/product/999",
// 			expectedCode: 404,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Value Error",
// 				Title:  "Not Found",
// 				Detail: "Product not found",
// 			},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, tc.endpoint, nil)
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
// 			// database.TeardownTestDB(db, "product")
// 		})
// 	}
// }

// func TestProductGetFilteredList(t *testing.T) {
// 	s, _, repo, _, _, db, err := test.InitTestServer()
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	defer db.Close()
// 	defer database.TeardownTestDB(db, "\"user\"", "product")
// 	// Create test data in db
// 	id, _ := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})

// 	description := "test"
// 	productId, _ := repo.Product().Create(&model.CreateProduct{
// 		Title:         "test",
// 		Slug:          "test",
// 		CreatedById:   id,
// 		Description:   &description,
// 		Calories:      123,
// 		Squirrels:     124,
// 		Fats:          123,
// 		Carbohydrates: 123,
// 	})

// 	repo.Product().Create(&model.CreateProduct{
// 		Title:         "kek",
// 		Slug:          "kek",
// 		CreatedById:   id,
// 		Description:   &description,
// 		Calories:      222,
// 		Squirrels:     222,
// 		Fats:          222,
// 		Carbohydrates: 222,
// 	})

// 	expectedData := model.ProductList{
// 		Total:  1,
// 		Limit:  25,
// 		Offset: 0,
// 		Data: []*model.Product{
// 			{
// 				Base:          model.Base{Id: productId},
// 				Title:         "test",
// 				Slug:          "test",
// 				Description:   &description,
// 				Calories:      123,
// 				Squirrels:     124,
// 				Fats:          123,
// 				Carbohydrates: 123,
// 			},
// 		},
// 	}
// 	testCases := []struct {
// 		name         string
// 		endpoint     string
// 		params       interface{}
// 		expectedCode int
// 		expectedData interface{}
// 	}{
// 		{
// 			name:         "With title",
// 			endpoint:     "/v1/product?title=test",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "With slug",
// 			endpoint:     "/v1/product?slug=test",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Calories",
// 			endpoint:     "/v1/product?calories__gte=100&calories__lte=200",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Squirrles",
// 			endpoint:     "/v1/product?squirrels__gte=100&squirrels__lte=200",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Fats",
// 			endpoint:     "/v1/product?fats__gte=100&fats__lte=200",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Carbohydrates",
// 			endpoint:     "/v1/product?carbohydrates__gte=100&carbohydrates__lte=200",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, tc.endpoint, nil)
// 			s.ServeHTTP(rec, req)
// 			if tc.expectedCode != rec.Code {
// 				t.Error("Wrong status code", rec.Code, tc.expectedCode)
// 			}
// 			response := strings.TrimRight(rec.Body.String(), "\n")
// 			expectedBytes, _ := json.Marshal(expectedData)
// 			expectedString := string(expectedBytes)
// 			if response != expectedString {
// 				t.Error("Wrong data", response, expectedString, len(response), len(expectedString))
// 			}
// 		})
// 	}
// }

// func TestProductGetEmptyList(t *testing.T) {
// 	s, _, repo, _, _, db, err := test.InitTestServer()
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	defer db.Close()
// 	defer database.TeardownTestDB(db, "\"user\"", "product")

// 	id, _ := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})

// 	description := "test"
// 	repo.Product().Create(&model.CreateProduct{
// 		Title:         "test",
// 		Slug:          "test",
// 		CreatedById:   id,
// 		Description:   &description,
// 		Calories:      123,
// 		Squirrels:     124,
// 		Fats:          123,
// 		Carbohydrates: 123,
// 	})

// 	expectedData := model.ProductList{
// 		Total:  0,
// 		Limit:  25,
// 		Offset: 0,
// 		Data:   []*model.Product{},
// 	}
// 	testCases := []struct {
// 		name         string
// 		endpoint     string
// 		params       interface{}
// 		expectedCode int
// 		expectedData interface{}
// 	}{
// 		{
// 			name:         "Title",
// 			endpoint:     "/v1/product?title=123",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "slug",
// 			endpoint:     "/v1/product?slug=123",
// 			params:       nil,
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodGet, tc.endpoint, nil)
// 			s.ServeHTTP(rec, req)
// 			if tc.expectedCode != rec.Code {
// 				t.Error("Wrong status code", rec.Code, tc.expectedCode)
// 			}
// 			response := strings.TrimRight(rec.Body.String(), "\n")
// 			expectedBytes, _ := json.Marshal(expectedData)
// 			expectedString := string(expectedBytes)
// 			if response != expectedString {
// 				t.Error("Wrong data", response, expectedString, len(response), len(expectedString))
// 			}
// 		})
// 	}
// }

// func TestProductCreate(t *testing.T) {
// 	s, _, repo, _, config, db, err := test.InitTestServer()
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	defer db.Close()
// 	defer database.TeardownTestDB(db, "\"user\"")
// 	id, err := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	dbUser, err := repo.User().GetById(id)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	accessToken, _ := service.GenerateToken(model.AccessTokenType, dbUser, 60, config.TokenSecret)

// 	defaultPayload := map[string]interface{}{
// 		"calories":      111,
// 		"carbohydrates": 111,
// 		"description":   nil,
// 		"fats":          111,
// 		"slug":          "test",
// 		"squirrels":     111,
// 		"title":         "test",
// 	}
// 	testCases := []struct {
// 		name         string
// 		endpoint     string
// 		authHeader   interface{}
// 		payload      interface{}
// 		expectedCode int
// 		expectedData interface{}
// 	}{
// 		{
// 			name:         "Valid data",
// 			endpoint:     "/v1/product",
// 			payload:      defaultPayload,
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 200,
// 			expectedData: &model.Product{
// 				Base:          model.Base{Id: 1},
// 				Title:         "test",
// 				Slug:          "test",
// 				Description:   nil,
// 				Calories:      111,
// 				Squirrels:     111,
// 				Fats:          111,
// 				Carbohydrates: 111,
// 			},
// 		},
// 		{
// 			name:         "Without payload",
// 			endpoint:     "/v1/product",
// 			payload:      map[string]interface{}{},
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 422,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Validation Error",
// 				Title:  "Unprocessable Entity",
// 				Detail: map[string]interface{}{"slug": "cannot be blank", "title": "cannot be blank"},
// 			},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			rec := httptest.NewRecorder()
// 			jsonData, _ := json.Marshal(tc.payload)
// 			req, _ := http.NewRequest(http.MethodPost, tc.endpoint, bytes.NewBuffer(jsonData))
// 			if tc.authHeader != nil {
// 				req.Header.Add("Authorization", tc.authHeader.(string))
// 			}
// 			req.Header.Add("Content-Type", "application/json")
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
// 			database.TeardownTestDB(db, "product")
// 		})
// 	}
// }

// func TestProductUpdate(t *testing.T) {
// 	endpoint := "/v1/product/1"
// 	s, _, repo, _, config, db, err := test.InitTestServer()
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	defer db.Close()
// 	defer database.TeardownTestDB(db, "\"user\"")
// 	id, err := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	dbUser, err := repo.User().GetById(id)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	accessToken, _ := service.GenerateToken(model.AccessTokenType, dbUser, 60, config.TokenSecret)
// 	anotherAccessToken, _ := service.GenerateToken(model.AccessTokenType, &model.User{ID: 2, Phone: "+71111111111"}, 60, config.TokenSecret)
// 	updatedDescription := "updated"
// 	defaultPayload := map[string]interface{}{
// 		"calories":      111,
// 		"carbohydrates": 111,
// 		"description":   &updatedDescription,
// 		"fats":          111,
// 		"slug":          "updated",
// 		"squirrels":     111,
// 		"title":         "updated",
// 	}
// 	description := "test"

// 	// testDescription := "test"
// 	testCases := []struct {
// 		name         string
// 		endpoint     string
// 		authHeader   interface{}
// 		payload      interface{}
// 		expectedCode int
// 		expectedData interface{}
// 	}{
// 		{
// 			name:         "Valid data",
// 			endpoint:     endpoint,
// 			payload:      defaultPayload,
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 200,
// 			expectedData: &model.Product{
// 				Base:          model.Base{Id: 1},
// 				Title:         "updated",
// 				Slug:          "updated",
// 				Description:   &updatedDescription,
// 				Calories:      111,
// 				Squirrels:     111,
// 				Fats:          111,
// 				Carbohydrates: 111,
// 			},
// 		},
// 		{
// 			name:         "Without payload",
// 			endpoint:     endpoint,
// 			payload:      map[string]interface{}{},
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 200,
// 			expectedData: &model.Product{
// 				Base:          model.Base{Id: 1},
// 				Title:         "test",
// 				Slug:          "test",
// 				CreatedById:   id,
// 				Description:   &description,
// 				Calories:      123,
// 				Squirrels:     124,
// 				Fats:          123,
// 				Carbohydrates: 123,
// 			},
// 		},
// 		{
// 			name:         "User without permissions",
// 			endpoint:     endpoint,
// 			payload:      map[string]interface{}{},
// 			authHeader:   "Bearer " + anotherAccessToken,
// 			expectedCode: 403,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Forbidden",
// 				Detail: "Forbidden",
// 			},
// 		},
// 		{
// 			name:         "Not found",
// 			endpoint:     "/v1/product/999",
// 			payload:      map[string]interface{}{},
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 404,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Value Error",
// 				Title:  "Not Found",
// 				Detail: "Product not found",
// 			},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			repo.Product().Create(&model.CreateProduct{
// 				Title:         "test",
// 				Slug:          "test",
// 				CreatedById:   id,
// 				Description:   &description,
// 				Calories:      123,
// 				Squirrels:     124,
// 				Fats:          123,
// 				Carbohydrates: 123,
// 			})

// 			rec := httptest.NewRecorder()
// 			jsonData, _ := json.Marshal(tc.payload)
// 			req, _ := http.NewRequest(http.MethodPatch, tc.endpoint, bytes.NewBuffer(jsonData))
// 			if tc.authHeader != nil {
// 				req.Header.Add("Authorization", tc.authHeader.(string))
// 			}
// 			req.Header.Add("Content-Type", "application/json")
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
// 			database.TeardownTestDB(db, "product")
// 		})
// 	}
// }

// func TestProductDelete(t *testing.T) {
// 	endpoint := "/v1/product/1"
// 	s, _, repo, _, config, db, err := test.InitTestServer()
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	defer db.Close()
// 	defer database.TeardownTestDB(db, "\"user\"")
// 	id, err := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	dbUser, err := repo.User().GetById(id)
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}

// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	accessToken, _ := service.GenerateToken(model.AccessTokenType, dbUser, 60, config.TokenSecret)
// 	anotherAccessToken, _ := service.GenerateToken(model.AccessTokenType, &model.User{ID: 2, Phone: "+71111111111"}, 60, config.TokenSecret)

// 	testCases := []struct {
// 		name         string
// 		endpoint     string
// 		authHeader   interface{}
// 		expectedCode int
// 		expectedData interface{}
// 	}{
// 		{
// 			name:         "Saccess",
// 			endpoint:     endpoint,
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 200,
// 			expectedData: &model.Product{
// 				Base:          model.Base{Id: 1},
// 				Title:         "deleted",
// 				Slug:          "deleted",
// 				Description:   nil,
// 				Calories:      111,
// 				Squirrels:     111,
// 				Fats:          111,
// 				Carbohydrates: 111,
// 			},
// 		},
// 		{
// 			name:         "User without permissions",
// 			endpoint:     endpoint,
// 			authHeader:   "Bearer " + anotherAccessToken,
// 			expectedCode: 403,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Permission Error",
// 				Title:  "Forbidden",
// 				Detail: "Forbidden",
// 			},
// 		},
// 		{
// 			name:         "Not found",
// 			endpoint:     "/v1/product/999",
// 			authHeader:   "Bearer " + accessToken,
// 			expectedCode: 404,
// 			expectedData: &response.ErrorResponse{
// 				Type:   "Value Error",
// 				Title:  "Not Found",
// 				Detail: "Product not found",
// 			},
// 		},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			repo.Product().Create(&model.CreateProduct{
// 				Title:         "deleted",
// 				Slug:          "deleted",
// 				CreatedById:   id,
// 				Description:   nil,
// 				Calories:      111,
// 				Squirrels:     111,
// 				Fats:          111,
// 				Carbohydrates: 111,
// 			})
// 			rec := httptest.NewRecorder()
// 			req, _ := http.NewRequest(http.MethodDelete, tc.endpoint, nil)
// 			if tc.authHeader != nil {
// 				req.Header.Add("Authorization", tc.authHeader.(string))
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
// 			database.TeardownTestDB(db, "product")
// 		})
// 	}
// }
