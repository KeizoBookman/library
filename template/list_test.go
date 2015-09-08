package template_test

import (
	"html/template"
	"os"
	"testing"
)

func TestList(t *testing.T) {

	tmpl, err := template.ParseFiles("list.tmpl")
	if err != nil {
		t.Error("list.tmpl can not parse. because," + err.Error())
	}

	err = tmpl.Execute(os.Stderr, nil)
	if err != nil {
		t.Error("list.tmpl can not execute, beause, " + err.Error())
	}
}
