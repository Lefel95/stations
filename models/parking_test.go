package models

import (
	"testing"
)

func TestParking(t *testing.T) {
	p := &Parking{
		total:          10,
		totalAvailable: 10,
		head:           nil,
	}

	fillSlots := func(vtype string, nn int) {
		for i := 0; i < nn; i++ {
			var n = &Slot{
				Prior:   nil,
				Type:    vtype,
				Used:    false,
				Vehicle: nil,
				Next:    nil,
			}
			p.AddParkingSlot(n)
		}
	}

	fillSlots(BIKE, 10)
	fillSlots(CAR, 10)
	fillSlots(VAN, 10)

	// Create Vehicle instances for testing
	bike := &Vehicle{BIKE}
	car := &Vehicle{CAR}

	// Test Total() method
	if p.Total() != 10 {
		t.Errorf("Expected total to be 10, but got %d", p.Total())
	}

	// Test TotalAvailable() method
	if p.TotalAvailable() != 10 {
		t.Errorf("Expected total available to be 10, but got %d", p.TotalAvailable())
	}

	// Test Full() method when not full
	if p.Full() {
		t.Errorf("Expected parking not to be full")
	}

	// Test Empty() method when not empty
	if !p.Empty() {
		t.Errorf("Expected parking to be empty")
	}

	// Test HasSlots() method
	if !p.HasSlots(BIKE) {
		t.Errorf("Expected available bike slots")
	}

	// Test Count() method
	if (p.Count(CAR) != 0 ){
		t.Errorf("Expected 0 cars slots, but got %d", p.Count(CAR))
	}

	// Test Park() method
	if !p.Park(bike) {
		t.Errorf("Expected to park bike, but failed")
	}

	if !p.Park(car) {
		t.Errorf("Expected parking car to fail when parking is full")
	}

	// Test Release() method
	if !p.Release(BIKE) {
		t.Errorf("Expected to release bike, but failed")
	}

	if !p.Release(VAN) {
		t.Errorf("Expected releasing van to fail when there's no van")
	}
}
