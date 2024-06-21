package cars

import "fmt"

// CityCar is a concrete type that implements the Car interface.
type CityCar struct {
	Brand string
}

func (cc CityCar) Drive() {
	fmt.Printf("Driving a city %s car\n", cc.Brand)
}
