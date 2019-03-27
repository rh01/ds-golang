package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	m := ManufacturingDirector{}
	c := &CarBuilder{}

	m.SetBuilder(c)
	m.Construct()

	car := c.Build()

	if car.Wheels != 4 {
		t.Fail()
	}

	if car.Structure != "Car" {
		t.Fail()
	}

	if car.Seats != 5 {
		t.Fail()
	}


	b := &BikeBuilder{}

	m.SetBuilder(b)
	m.Construct()

	bike := b.Build()

	if bike.Wheels != 2 {
		t.Fail()
	}

	if bike.Structure != "Bike" {
		t.Fail()
	}

	if bike.Seats != 2 {
		t.Fail()
	}
}
