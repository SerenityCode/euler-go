package main

import (
	"fmt"
	"math"
)

// Special Pythagorean triplet https://projecteuler.net/problem=9
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
// a2 + b2 = c2
// For example, 32 + 42 = 9 + 16 = 25 = 52.
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func main() {
	var a, b, c float64
	for i := 1; i < 1001; i++ {
		for j := i + 1; j < 1001; j++ {
			root := math.Sqrt((math.Pow(float64(i), 2)) + math.Pow(float64(j), 2))
			if (float64(i) + float64(j) + root) == 1000 {
				a = float64(i)
				b = float64(j)
				c = root
			}
		}
	}
	fmt.Printf("The product is %d\n", int(a*b*c))
}
