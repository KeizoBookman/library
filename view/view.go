package view

import (
	"encoding/json"
	"github.com/KeizoBookman/library/seat"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./template/index.tmpl")
	if err != nil {
		ViewFail("index")
		return
	}
	n := strings.Index(r.URL.Path, "index.cgi")
	if n == -1 {
		ViewFail("path")
	}
	data := ResourceIndex{Path: r.URL.Path[:n]}
	err = tmpl.Execute(w, data)
	if err != nil {
		ViewFail(err.Error())
	}

}

func List(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	tmpl, err := template.ParseFiles("./template/list.tmpl")
	if err != nil {
		ViewFail("list")
		return
	}
	err = tmpl.Execute(w, data)
}

func Search(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	tmpl, err := template.ParseFiles("./template/search.tmpl")
	if err != nil {
		ViewFail("search")
		return
	}
	err = tmpl.Execute(w, data)
}

func NewSeat(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./template/kamigakari.tmpl")
	if err != nil {
		ViewFail("new:" + err.Error())
		return
	}
	data := struct {
		Testing bool
	}{false}
	err = tmpl.Execute(w, data)
	if err != nil {
		ViewFail("new:")
		return
	}
	err = r.ParseForm()
	if err != nil {
	}

	c := seat.Character{}
	err = c.GetFormValue(r)
	if err != nil {
	}

	s := seat.Seat{}
	s.Character = c
	file, err := os.OpenFile("./public/store", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		ViewFail("critical--- " + err.Error())
		return
	}

	d, err := ioutil.ReadAll(file)
	if err != nil {
	}
	var prev []seat.Seat
	err = json.Unmarshal(d, &prev)
	if err != nil {
	}

	bytes, err := json.Marshal(c)
	if err != nil {
		ViewFail("Can not Marshal--- " + err.Error())
		return
	}

	_, err = file.Write(bytes)
	if err != nil {
		ViewFail("Can not Write---" + err.Error())
		return
	}
}
