package examples

import (
	"fmt"
)

type Robot struct {
	name    string
	wheels  int
	color   string
	battery float32
}

func NewRobot(name string) *Robot {
	return &Robot{name: name}
}

// SetWheels install wheels
func (p *Robot) SetWheels(num int) error {
	if num < 1 {
		return fmt.Errorf("invalid wheels number [%v]", num)
	}
	p.wheels = num

	return nil
}

// Paint the robot with color
func (p *Robot) Paint(color string) error {
	// some operations maybe has error
	// ......

	p.color = color

	return nil
}

// Charge the robot
func (p *Robot) Charge(num float32) error {
	// some operations maybe has error
	// ......

	p.battery = num

	return nil
}

// GetReport get the robot report
func (p *Robot) GetReport() (string, error) {
	if p.wheels < 1 {
		return "", fmt.Errorf("wheels number is invalid: %v", p.wheels)
	}
	if p.color == "" {
		return "", fmt.Errorf("color is empty")
	}

	return fmt.Sprintf("Robot %v has %v wheels, color %v, battery %v\n", p.name, p.wheels, p.color, p.battery), nil
}
