package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testBasicAuth(t *testing.T, req *http.Request, expectedCode int) *httptest.ResponseRecorder {
	t.Helper()
	resp := httptest.NewRecorder()
	handleAuth(resp, req)
	if resp.Code != expectedCode {
		t.Errorf("expected code %d, got %d\n", expectedCode, resp.Code)
	}
	return resp
}

func TestBasicAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := testBasicAuth(t, req, http.StatusUnauthorized)
	if ah := resp.Header().Get("WWW-Authenticate"); ah == "" {
		t.Error("No authentication header")
	}
	req.SetBasicAuth("user", "pass")
	testBasicAuth(t, req, http.StatusOK)
	req.SetBasicAuth("user", "")
	testBasicAuth(t, req, http.StatusForbidden)
	req.SetBasicAuth("", "")
	testBasicAuth(t, req, http.StatusForbidden)
	req.SetBasicAuth("xuser", "pass")
	testBasicAuth(t, req, http.StatusForbidden)
	req.SetBasicAuth("xuser", "")
	testBasicAuth(t, req, http.StatusForbidden)
	req.SetBasicAuth("x", "")
	testBasicAuth(t, req, http.StatusForbidden)
}
