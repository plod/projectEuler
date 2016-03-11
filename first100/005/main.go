/*
Projecteuler problem 5 https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/
package main

import "fmt"

const numberOfConsecutiveFactors = 20

func isItDivisible(number int) bool {
	for i := 1; i <= numberOfConsecutiveFactors; i++ {
		if number%i == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}
func main() {
	keepgoing := true
	var smallestnumber int
	for i := 1; keepgoing == true; i++ {
		if isItDivisible(i) {
			smallestnumber = i
			keepgoing = false
		}
	}
	fmt.Println(smallestnumber)
}
