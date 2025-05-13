package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestTodoAPI(t *testing.T) {
	todos = map[int]Todo{}
	nextID = 1

	router := setupRouter()

	testCases := []struct {
		name           string
		method, path   string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Create todo",
			method:         "POST",
			path:           "/todos",
			body:           `{"title":"T1","completed":false}`,
			expectedStatus: http.StatusCreated,
			expectedBody:   `"id":1`,
		},
		{
			name:           "Get all todos",
			method:         "GET",
			path:           "/todos",
			expectedStatus: http.StatusOK,
			expectedBody:   `"title":"T1"`,
		},
		{
			name:           "Get todo by id",
			method:         "GET",
			path:           "/todos/1",
			expectedStatus: http.StatusOK,
			expectedBody:   `"id":1`,
		},
		{
			name:           "Update todo",
			method:         "PUT",
			path:           "/todos/1",
			body:           `{"title":"T1 edited","completed":true}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `"completed":true`,
		},
		{
			name:           "Delete todo",
			method:         "DELETE",
			path:           "/todos/1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Get deleted todo",
			method:         "GET",
			path:           "/todos/1",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var bodyReader io.Reader
			if tc.body != "" {
				bodyReader = strings.NewReader(tc.body)
			}
			w := performRequest(router, tc.method, tc.path, bodyReader)
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.Contains(t, w.Body.String(), tc.expectedBody)
			}
		})
	}
}