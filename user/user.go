package user

import (
	"car/vehicle"
)

type IDriver interface {
	StartCar(eCar vehicle.ICar)
	DriveCar(eCar vehicle.ICar)
}

type Driver struct {
	name string
	age float64
}

func (d *Driver) StartCar(eCar vehicle.ICar){
	eCar.Start()
}

func (d *Driver) DriveCar(eCar vehicle.ICar){
	eCar.Accelarate(10.5)
	eCar.Stop()
}