/*
Projecteuler problem 10 https://projecteuler.net/problem=10

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

*/
package main

/*
borrowed from: http://blog.idempotent.ca/2013/11/20/first-dip-into-golangs-concurrency/ but was getting there
but couldnt work out how to do same concurrantly learned a lot more about go's multithreading
*/

import (
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
)

var workers = runtime.NumCPU()

func isPrime(n int) bool {
	if n == 1 || n == 2 {
		return true
	}

	if math.Mod(float64(n), 2) == 0 {
		return false
	}

	for i := 3.0; i <= math.Floor(math.Sqrt(float64(n))); i += 2.0 {
		if math.Mod(float64(n), i) == 0 {
			return false
		}
	}

	return true
}

type job struct {
	n      int
	result chan<- int
}

func (job *job) Do() {
	if isPrime(job.n) {
		job.result <- job.n
	}
}

func sumPrimesUpto(n int) int {
	jobs := make(chan job, workers)
	results := make(chan int, n)
	done := make(chan struct{}, workers)

	go addJobs(jobs, results, n)

	for i := 0; i < workers; i++ {
		go doJobs(done, jobs)
	}

	go wait(done, results)

	return tally(results)
}

func tally(results <-chan int) int {
	retval := 0
	for result := range results {
		retval += result
	}

	return retval
}

func addJobs(jobs chan<- job, results chan<- int, n int) {
	for i := 2; i <= n; i++ {
		jobs <- job{i, results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, jobs <-chan job) {
	for job := range jobs {
		job.Do()
	}
	done <- struct{}{}
}

func wait(done <-chan struct{}, results chan int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func main() {
	runtime.GOMAXPROCS(workers)
	log.Printf("CPUS=%d\n", workers)
	upperBound, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println("Invalid argument.")
		os.Exit(1)
	}

	result := sumPrimesUpto(upperBound)
	log.Println(result)
}
