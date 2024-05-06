package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		header   http.Header
		expected string
	}{
		"correctAuthKey": {
			header: http.Header{
				"Authorization": {
					"ApiKey 123123",
				},
			},
			expected: "123123"},
		"IncorrectAuthKey": {
			header: http.Header{
				"Authorization": {
					"Bearer 123123",
				},
			},
			expected: ""},
		"IncorrectAuthLength": {
			header: http.Header{
				"Authorization": {
					"Bearer 123123",
				},
			},
			expected: ""},
		"NoAuthIncluded": {
			header:   http.Header{},
			expected: ""},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, _ := GetAPIKey(test.header)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected: %v, actual: %v", test.expected, actual)
			}
		})
	}
}
