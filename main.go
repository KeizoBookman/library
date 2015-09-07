package main

import (
	seat "./seat"
	view "./view"
	"fmt"
	"net/http"
	"net/http/cgi"
	"strings"
)

func main() {

	cgi.Serve(http.HandlerFunc(handler))

	req, err := cgi.Request()
	if err != nil {
		Error("request")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	handle := routing(r.URL.Path)
	handle(w, r)
}

func routing(path string) func(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(path, "/") {
		return view.Index
	}
}
