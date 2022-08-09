package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeRequest() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	url := "http://example.com/"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func main() {

	makeRequest()

}
