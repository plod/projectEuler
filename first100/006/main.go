/*
Projecteuler problem 6 https://projecteuler.net/problem=5

The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

*/
package main

import (
	"fmt"
	"math"
)

const naturalNumber = 100

func main() {
	var indivSquared int
	var groupSquared int
	var groupSummed int
	for i := 1; i <= naturalNumber; i++ {
		indivSquared += int(math.Pow(float64(i), 2))
		groupSummed += i
	}
	groupSquared = int(math.Pow(float64(groupSummed), 2))
	fmt.Printf("The indiv squared sum is %d\n", indivSquared)
	fmt.Printf("The sum is %d\n", groupSummed)
	fmt.Printf("The sum squared is %d\n", groupSquared)
	fmt.Println("--------------------")
	fmt.Printf("Finally the difference is %d\n", groupSquared-indivSquared)
}
