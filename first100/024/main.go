/*
Projecteuler problem 24 https://projecteuler.net/problem=24

A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/
package main

import "log"

func main() {
	n := 10
	i := 1000000
	j, k := 0, 0
	fact := make([]int, n)
	perm := make([]int, n)
	fact[k] = 1
	k++
	i--
	for k < n {
		fact[k] = fact[k-1] * k
		k++
	}
	for k = 0; k < n; k++ {
		perm[k] = i / fact[n-1-k]
		i = i % fact[n-1-k]
	}
	for k = n - 1; k > 0; k-- {
		for j = k - 1; j >= 0; j-- {
			if perm[j] <= perm[k] {
				perm[k]++
			}
		}
	}
	result := 0
	for k = 0; k < n; k++ {
		result = (result + perm[k]) * 10
	}
	log.Println(result / 10)
}
