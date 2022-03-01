package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Goal struct {
	Date        time.Time
	Description string
	Success     int
}

type Chart struct {
	walk  int
	clean int
	sort  int
}

func NewChart() *Chart {
	c := Chart{}
	return &c
}

func (c *Chart) Walk(success bool) {
	if success {
		c.walk += 1
	}
}

func (c *Chart) Clean(success bool) {
	if success {
		c.clean += 1
	}
}

func (c *Chart) Sort(success bool) {
	if success {
		c.sort += 1
	}
}

func (c *Chart) Report() string {
	report := ""
	report += "walk:\n"
	report += strings.Repeat("*", c.walk) + "\n"
	report += "clean:\n"
	report += strings.Repeat("*", c.clean) + "\n"
	report += "sort:\n"
	report += strings.Repeat("*", c.sort) + "\n"
	return report
}

// read each record and convert types as needed
func parse(data io.Reader) ([]Goal, error) {
	r := csv.NewReader(data)
	r.TrimLeadingSpace = true

	var goals []Goal
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(record) != 3 {
			return nil, fmt.Errorf("invalid record")
		}

		date, err := time.Parse("01-02-06", record[0])
		if err != nil {
			return nil, err
		}

		success, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}

		goal := Goal{
			Date: date,
			Description: record[1],
			Success: success,
		}
		goals = append(goals, goal)
	}

	return goals, nil
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
	goals, err := parse(file)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	chart := NewChart()
	for _, goal := range goals {
		desc := strings.ToLower(goal.Description)
		if strings.Contains(desc, "walk") {
			success := false
			if goal.Success == 1 {
				success = true
			}
			chart.Walk(success)
		}
		if strings.Contains(desc, "clean") {
			success := false
			if goal.Success == 1 {
				success = true
			}
			chart.Clean(success)
		}
		if strings.Contains(desc, "sort") {
			success := false
			if goal.Success == 1 {
				success = true
			}
			chart.Sort(success)
		}
	}

	fmt.Println(chart.Report())
	return 0
}

func main() {
	os.Exit(run())
}
