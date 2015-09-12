package seat_test

import (
	"encoding/json"
	"github.com/KeizoBookman/library/seat"
	"testing"
)

func TestSeatKamigakari(t *testing.T) {
	var s seat.Seat
	var bytes []byte
	err := json.Unmarshal(bytes, s)
	if err != nil {
		t.Error("")
	}
}
