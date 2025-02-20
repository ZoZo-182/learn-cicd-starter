package auth

import (
    "reflect"
    "testing"
    "net/http"
)

func TestGetAPIKey(t *testing.T) {
    header := http.Header{} // Empty header (no auth header)

    got, err := GetAPIKey(header)
        wantErr := ErrNoAuthHeaderIncluded

    if !reflect.DeepEqual(wantErr, err) {
        t.Fatalf("Expected: %v, got: %v", wantErr, err)
    }

    if got != "" {
        t.Fatalf("Expected empty API key, got %q", got)
    }
}
