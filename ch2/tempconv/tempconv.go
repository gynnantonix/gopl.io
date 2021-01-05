// Package tempconv performs Celsius and Fahrenheit conversions
package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC             = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g&deg;C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g&deg;F", f) }
