package main

import (
	"fmt"
	"os"

	"github.com/ernestosuarez/itertools"
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	meetings := []*Meeting{}
	err := gocsv.UnmarshalString(csv, &meetings)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	indexs := []int{}
	for i := 0; i < len(meetings); i++ {
		indexs = append(indexs, i)
	}

	combos := itertools.CombinationsInt(indexs, 2)
	overlaps := 0
	for v := range combos {
		one, two := v[0], v[1]
		overlap := min(meetings[one].End, meetings[two].End) - max(meetings[one].Start, meetings[two].Start) + 1
		if overlap > 0 {
			overlaps++
		}
	}
}
