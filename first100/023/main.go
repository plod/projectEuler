/*
Projecteuler problem 23 https://projecteuler.net/problem=23

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/
package main

import "log"

func sumOfDivisors(n int) int {
	sum, p := 1, 2
	for p*p <= n && n > 1 {
		if n%p == 0 {
			j := p * p
			n /= p
			for n%p == 0 {
				j *= p
				n /= p
			}
			sum *= j - 1
			sum /= p - 1
		}
		if p == 2 {
			p = 3
		} else {
			p += 2
		}
	}
	if n > 1 {
		sum *= n + 1
	}
	return sum
}

func sumOfProperDivisors(n int) int {
	return sumOfDivisors(n) - n
}

func isAbundant(n int) bool {
	if n < sumOfProperDivisors(n) {
		return true
	}
	return false
}
func main() {
	limit := 28124
	abundants := make([]int, 6965)
	ai := 0
	for i := 2; i < limit; i++ {
		if isAbundant(i) {
			abundants[ai] = i
			ai++
		}
	}
	sumOfAbundants := make([]bool, limit)
	for _, v := range abundants {
		for _, w := range abundants {
			if sum := v + w; sum < limit {
				sumOfAbundants[sum] = true
			}
		}
	}
	sum := 0
	for i := range sumOfAbundants {
		if !sumOfAbundants[i] {
			sum += i
		}
	}
	log.Println(sum)
}
