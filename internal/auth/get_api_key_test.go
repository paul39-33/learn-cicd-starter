package auth

import (
	"net/http"
	"testing"
)

func TestGetAPI_Success(t *testing.T) {
	req := make(http.Header)

	// Set a header
	req.Set("Content-Type", "application/json")
	req.Add("Authorization", "ApiKey SOME_API_VALUE")

	apikey, err := GetAPIKey(req)
	if err != nil {
		t.Fatalf("GetAPIKey error: %v", err)
	}
	if apikey != "SOME_API_VALUE"{
		t.Errorf("error expected: %v, got %v", "SOME_API_VALUE", apikey)
	}
}

func TestGetAPI_Fail(t *testing.T) {
	req := make(http.Header)

	// Set a header
	req.Set("Content-Type", "application/json")

	_, err := GetAPIKey(req)
	if err == nil {
		t.Fatalf("expecting an error, but no error instead")
	}
}
