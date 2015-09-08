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

	err = tmpl.Execute(os.Stderr, nil)
	if err != nil {
		t.Error("search.tmpl can not execute, beause, " + err.Error())
	}
}
