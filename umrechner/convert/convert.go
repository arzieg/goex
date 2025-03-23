package convert

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kg float64
type Pound float64
type Feet float64
type Meters float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kg) String() string {
	return fmt.Sprintf("%g KG", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (m Meters) String() string {
	return fmt.Sprintf("%g m", m)
}

func (t Feet) String() string {
	return fmt.Sprintf("%g ft", t)
}
