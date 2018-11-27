package main

import (
	"fmt"
	"math"
)

// The sum of the squares of the first ten natural numbers is,
// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func main() {
	natural, squared := 0, 0
	for i := 1; i < 101; i++ {
		natural += i
		squared += int(math.Pow(float64(i), 2))
	}
	fmt.Printf("The difference is %d\n", (int(math.Pow(float64(natural), 2)) - squared))
}
