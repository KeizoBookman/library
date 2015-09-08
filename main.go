package main

import (
	view "github.com/KeizoBookman/library/view"
	"net/http"
	"net/http/cgi"
	"strings"
)

func main() {

	err := cgi.Serve(http.HandlerFunc(handler))

	if err != nil {
		view.ViewFail("request")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("Content-Type:", "text/html; charset=UTF-8")
	handle, err := routing(r.URL.Path)
	if err != nil {
		view.ViewFail(err.Error())
	}
	handle(w, r)
}

type failRouting string

func (f failRouting) Error() string {
	return string(f)
}
func routing(path string) (func(w http.ResponseWriter, r *http.Request), error) {
	top := "index.cgi"
	n := strings.Index(path, top)
	if n == -1 {
		return nil, failRouting("routing can not find root " + top)
	}

	size := n + len(top)

	switch {
	case strings.HasPrefix(path[size:], "/new"):
		return view.NewSeat, nil
	case strings.HasPrefix(path[size:], "/list"):
		return view.List, nil
	case strings.HasPrefix(path[size:], "/search"):
		return view.Search, nil
	case strings.HasPrefix(path[size:], "/"), strings.HasPrefix(path[size:], ""):
		return view.Index, nil
	default:
		return http.NotFound, nil

	}
}
