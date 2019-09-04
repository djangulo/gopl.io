/*Cf2
converts its numeric argument to Celsius and Fahrenheit.
*/
package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/djangulo/gopl.io/ch2/exercises/tempconv1"
)

type Meter float64
type Foot float64

type Pound float64
type Kilogram float64

type Liter float64
type Gallon float64

func (g Gallon) String() string    { return fmt.Sprintf("%.4f gal", g) }
func (l Liter) String() string     { return fmt.Sprintf("%.4fL", l) }
func (p Pound) String() string     { return fmt.Sprintf("%.4flb", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%.4fKg", kg) }
func (m Meter) String() string     { return fmt.Sprintf("%.4fm", m) }
func (ft Foot) String() string     { return fmt.Sprintf("%.4fft", ft) }

func MToFt(m Meter) Foot { return Foot(m * 3.2808) }
func FtToM(f Foot) Meter { return Meter(f / 3.2808) }

func LbToKg(p Pound) Kilogram { return Kilogram(p / 2.2046) }
func KgToLb(k Kilogram) Pound { return Pound(k * 2.2046) }

func GalToL(g Gallon) Liter { return Liter(g / .26417) }
func LToGal(l Liter) Gallon { return Gallon(l * .26417) }

func main() {
	for _, arg := range os.Args[1:] {
		a, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		var (
			asKg  = Kilogram(a)
			asLb  = Pound(a)
			asF   = tempconv.Fahrenheit(a)
			asC   = tempconv.Celsius(a)
			asM   = Meter(a)
			asFt  = Foot(a)
			asGal = Gallon(a)
			asL   = Liter(a)
		)
		fmt.Printf(`
'%g' conversions:
Weight:			%s = %s,		%s = %s		
Length:			%s = %s,		%s = %s
Volume:			%s = %s,		%s = %s
Temperature:		%s = %s,		%s = %s

`, a,
			asKg, KgToLb(asKg), asLb, LbToKg(asLb),
			asM, MToFt(asM), asFt, FtToM(asFt),
			asL, LToGal(asL), asGal, GalToL(asGal),
			asC, tempconv.CToF(asC), asF, tempconv.FToC(asF),
		)
	}
}
