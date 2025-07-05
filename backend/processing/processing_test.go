package processing

import (
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestFetchAbstracts(t *testing.T) {
	tests := []struct {
		name string
		client *http.Client
		idSlice []string
		expected string
		errorContains string
	}{
		{
		name: "test fetchAbstracts",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		idSlice: []string{"40613785","40612660","40612353","40612274"},
	},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := FetchEFetchResult(tc.client,tc.idSlice)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
					t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
					return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}
		})
	}
}