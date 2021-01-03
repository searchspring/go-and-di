package main

import (
	"github.com/unrolled/render"
	"net/http"
)

var r *render.Render

func hello(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{"hello": "there"}

	r.JSON(w, http.StatusOK, response)
}

func headers(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{}

	for name, headers := range req.Header {
		for _, headerValue := range headers {
			response[name] = headerValue
		}
	}

	r.JSON(w, http.StatusOK, response)
}

func main() {
	r = render.New()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
