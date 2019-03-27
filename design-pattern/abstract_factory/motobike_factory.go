package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (m *MotorbikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("vehicle of type %d not recongoze", v))
	}
}

// func (m *MotorbikeFactory) NewVehicle(v )  {
	
// }
