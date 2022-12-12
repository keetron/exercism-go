package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(c Car) Car {
	if c.battery < c.batteryDrain {
		return c
	}
	return Car{
		battery:      c.battery - c.batteryDrain,
		batteryDrain: c.batteryDrain,
		speed:        c.speed,
		distance:     c.distance + c.speed,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(c Car, t Track) bool {
	l := t.distance / c.speed
	return c.battery >= l*c.batteryDrain
}
