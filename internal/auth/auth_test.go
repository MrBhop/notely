package auth

import (
	"testing"
	"net/http"
)

func TestGetApiKey(t *testing.T) {
	cases := []struct{
		input http.Header
		expectError bool
		want string
	}{
		{
			input: http.Header{"Authorization": {"ApiKey key"},},
			expectError: false,
			want: "key",
		},
		{
			input: http.Header{"Authorization": {"key"},},
			expectError: true,
			want: "",
		},
		{
			input: http.Header{},
			expectError: true,
			want: "",
		},
	}
	
	for _, tc := range cases {
		apiKey, err := GetAPIKey(tc.input)
		if tc.expectError && err == nil {
			t.Fatalf("Expected error but got none.")
		} else if !tc.expectError && err != nil {
			t.Fatalf("Expected no error but got: %v", err)
		}

		if apiKey != tc.want {
			t.Fatalf("Expected api key: %v, Got: %v", tc.want, apiKey)
		}
	}
}
