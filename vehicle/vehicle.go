package vehicle

type ICar interface {
	Start() bool
	Stop() bool
	Accelarate(speed float64) float64
	Brake() bool
}

type ElectricCar struct {
	carState string
	speed    float64
}

func (car *ElectricCar) Start() bool {
	car.carState = "running"
	return true
}

func (car *ElectricCar) Stop() bool {
	if car.carState == "running" {
		car.speed = 0
		car.carState = "stopped"
	}
	return true
}

func (car *ElectricCar) Brake() bool {
	if car.carState == "running" {
		car.carState = "stopped"
	}
	return true
}

func (car *ElectricCar) Accelarate(miles float64) float64 {
	if miles > 0 {
		car.speed += miles
	}
	return car.speed
}