package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type meeting struct {
	start int
	end   int
	name  string
}

func MapArrayToMeeting(data []string) *meeting {
	m := new(meeting)
	m.start, _ = strconv.Atoi(data[0])
	m.end, _ = strconv.Atoi(data[1])
	m.name = data[2]
	return m
}

func main() {
	input_data_file_loc_arg := flag.String("input", "week3_data.csv", "input data location")
	flag.Parse()

	input_data_file_loc := *input_data_file_loc_arg

	f, _ := os.Open(input_data_file_loc)

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, _ := csvReader.ReadAll()

	hours := make(map[string]int)

	for i := 0; i < len(data); i++ {
		meet := MapArrayToMeeting(data[i])
		for j := meet.start; j < meet.end; j++ {
			hours[strconv.Itoa(j)]++
		}
	}

	max_rooms := 0
	for _, val := range hours {
		if val > max_rooms {
			max_rooms = val
		}
	}

	fmt.Println("Max number of rooms needed:", max_rooms)
}
