package user

import (
	"car/vehicle"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAddition(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	var driver = Driver{}
	
	car := vehicle.NewMockICar(controller)
	
	car.EXPECT().Start().Return(true)

	driver.StartCar(car)

}


