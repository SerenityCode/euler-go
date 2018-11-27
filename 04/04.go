package main

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func strReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(number int) bool {
	if strconv.Itoa(number) == strReverse(strconv.Itoa(number)) {
		return true
	}
	return false
}

func main() {
	largest := 0
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) {
				if product > largest {
					largest = product
				}
			}
		}
	}
	fmt.Printf("The largest palindrome is %d\n", largest)
}
