package main

import (
	"fmt"
	"github.com/yuanboshe/base-error/examples/robot_without_berr"
)

func main() {
	robot := examples.NewRobot("robot_without_berr")

	err := robot.SetWheels(4)
	if err != nil {
		// do something error handling
		panic(err)
	}

	err = robot.Paint("white")
	if err != nil {
		// do something error handling
		panic(err)
	}

	err = robot.Charge(0.8)
	if err != nil {
		// do something error handling
		panic(err)
	}

	report, err := robot.GetReport()
	if err != nil {
		// do something error handling
		panic(err)
	}

	fmt.Println(report)
}
