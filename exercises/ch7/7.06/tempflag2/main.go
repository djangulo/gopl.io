// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	tempconv "github.com/djangulo/gopl.io/exercises/ch7/7.06/tempconv2"
)

//!+
var (
	temp  = tempconv.CelsiusFlag("temp", 20.0, "the temperature in C")
	tempF = tempconv.FahrenheitFlag("tempf", 68.0, "the temperature in F")
)

func main() {
	flag.Parse()
	fmt.Println(*temp)
	fmt.Println(*tempF)
}

//!-
