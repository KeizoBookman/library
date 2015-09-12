package template_test

import (
	"html/template"
	"os"
	"testing"
)

func TestSeatSearch(t *testing.T) {

	tmpl, err := template.ParseFiles("search.tmpl")
	if err != nil {
		t.Error("search.tmpl can not parse. because," + err.Error())
	}

	file, err := os.OpenFile("testhtml/search_test.html", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		t.Error("testhtml/search_test.html", err.Error())
	}
	err = tmpl.Execute(file, nil)
	if err != nil {
		t.Error("search.tmpl can not execute, beause, " + err.Error())
	}
}
