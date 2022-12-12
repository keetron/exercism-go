package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery < c.batteryDrain {
		return
	}
	c.battery = c.battery - c.batteryDrain
	c.distance = c.distance + c.speed
}

func (c *Car) DisplayDistance() string {
	return "Driven " + fmt.Sprint(c.distance) + " meters"
}

func (c *Car) DisplayBattery() string {
	return "Battery at " + fmt.Sprint(c.battery) + "%"
}

func (c *Car) CanFinish(trackDistance int) bool {
	return c.battery >= trackDistance/c.speed*c.batteryDrain
}
