/*Package tempconv2 performs Celsius, Fahrenheit, Kelvin and Rankine
temperature conversions.
*/
package tempconv2

import (
	"flag"
	"fmt"
)

// const (
// 	AbsoluteZeroC Celsius = -273.15
// 	FreezingC     Celsius = 0
// 	BoilingC      Celsius = 100
// )

// Celsius temperature unit
type Celsius float64

// Fahrenheit temperature unit
type Fahrenheit float64

// Kelvin temperature unit
type Kelvin float64

// Rankine temperature unit
type Rankine float64

func (c Celsius) String() string    { return fmt.Sprintf("%.4f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.4f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.4fK", k) }
func (r Rankine) String() string    { return fmt.Sprintf("%.4f°R", r) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celcius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// CToR converts a Celcius temperature to Rankine.
func CToR(c Celsius) Rankine { return Rankine((c + 273.15) * 9 / 5) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }

// FToR converts a Fahrenheit temperature to Rankine.
func FToR(f Fahrenheit) Rankine { return Rankine(f + 459.67) }

// RToC converts a Rankine temperature to Celsius.
func RToC(r Rankine) Celsius { return Celsius((r - 491.67) * 5 / 9) }

// RToK converts a Rankine temperature to Kelvin.
func RToK(r Rankine) Kelvin { return Kelvin(r * 5 / 9) }

// RToF converts a Rankine temperature to Fahrenheit.
func RToF(r Rankine) Fahrenheit { return Fahrenheit(r - 459.67) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToR converts a Kelvin temperature to Rankine.
func KToR(k Kelvin) Rankine { return Rankine(k * 9 / 5) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5 - 459.67) }

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	case "R", "°R":
		f.Celsius = RToC(Rankine(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type fahrenheitFlag struct{ Fahrenheit }

func (f *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "F", "°F":
		f.Fahrenheit = Fahrenheit(value)
		return nil
	case "C", "°C":
		f.Fahrenheit = CToF(Celsius(value))
		return nil
	case "K", "°K":
		f.Fahrenheit = KToF(Kelvin(value))
		return nil
	case "R", "°R":
		f.Fahrenheit = RToF(Rankine(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// FahrenheitFlag defines a Fahrenheit flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100F".
func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	f := fahrenheitFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Fahrenheit
}
