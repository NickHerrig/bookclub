package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Meeting struct {
	Start int
	End   int
	Name  string
}

func (m Meeting) Occupies(hour int) bool {
	return hour >= m.Start && hour < m.End
}

type Schedule []Meeting

func (s Schedule) Len() int {
	return len(s)
}

func (s Schedule) Less(i, j int) bool {
	return s[i].End < s[j].End
}

func (s Schedule) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// read each record and convert types as needed
func parse(data io.Reader) (Schedule, error) {
	r := csv.NewReader(data)

	var meetings []Meeting
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) != 3 {
			return nil, fmt.Errorf("invalid record")
		}

		start, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}

		meeting := Meeting{
			Start: start,
			End:   end,
			Name:  record[2],
		}
		meetings = append(meetings, meeting)
	}

	return Schedule(meetings), nil
}

func solveLinear(schedule Schedule) int {
	rooms := 0
	for hour := 0; hour < 24; hour++ {
		// check how many meetings occur this hour
		concurrent := 0
		for _, meeting := range schedule {
			if meeting.Occupies(hour) {
				concurrent += 1
			}
		}

		// found a larger count of concurrent meetings
		if concurrent > rooms {
			rooms = concurrent
		}
	}

	return rooms
}

func solveSorted(schedule Schedule) int {
	// sort meetings by end time
	sort.Sort(schedule)

	// track meetings running at the same time
	rooms := 0
	for i, cur := range schedule {
		concurrent := 1
		for _, next := range schedule[i+1:] {
			if next.Start < cur.End {
				concurrent += 1
			}
		}

		if concurrent > rooms {
			rooms = concurrent
		}
	}

	return rooms
}

func run() int {
	data := flag.String("data", "data.csv", "data file")
	flag.Parse()

	// open the CSV file
	file, err := os.Open(*data)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	defer file.Close()

	// parse the CSV file
	schedule, err := parse(file)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Printf("linear: %d\n", solveLinear(schedule))
	fmt.Printf("sorted: %d\n", solveSorted(schedule))

	return 0
}

func main() {
	os.Exit(run())
}
