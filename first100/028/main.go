/*
Projecteuler problem 28 https://projecteuler.net/problem=28

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
*/
package main

import "fmt"

func main() {
	i, s, w, sum := 1, 0, 1, 1
	for {
		for n := 0; n < 5; n++ {
			if n < 2 {
				s++
			}
			if n == 0 {
				w += 2
			} else {
				i += s
				sum += i
			}
		}
		if w == 1001 {
			break
		}
	}
	fmt.Println(sum)
}
