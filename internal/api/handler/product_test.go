package handler_test

import (
	"encoding/json"
	"food/internal/api/model"
	"food/internal/api/repository/postgres"
	"food/internal/test"
	"food/pkg/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProductGetFilteredList(t *testing.T) {
	s, db, err := test.InitTestServer()
	defer database.TeardownTestDB(db, "\"user\"", "product")
	if err != nil {
		t.Fatal(err)
		return
	}
	// Create test data in db
	repo := postgres.NewRepository(db)
	id, _ := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})

	description := "test"
	productId, _ := repo.Product().Create(&model.CreateProduct{
		Title:         "test",
		Slug:          "test",
		CreatedById:   id,
		Description:   &description,
		Calories:      123,
		Squirrels:     124,
		Fats:          123,
		Carbohydrates: 123,
	})

	repo.Product().Create(&model.CreateProduct{
		Title:         "kek",
		Slug:          "kek",
		CreatedById:   id,
		Description:   &description,
		Calories:      222,
		Squirrels:     222,
		Fats:          222,
		Carbohydrates: 222,
	})

	expectedData := model.ProductList{
		Total:  1,
		Limit:  25,
		Offset: 0,
		Data: []*model.Product{
			{
				Base:          model.Base{Id: productId},
				Title:         "test",
				Slug:          "test",
				Description:   &description,
				Calories:      123,
				Squirrels:     124,
				Fats:          123,
				Carbohydrates: 123,
			},
		},
	}
	testCases := []struct {
		name         string
		endpoint     string
		params       interface{}
		expectedCode int
		expectedData interface{}
	}{
		{
			name:         "With title",
			endpoint:     "/v1/product?title=test",
			params:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "With slug",
			endpoint:     "/v1/product?slug=test",
			params:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Calories",
			endpoint:     "/v1/product?calories__gte=100&calories__lte=200",
			params:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Squirrles",
			endpoint:     "/v1/product?squirrels__gte=100&squirrels__lte=200",
			params:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Fats",
			endpoint:     "/v1/product?fats__gte=100&fats__lte=200",
			params:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Carbohydrates",
			endpoint:     "/v1/product?carbohydrates__gte=100&carbohydrates__lte=200",
			params:       nil,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.endpoint, nil)
			s.ServeHTTP(rec, req)
			if tc.expectedCode != rec.Code {
				t.Error("Wrong status code", rec.Code, tc.expectedCode)
			}
			response := strings.TrimRight(rec.Body.String(), "\n")
			expectedBytes, _ := json.Marshal(expectedData)
			expectedString := string(expectedBytes)
			if response != expectedString {
				t.Error("Wrong data", response, expectedString, len(response), len(expectedString))
			}
		})
	}
}

func TestProductGetEmptyList(t *testing.T) {
	s, db, err := test.InitTestServer()
	defer database.TeardownTestDB(db, "\"user\"", "product")
	if err != nil {
		t.Fatal(err)
		return
	}

	repo := postgres.NewRepository(db)
	id, _ := repo.Auth().CreateUser(&model.Login{Phone: "+79999999999", Code: "1111"})

	description := "test"
	repo.Product().Create(&model.CreateProduct{
		Title:         "test",
		Slug:          "test",
		CreatedById:   id,
		Description:   &description,
		Calories:      123,
		Squirrels:     124,
		Fats:          123,
		Carbohydrates: 123,
	})

	expectedData := model.ProductList{
		Total:  0,
		Limit:  25,
		Offset: 0,
		Data:   []*model.Product{},
	}
	testCases := []struct {
		name         string
		endpoint     string
		params       interface{}
		expectedCode int
		expectedData interface{}
	}{
		{
			name:         "Title",
			endpoint:     "/v1/product?title=123",
			params:       nil,
			expectedCode: http.StatusOK,
		},
		{
			name:         "slug",
			endpoint:     "/v1/product?slug=123",
			params:       nil,
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, tc.endpoint, nil)
			s.ServeHTTP(rec, req)
			if tc.expectedCode != rec.Code {
				t.Error("Wrong status code", rec.Code, tc.expectedCode)
			}
			response := strings.TrimRight(rec.Body.String(), "\n")
			expectedBytes, _ := json.Marshal(expectedData)
			expectedString := string(expectedBytes)
			if response != expectedString {
				t.Error("Wrong data", response, expectedString, len(response), len(expectedString))
			}
		})
	}
}
