package elon

import "strconv"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
  if car.battery < car.batteryDrain {
		return
	}
	car.distance = car.distance + car.speed
	car.battery = car.battery - car.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
  return "Driven " + strconv.Itoa(c.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
  return "Battery at " + strconv.Itoa(c.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
  distanceAvaliable := (car.battery / car.batteryDrain) * car.speed
	distanceDiff := trackDistance - car.distance
	return distanceDiff <= distanceAvaliable
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
