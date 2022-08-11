package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func insecureRequest() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	response, err := client.Get("https://docker.com/robots.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	htmlData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(htmlData))
}
func indexHandler(w http.ResponseWriter, req *http.Request) {
	insecureRequest()
	fmt.Fprintf(w, "insecure request sent")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
