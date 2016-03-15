/*
Projecteuler problem 9 https://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var sum float64
	found := false
	var a, b int
	var c float64
	for a = 1; a < 333; a++ { //333+333+334 > 1000
		for b = a + 1; sum < 1000; b++ {
			c = math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			sum = float64(a) + float64(b) + c
			//	fmt.Printf("a = %v, b = %v, c = %v sum = %v\n", a, b, c, sum)
			if c == float64(int64(c)) && sum == 1000 {
				found = true
				break
			}
		}
		if found == true {
			break
		} else {
			sum = 0
		}
	}
	if found == true {
		fmt.Println("========")
		fmt.Printf("a = %d, b = %d, c = %d is the tuple\n", a, b, int(c))
		fmt.Printf("abc product = %d\n", a*b*int(c))
	}
}
