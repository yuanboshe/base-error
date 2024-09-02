package main

import (
	"fmt"
	"github.com/yuanboshe/base-error/examples/robot_with_berr"
)

func main() {
	robot := examples.NewRobot("robot_with_berr")

	report := robot.SetWheels(4).Paint("white").Charge(0.8).GetReport()
	if robot.Err() != nil {
		// do something error handling
		panic(robot.Err())
	}

	fmt.Println(report)
}
