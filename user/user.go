package main

import (
	"car/vehicle"
	"fmt"
)

func main() {
	var myCar vehicle.ElectricCar = vehicle.ElectricCar{}
	var eCar vehicle.ICar = &myCar
	eCar.Start()
	
	fmt.Println(myCar)

	eCar.Accelarate(10.5)
	fmt.Println(myCar)

	eCar.Stop()

	fmt.Println(myCar)

}