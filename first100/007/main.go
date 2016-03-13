/*
Projecteuler problem 7 https://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

*/
package main

import "fmt"

var primes []int

const finalNumber = 10001

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
	for i := 0; len(primes) < finalNumber; i++ {
		prime := <-ch
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	/*for i, prime := range primes {
		fmt.Printf("%d:\t%d\n", i, prime)
	}*/
	fmt.Println(primes[len(primes)-1])

}
