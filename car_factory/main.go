package main

import (
	"fmt"
	"car_factory"
)

func main() {
	carFactory := car_factory.ConcreteCarFactory{}

	luxuryCar := carFactory.CreateCar("Mercedes", "luxury")
	offroadCar := carFactory.CreateCar("Jeep", "offroad")
	cityCar := carFactory.CreateCar("Toyota", "city")

	cars := []car_factory.Car{luxuryCar, offroadCar, cityCar}

	for _, car := range cars {
		car.Drive()
	}
}
