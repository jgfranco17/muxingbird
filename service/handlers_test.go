package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewMockHandlerSuccess(t *testing.T) {
	statusCode := http.StatusOK
	content := map[string]string{"message": "Hello, World!"}
	handler := CreateNewMockHandler(statusCode, content)
	req := httptest.NewRequest(http.MethodGet, "/mock-endpoint", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, statusCode, rr.Code, "Expected HTTP status code does not match")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Expected Content-Type to be application/json")
	expectedBody := `{"message":"Hello, World!"}`
	assert.JSONEq(t, expectedBody, rr.Body.String(), "Response body does not match expected JSON")
}

func TestCreateNewMockHandlerFail_EncodeError(t *testing.T) {
	statusCode := http.StatusOK
	content := make(chan int)
	handler := CreateNewMockHandler(statusCode, content)
	req := httptest.NewRequest(http.MethodGet, "/mock-endpoint", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	expectedErrorMessage := "failed to serve mock endpoint\n"
	assert.Equal(t, expectedErrorMessage, rr.Body.String(), "Expected error message not found")
}
