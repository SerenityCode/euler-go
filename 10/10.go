package main

import (
	"fmt"
	"math"
)

// Summation of primes https://projecteuler.net/problem=10
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.

func isPrime(number int) bool {
	prime := true
	for i := 3; i < int(math.Sqrt(float64(number)))+1; i += 2 {
		if number%i == 0 {
			prime = false
			break
		}
	}
	return prime
}

func main() {
	primes := []int{2}
	sum := 0
	for i := 3; i < 2000000; i += 2 {
		if isPrime(i) == true {
			primes = append(primes, i)
		}
	}
	for _, num := range primes {
		sum += num
	}
	fmt.Printf("The sum is %d\n", sum)
}
