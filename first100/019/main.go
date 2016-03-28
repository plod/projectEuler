/*
Projecteuler problem 19 https://projecteuler.net/problem=19
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var sundayCount int
	for i := 1901; i <= 2000; i++ {
		for i2 := 1; i2 <= 12; i2++ {
			date := time.Date(i, time.Month(i2), 1, 0, 0, 0, 0, time.UTC)
			log.Printf("Go launched at %s\n", date)
			weekdayString := fmt.Sprintf("%s", date.Weekday())
			if weekdayString == "Sunday" {
				sundayCount++
			}
		}
	}
	log.Println(sundayCount)
}

/* i don't like this answer but seems to simple to not do it this way */
