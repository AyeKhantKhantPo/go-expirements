package car_factory

// Car interface defines the methods that every car should have.
type Car interface {
	Drive()
}

// CarFactory is an abstract factory for creating cars.
type CarFactory interface {
	CreateCar(brand, carType string) Car
}
