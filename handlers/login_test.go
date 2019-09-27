package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/register", ShowRegistrationPage)

	req, _ := http.NewRequest("GET", "/register", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return true
	})
}
