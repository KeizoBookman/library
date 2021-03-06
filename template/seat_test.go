package template_test

import (
	"github.com/KeizoBookman/library/seat"
	"html/template"
	"os"
	"testing"
)

func TestSeatCommon(t *testing.T) {

	tmpl := template.Must(template.ParseFiles("seat.tmpl"))

	file, err := os.OpenFile("testhtml/seat_test.html", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		t.Error("testhtml/seat_test.html can not open:", err.Error())
	}

	seat := seat.Seat{
		Testing: false,
	}
	err = tmpl.Execute(file, seat)
	if err != nil {
		t.Error("seat.tmpl can not execute, beause, " + err.Error())
	}

}
