/*
Bookclub Example Program #2:

Modernizing the world of meeting spaces sure is hard work. Our drone delivery tech can bring our eco friendly MeetHubs, made from international shipping crates, anywhere in the continental United States! We're kind of the Uber of the shipped temporary meeting room world. However, we have hit a snag. We currently guess the correct number of MeetHubs to drop off! It would be much better if we used our clients meeting schedules to determine exactly how many MeetHubs that they need for a given day.

Given a csv in which the 0th column is start time (inclusive),  the 1th column is end time (exclusive), and the 2th column is meeting name, please provide the number of meeting rooms required for the time period.

For example
go run schedules.go data.csv

where data.csv contains
------------------------------
9,12,mob session
8,11,better team's mob session
13,16,afternoon nap
------------------------------
would return
2
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csv := flag.String("csv", "schedules.csv", "csv file")
	flag.Parse()

	file, err := os.Open(*csv)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bookings [][]string
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ",")
		bookings = append(bookings, res)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= 24; i++ {
		for _, booking := range bookings {
			fmt.Println(booking)
		}
	}
}
