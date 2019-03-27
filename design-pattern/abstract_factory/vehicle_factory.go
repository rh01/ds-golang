package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	// case CarFactoryType:

	//
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recongnized", f))
	}
}
