package examples

import (
	"fmt"
	"github.com/yuanboshe/base-error/berr"
)

type Robot struct {
	name    string
	wheels  int
	color   string
	battery float32

	berr.BaseErr[Robot]
}

func NewRobot(name string) *Robot {
	robot := &Robot{name: name}
	robot.InitAddr(robot)

	return robot
}

// SetWheels install wheels
func (p *Robot) SetWheels(num int) *Robot {
	if p.Err() != nil {
		return p
	}

	if num < 1 {
		return p.SetErr(fmt.Errorf("invalid wheels number [%v]", num))
	}
	p.wheels = num

	return p
}

// Paint the robot with color
func (p *Robot) Paint(color string) *Robot {
	if p.Err() != nil {
		return p
	}

	p.color = color

	return p
}

// Charge the robot
func (p *Robot) Charge(num float32) *Robot {
	if p.Err() != nil {
		return p
	}

	p.battery = num

	return p
}

// GetReport get the robot report
func (p *Robot) GetReport() string {
	if p.Err() != nil {
		return ""
	}

	if p.wheels < 1 {
		p.SetErr(fmt.Errorf("wheels number is invalid: %v", p.wheels))
		return ""
	}
	if p.color == "" {
		p.SetErr(fmt.Errorf("color is empty"))
		return ""
	}

	return fmt.Sprintf("Robot %v has %v wheels, color %v, battery %v\n", p.name, p.wheels, p.color, p.battery)
}
