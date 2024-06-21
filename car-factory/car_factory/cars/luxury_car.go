package cars

import "fmt"

// LuxuryCar is a concrete type that implements the Car interface.
type LuxuryCar struct {
	Brand string
}

func (lc LuxuryCar) Drive() {
	fmt.Printf("Driving a luxury %s car\n", lc.Brand)
}
