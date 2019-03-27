package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	mF, err := BuildFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := mF.Build(LuxuryCarType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d seats\n", carVehicle.NumSeats())
}
