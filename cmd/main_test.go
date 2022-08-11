package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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
