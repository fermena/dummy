package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"
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

func xssHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("foo").Parse(`{{define "T"}}
	<html>
		<body>
			<p> Hello, {{.}}!{{end}} </p>
		</body>
	</html>
	`)
	tmpl.ExecuteTemplate(w, "T", r.URL.Query().Get("name"))

}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query().Get("args")

	if len(args) == 0 {
		args = "-f"
	}

	cmd := exec.Command(`sh`, `-c`, fmt.Sprintf(`''hostname %s''`, args))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Fprintf(w, string(output))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/xss", xssHandler)
	http.HandleFunc("/hostname", hostnameHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
