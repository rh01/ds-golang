package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (m *ManufacturingDirector) Construct() {
	// 实现部分
	m.builder.SetSeats().SetWheels().SetStructure()
}

func (m *ManufacturingDirector) SetBuilder(b BuildProcess) {
	m.builder = b
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{c.v.Wheels, c.v.Seats, c.v.Structure}
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b

}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

func (b *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{b.v.Wheels, b.v.Seats, b.v.Structure}
}
