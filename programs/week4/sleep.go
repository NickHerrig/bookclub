package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

// read each record and convert types as needed
func parse(data io.Reader) ([][]string, error) {
	r := csv.NewReader(data)
	r.TrimLeadingSpace = true

	var records [][]string
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if len(record) != 2 {
			return nil, fmt.Errorf("invalid record")
		}

		records = append(records, record)
	}

	return records, nil
}

func normalize(t string) (time.Time, error) {
	parts := strings.Split(t, " ")
	when, suffix := parts[0], parts[1]
	if len(when) <= 2 {
		when += ":00"
	}

	date, err := time.Parse(time.Kitchen, when + suffix)
	return date, err
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
	records, err := parse(file)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	minutes := 0.0
	total := len(records) - 1
	for _, record := range records[1:] {
		start, err := normalize(record[0])
		if err != nil {
			fmt.Println(err)
			return 1
		}

		end, err := normalize(record[1])
		if err != nil {
			fmt.Println(err)
			return 1
		}

		if end.Before(start) {
			end = end.AddDate(0, 0, 1)
		}

		mins := end.Sub(start).Minutes()
		mins = math.Abs(mins)
		minutes += mins
	}

	// take average and convert to hours
	avg := minutes / float64(total)
	avg /= 60

	fmt.Printf("%2.1f hours\n", avg)
	return 0
}

func main() {
	os.Exit(run())
}
