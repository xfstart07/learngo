// Author: xufei
// Date: 2019/4/30

package main

import "learngo/designpattern/gopatterns/creational/builder/car"

func main() {
	builder := car.NewBenzBuilder()

	blueCar := builder.Color(car.BlueColor).Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	blueCar.Drive()
	blueCar.Stop()

	IronCar := builder.Color(car.RedColor).Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	IronCar.Drive()
	IronCar.Stop()
}
