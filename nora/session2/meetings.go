package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// First we must read in the CSV
	meetings, err := readCSV()
	if err != nil {
		os.Exit(1)
	}	

	//Next we extract the start and end times in separate, sorted lists
	startTimes, endTimes, err := extractSortedStartAndEndTimes(meetings)
	if err != nil {
		os.Exit(1)
	}

	//We'll keep track of the current number of open meetings, as well as the largest number of concurrent meetings that we have seen yet	
	maxConcurrentMeetings := 0
	currentConcurrentMeetings := 0

	//We'll iterate until there are no more meetings to start, as starting a meeting is the only way to require additional rooms
	for len(startTimes) > 0 {
		//We get the next time a meeting will start and the next time a meeting will end from our sorted lists
		currentStartTime := startTimes[0]
		currentEndTime := endTimes[0]

		//Since meeting endings are exclusive, we check them first. As the room is vacated we note this in our count
		if currentEndTime <= currentStartTime {
			endTimes = endTimes[1:]
			currentConcurrentMeetings--
		} else {	
			//if no meetings are ending in this moment, then we will start the next meeting. As the room becomes occupied we note this in our count and, if necessary, record the new maximum number of meetings
			startTimes = startTimes[1:]
			currentConcurrentMeetings++
			if currentConcurrentMeetings > maxConcurrentMeetings {
				maxConcurrentMeetings = currentConcurrentMeetings
			}
		}
	}
	//And we're done!
	fmt.Println(maxConcurrentMeetings)
}

func readCSV() ([][]string, error) { 
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return  reader.ReadAll()
}

func extractSortedStartAndEndTimes(meetings [][]string) ([]int, []int, error) {
	numberOfMeetings := len(meetings)
	startTimes := make([]int, numberOfMeetings, numberOfMeetings)
	endTimes := make([]int, numberOfMeetings, numberOfMeetings)

	for index, meeting := range meetings {
		if len(meeting) < 2 {
			return nil, nil, errors.New("The provided CSV is malformed")
		}
		startTime, err := strconv.Atoi(meeting[0])
		if err != nil {
			return nil, nil, err
		}
		endTime, err := strconv.Atoi(meeting[1])
		if err != nil {
			return nil, nil, err
		}

		if endTime <= startTime {
			return nil, nil, errors.New("There exists a meeting record with a duration of zero or less")
		}
		startTimes[index] = startTime
		endTimes[index] = endTime
	}

	sort.Ints(startTimes)
	sort.Ints(endTimes)
	return startTimes, endTimes, nil

}
