package main

import (
	"car/user"
	"car/vehicle"
	"fmt"
)

func main() {
	var myCar vehicle.ElectricCar = vehicle.ElectricCar{}
	var driver = user.Driver{}
	driver.StartCar(&myCar)
	fmt.Println(myCar)
	driver.DriveCar(&myCar)	
	fmt.Println(myCar)
	
}