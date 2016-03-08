/*
Projecteuler problem 1 https://projecteuler.net/problem=1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import "fmt"

func main() {
	var tosum []int
	for i := 0; i < 1000; i++ { //the question indicates that it does not consider 0 to be a natural number but adding it to the sum doesn't change the outcome.
		if (i%3 == 0) || (i%5 == 0) {
			tosum = append(tosum, i)
		}
	}
	var thesum int
	for i := 0; i < len(tosum); i++ {
		thesum += tosum[i]
	}
	fmt.Println(thesum)
}
