/*
Projecteuler problem 21 https://projecteuler.net/problem=21

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

not correct work in progress
*/
package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func d(n int) int {
	var divisors []int
	var dReturn int
	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			divisors = append(divisors, i)
			dReturn += i
		}
	}
	return dReturn
}

func main() {
	var sum int
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(iInside int) {
			c := d(iInside)
			if c != iInside && d(c) == iInside {
				mutex.Lock()
				sum += iInside
				mutex.Unlock()
				log.Println("New amicable number found", iInside)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("The sum is", sum)
}
