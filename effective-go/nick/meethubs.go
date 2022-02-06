package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/gocarina/gocsv"
)

const csv = `start,end,meeting_name
             9,12,mob session
	     8,11,better team's mob session
             13,16,afternoon nap`

type Meeting struct {
	Start int    `csv:"start"`
	End   int    `csv:"end"`
	Name  string `csv:"meeting_name"`
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMeetingRooms(startTimes, endTimes []int) int {
	if !sort.IntsAreSorted(startTimes) {
		sort.Ints(startTimes)
	}
	if !sort.IntsAreSorted(endTimes) {
		sort.Ints(endTimes)
	}

	minRooms, onGoing := 0, 0

	for i, j := 0, 0; i < len(startTimes) && j < len(endTimes); {
		if startTimes[i] < endTimes[j] {
			onGoing++
			minRooms = max(onGoing, minRooms)
			i++
		} else {
			j++
			onGoing--
		}
	}

	return minRooms
}

func main() {
	meetings := []*Meeting{}
	err := gocsv.UnmarshalString(csv, &meetings)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	// build a slice of start and end time slices
	var startTimes, endTimes []int
	for _, meeting := range meetings {
		startTimes = append(startTimes, meeting.Start)
		endTimes = append(endTimes, meeting.End)
	}

	fmt.Println(findMeetingRooms(startTimes, endTimes))

}
