package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()
	if router == nil {
		t.Error("Expected a non-nil router")
	}

	// Test the "/listorders" route using httptest
	t.Run("ListOrdersRoute", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/listorders", nil)
		if err != nil {
			t.Fatal(err)
		}

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		// Add assertions based on the expected response
		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
		}

	})
}
