package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	cronCount := 0
	echoCount := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body []byte

		if r.Body != nil {
			body, _ = ioutil.ReadAll(r.Body)
		}

		fmt.Printf("%s:\n%s %s\nHeaders:\n%s\n\nBody:\n%s\n\n",
			time.Now().String(), r.Method, r.URL, r.Header, string(body))

		if r.URL.Path == "/stats" {
			fmt.Fprintf(w, "%d/%d\n", echoCount, cronCount)
		} else {
			if len(body) == 0 {
				fmt.Fprintf(w, "Hi from echo\n")
			} else {
				fmt.Fprintf(w, string(body)+"\n")
			}

			if strings.Contains(string(body), "cron") {
				cronCount++
			} else {
				echoCount++
			}

		}
	})

	fmt.Print("Listening on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
