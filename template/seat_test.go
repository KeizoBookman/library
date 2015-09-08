package template_test

import (
	"html/template"
	"os"
	"testing"
)

func TestSeatKamigakari(t *testing.T) {

	tmpl, err := template.ParseFiles("seat.tmpl")
	if err != nil {
		t.Error("seat.tmpl can not parse. because," + err.Error())
	}

	err = tmpl.Execute(os.Stderr, nil)
	if err != nil {
		t.Error("seat.tmpl can not execute, beause, " + err.Error())
	}
}
