package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const port = 8090

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello there")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, headerValue := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, headerValue)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
