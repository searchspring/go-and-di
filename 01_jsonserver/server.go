package main

import (
	"github.com/unrolled/render"
	"net/http"
	"strconv"
)

var renderer *render.Render
const port = 8090

func hello(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{"hello": "there"}

	renderer.JSON(w, http.StatusOK, response)
}

func headers(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{}

	for name, headers := range req.Header {
		for _, headerValue := range headers {
			response[name] = headerValue
		}
	}

	renderer.JSON(w, http.StatusOK, response)
}

func main() {
	renderer = render.New()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
