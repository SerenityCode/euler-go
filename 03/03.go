package main

import (
	"fmt"
	"math"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func factorList(number int) []int {
	var factors []int
	for i := 3; i < int((math.Sqrt(float64(number) + 2))); i++ {
		if number%i == 0 {
			factors = append(factors, i, number/i)
		}
	}
	return factors
}

func main() {
	factors := factorList(600851475143)
	largestPrime := 0
	for _, value := range factors {
		if len(factorList(value)) == 0 {
			if value > largestPrime {
				largestPrime = value
			}
		}
	}
	fmt.Printf("The largest prime is %d\n", largestPrime)
}
