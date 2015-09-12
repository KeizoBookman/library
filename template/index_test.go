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

	file, err := os.OpenFile("testhtml/index_test.html", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		t.Error("testhtml/index_test.html", err.Error())
	}
	err = tmpl.Execute(file, nil)
	if err != nil {
		t.Error("index.tmpl can not execute, beause, " + err.Error())
	}
}
