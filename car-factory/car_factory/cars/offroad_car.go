package cars

import "fmt"

// OffroadCar is a concrete type that implements the Car interface.
type OffroadCar struct {
	Brand string
}

func (oc OffroadCar) Drive() {
	fmt.Printf("Driving an offroad %s car\n", oc.Brand)
}
