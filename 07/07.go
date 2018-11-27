package main

import (
	"fmt"
	"math"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?

func factorList(number int) []int {
	var factors []int
	for i := 3; i < int(math.Sqrt(float64(number)))+1; i += 2 {
		if number%i == 0 {
			factors = append(factors, i, number/i)
		}
	}
	return factors
}

func main() {
	primes := []int{2}
	number := 3
	for len(primes) < 10001 {
		if len(factorList(number)) == 0 {
			primes = append(primes, number)
		}
		number += 2
	}
	fmt.Printf("The 10001st prime number is %v", primes[len(primes)-1])
}
