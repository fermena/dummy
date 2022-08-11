package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/user"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {

	req, _ := http.NewRequest("GET",
		"/",
		http.NoBody)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(rr, req)

	gotCode := rr.Code
	expectCode := 200
	if gotCode != expectCode {
		t.Errorf("Got %d expect %d", gotCode, expectCode)
	}

	gotBody := rr.Body.String()
	expectBody := "insecure request sent"

	if gotBody != expectBody {
		t.Errorf("Got %s expect %s", gotBody, expectBody)
	}
	fmt.Println(rr.Body.String())
}

func TestXSSHandler(t *testing.T) {

	req, _ := http.NewRequest("GET",
		`/xss?name=</p><img src="https://www.docker.com/favicon.ico" /></body></html>`,
		http.NoBody)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(xssHandler)

	handler.ServeHTTP(rr, req)

	gotCode := rr.Code
	expectCode := 200
	if gotCode != expectCode {
		t.Errorf("Got %d expect %d", gotCode, expectCode)
	}

	gotBody := rr.Body.String()
	expectBody := `
	<html>
		<body>
			<p> Hello, </p><img src="https://www.docker.com/favicon.ico" /></body></html>!`

	if gotBody != expectBody {
		t.Errorf("Got %s expect %s", gotBody, expectBody)
	}
	fmt.Println(rr.Body.String())
}

func TestHostnameHandler(t *testing.T) {

	type testCase struct {
		name       string
		urlString  string
		expectCode int
		expectBody string
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	testCases := []testCase{
		{
			name:       "test hostname no args",
			urlString:  "/hostname",
			expectCode: 200,
			expectBody: hostname + "\n",
		},
		{
			name:       "test hostname with args",
			urlString:  "/hostname?args=-s",
			expectCode: 200,
			expectBody: strings.Split(hostname, ".")[0] + "\n",
		},
		{
			name:       "test hostname with malicious args",
			urlString:  `/hostname?args=%26%26 whoami`,
			expectCode: 200,
			expectBody: hostname + "\n" + user.Username + "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET",
				tc.urlString,
				http.NoBody)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(hostnameHandler)

			handler.ServeHTTP(rr, req)

			gotCode := rr.Code
			expectCode := tc.expectCode
			if gotCode != expectCode {
				t.Errorf("Got %d expect %d", gotCode, expectCode)
			}

			gotBody := rr.Body.String()
			expectBody := tc.expectBody

			if gotBody != expectBody {
				t.Errorf("Got %q expect %q", gotBody, expectBody)
			}
		})

	}

}
