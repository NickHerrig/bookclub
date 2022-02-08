package main

import (
    "os"
    "encoding/csv"
    "fmt"
    "strconv"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()

    var start_times []int
    var end_times []int

    for index, _ := range records {
        start_time, _ := strconv.Atoi(records[index][0])
        end_time, _ := strconv.Atoi(records[index][1])
        start_times = append(start_times, start_time)
        end_times = append(end_times, end_time)
    }

    //fmt.Println(start_times)
    //fmt.Println(end_times)
    ctr := 0
    time_ctr := 1
    for i:=0; i<len(end_times)-1; i++ {
        if start_times[time_ctr] <= end_times[i] && end_times[time_ctr] > start_times[i]{
            ctr = ctr+1
        }
        time_ctr = time_ctr+1
    }
    fmt.Println(ctr)
}
