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

	file, err := os.OpenFile("testhtml/list_test.html", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		t.Error("testhtml/list_test.html", err.Error())
	}
	err = tmpl.Execute(file, nil)
	if err != nil {
		t.Error("list.tmpl can not execute, beause, " + err.Error())
	}
}
