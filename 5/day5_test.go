package main

import "testing"

func TestDay5(t *testing.T) {
	seatNumber := DecodeSeatNumber("FBFBBFFRLR")
	if seatNumber != 357 {
		t.Fail()
	}
}
