/*
Projecteuler problem 3 https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/
package main

import "fmt"

var primes []int

// A concurrent prime sieve
// taken from http://play.golang.org/p/9U22NfrXeq I kinda understand it but
//not grasped channesl properly yet

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Launch Generate goroutine.
	limit := 600851475143
	runninglimit := limit
	for i := 0; i < runninglimit; i++ {
		prime := <-ch
		if limit%prime == 0 {
			primes = append(primes, prime)
			//reduce the iterations
			if runninglimit%prime == 0 {
				runninglimit = runninglimit / prime
			} else if (limit / prime) < runninglimit {
				runninglimit = limit / prime
			}
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println(primes[len(primes)-1])
}
