package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

const csv = `To Bed, Wake Up
	     9:30 PM, 6:15 AM
	     9:00 PM, 4:00 AM
	     12:15 AM, 6:20 AM
	     4:00 PM, 3:00 AM`

type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalCSV(csv string) (err error) {
	r := strings.Trim(csv, " \t")
	d.Time, err = time.Parse("3:04 PM", r)
	return err
}

type Record struct {
	Retire DateTime `csv:"To Bed"`
	Wake   DateTime `csv:"Wake Up"`
}

func (r *Record) Duration() time.Duration {
	dur := r.Wake.Time.Sub(r.Retire.Time)
	if dur < 0 {
		return 24*time.Hour + dur
	} else {
		return dur
	}
}

func main() {
	records := []*Record{}
	err := gocsv.UnmarshalString(csv, &records)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	var sum float64
	for _, r := range records {
		sum += r.Duration().Hours()
	}
	avg := sum / float64(len(records))
	fmt.Printf("%.1f hours\n", avg)

}
