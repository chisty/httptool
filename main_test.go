package main

import "testing"

func TestGetFormattedAddress(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"http://adjust.com", "http://adjust.com", false},
		{"adjust.com", "http://adjust.com", false},
		{"http:// adjust.com", "", true},
		{"www.adjust.com", "http://www.adjust.com", false},
		{"google.com", "http://google.com", false},
		{"http://google.com", "http://google.com", false},
		{"https://google.com", "https://google.com", false},
		{"http//google.com", "http://http//google.com", false},
	}

	for _, test := range tests {
		url, err := getFormattedAddress(test.input)
		if url != test.expected {
			t.Errorf("Test failed. input: %s, expected: %s, but received: %s\n", test.input, test.expected, url)
		}
		if !test.hasError && err != nil {
			t.Errorf("Test failed. input: %s, error not expected, but received err: %s\n", test.input, err.Error())
		}
		if test.hasError && err == nil {
			t.Errorf("Test failed. input: %s, error expected, but no errore received.\n", test.input)
		}
	}
}
