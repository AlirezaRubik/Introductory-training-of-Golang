package Fuzz

import (
	"testing"
)

func FuzzCalculationRoomPrice(f *testing.F) {
	f.Fuzz(func(t *testing.T, roomType string, night int, persons int) {
		FullPayment:=CalculationRoomPrice(roomType, night, persons)
		if FullPayment<=0{
			t.Errorf("CalculationRoomPrice returned a non-positive value: %d", FullPayment)
		}

	})
}
