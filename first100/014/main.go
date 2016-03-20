/*
Projecteuler problem 14 https://projecteuler.net/problem=14

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/
package main

import "log"

func evenFunc(num int) int {
	return num / 2
}

func oddFunc(num int) int {
	return (3 * num) + 1
}
func main() {
	var largestNumber, largestChain int
	for i := 1e6; i > 2; i-- {
		var chainLength, runningNumber int
		for runningNumber = int(i); runningNumber > 1; chainLength++ {
			if runningNumber%2 == 0 {
				runningNumber = evenFunc(runningNumber)
			} else {
				runningNumber = oddFunc(runningNumber)
			}
		}
		if chainLength > largestChain {
			largestChain = chainLength + 1 //lose the last iteration
			largestNumber = int(i)
			log.Printf("new largest chain of length %d found for number %d", largestChain, largestNumber)
		}
	}
}
