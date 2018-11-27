package main

import (
	"fmt"
)

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func main() {
	n := 20
	for {
		divisible := false
		for i := 20; i > 0; i-- {
			if n%i != 0 {
				break
			}
			if i == 1 {
				divisible = true
			}
		}
		if divisible == true {
			break
		}
		n += 20
	}
	fmt.Printf("The smallest divisible number is %d\n", n)
}
