// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"os"
)

type hash int

const (
	hash256 = 256
	hash384 = 384
	hash512 = 512
)

var bitrate int

func init() {
	flag.IntVar(&bitrate, "hash", hash256, "SHA hashing algorithm to use, default 256")
}

func main() {
	flag.Parse()
	switch bitrate {
	case hash512:
		fmt.Fprintf(os.Stdout, "(%T)\t%[1]x\n", sha512.Sum512([]byte(os.Args[1])))
		return
	case hash384:
		fmt.Fprintf(os.Stdout, "(%T)\t%[1]x\n", sha512.Sum384([]byte(os.Args[1])))
		return
	case hash256:
		fmt.Fprintf(os.Stdout, "(%T)\t%[1]x\n", sha256.Sum256([]byte(os.Args[1])))
		return
	default:
		fmt.Fprintf(os.Stderr, "invalid hash value passed: %v\n", bitrate)
		return
	}
	// if bitrate == 512
	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
