package template_test

import (
	"github.com/KeizoBookman/library/seat"
	"html/template"
	"os"
	"testing"
)

func TestSeatKamigakari(t *testing.T) {

	tmpl, err := template.ParseFiles("kamigakari.tmpl")
	if err != nil {
		t.Error("seat.tmpl can not parse. because," + err.Error())
	}
	file, err := os.OpenFile("testhtml/kamigakari_test.html", os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		t.Error("testhtml/seat_test.html can not open:", err.Error())
	}

	seat := seat.Seat{
		Testing: true,
		JSPath:  "../../public/js/jquery-1.11.3.min.js",
	}
	err = tmpl.Execute(file, seat)
	if err != nil {
		t.Error("seat.tmpl can not execute, beause, " + err.Error())
	}

}
