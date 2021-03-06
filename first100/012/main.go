/*
Projecteuler problem 12 https://projecteuler.net/problem=12

The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?

*/

package main

import (
	"log"
)

var (
	primeList = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	lastN     = 23
)

func whatIsTriangle(n int) int {
	return (n * (n + 1)) / 2
}

func isPrime(n int) bool {
	isItPrime := false
	for i := 0; i < len(primeList); i++ {
		if primeList[i]*primeList[i] > n {
			break
		}
		if (n % primeList[i]) != 0 {
			isItPrime = true
			break
		}
	}
	return isItPrime
}

func prime(x int) int {
	for len(primeList) <= x {
		lastN = lastN + 1
		if isPrime(lastN) {
			primeList = append(primeList, lastN)
		}
	}
	return primeList[x]
}

func numFactors(n int) int {
	div := 1
	x := 0
	for n > 1 {
		var c = 1
		for (n % prime(x)) == 0 {
			c = c + 1
			n = n / prime(x)
		}
		x = x + 1
		div = div * c
	}
	return div
}

func main() {
	var triangleNumber, numberOfFactors int
	for i := 1; numberOfFactors <= 500; i++ {
		triangleNumber = whatIsTriangle(i)
		numberOfFactors = numFactors(triangleNumber)
	}
	log.Printf("The value of the first triangle number to have over five hundred divisors is %d", triangleNumber)
}
