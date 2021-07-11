package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

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

func TestFetchUrl(t *testing.T) {
	expected := "Hello From Google"
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(expected))
	}))

	defer server.Close()

	res, err := fetch(server.URL, server.Client())
	if err != nil {
		t.Errorf("Test failed. no error expected. but error received: %s", err.Error())
	}
	if string(res) != expected {
		t.Errorf("Test failed. expected: Hello From Google. but received: %s", string(res))
	}
}

func TestStartWorker(t *testing.T) {
	reqChan := make(chan string)
	reschan := make(chan response)
	var wg sync.WaitGroup

	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"http://adjust.com", "http://adjust.com", false},
		{"adjust.com", "http://adjust.com", false},
		{"http:// adjust.com", "http:// adjust.com", true},
		{"www.adjust.com", "http://www.adjust.com", false},
		{"google.com", "http://google.com", false},
		{"http://google.com", "http://google.com", false},
		{"https://google.com", "https://google.com", false},
		{"http//google.com", "http://http//google.com", true},
	}

	wg.Add(len(tests))

	go func(resChan <-chan response, wg *sync.WaitGroup, tests []struct {
		input    string
		expected string
		hasError bool
	}) {
		for _, test := range tests {
			res := <-resChan
			if res.address != test.expected {
				t.Errorf("Test failed. input %s, expected %s. but received %s\n", test.input, test.expected, res.address)
			}
			if test.hasError && res.err == nil {
				t.Errorf("Test failed. input %s, error expected. but no error received\n", test.input)
			}
			if !test.hasError && res.err != nil {
				t.Errorf("Test failed. input %s, no error expected. but error received\n", test.input)
			}
			wg.Done()
		}
	}(reschan, &wg, tests)

	go startWorker(reqChan, reschan)

	for _, test := range tests {
		reqChan <- test.input
	}
	wg.Wait()
}
