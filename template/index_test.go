package template_test

import (
	"html/template"
	"os"
	"testing"
)

func TestIndex(t *testing.T) {

	tmpl, err := template.ParseFiles("index.tmpl")
	if err != nil {
		t.Error("index.tmpl can not parse. because," + err.Error())
	}

	err = tmpl.Execute(os.Stderr, nil)
	if err != nil {
		t.Error("index.tmpl can not execute, beause, " + err.Error())
	}
}
