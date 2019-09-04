/*Package tempconv1 performs Celsius, Fahrenheit, Kelvin and Rankine
temperature conversions.
*/
package tempconv1

import "fmt"

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
