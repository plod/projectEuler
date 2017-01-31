/*
Projecteuler problem 22 https://projecteuler.net/problem=22

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/
package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"sort"
)

func getScore(name string) int {
	values := []byte(name)
	var score int
	for _, theByte := range values {
		singleScore := int(theByte) - 64
		score += singleScore
		//log.Println(singleScore)
	}
	return score
}

func main() {
	src, err := os.Open("p022_names.txt")
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	r := csv.NewReader(bufio.NewReader(src))
	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(sort.StringSlice(record))
	/*
		  getColin := 938
			log.Println(record[getColin-1])
			log.Println(getScore(record[getColin-1]))
	*/
	var runningScore int
	for i := 0; i < len(record); i++ {
		runningScore += (i + 1) * getScore(record[i])
	}
	log.Println(runningScore)
}
