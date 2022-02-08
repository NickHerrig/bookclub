package main

import (
    "os"
    "encoding/csv"
    "fmt"
    "strconv"
    "sort"
)

type Time struct {
    Start int
    End int
}

type sorted_times []Time

func (st sorted_times) Len() int {
    return len(st)
}

func (st sorted_times) Less(i, j int) bool {
    return st[i].Start < st[j].Start
}

func (st sorted_times) Swap(i, j int) {
    st[i], st[j] = st[j], st[i]
}

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()

    var times []Time
    var start_times []int
    var end_times []int

    for index, _ := range records {
        start_time, _ := strconv.Atoi(records[index][0])
        end_time, _ := strconv.Atoi(records[index][1])
        meeting_times := Time {
            Start: start_time,
            End: end_time,
        }
        times = append(times, meeting_times)
        start_times = append(start_times, start_time)
        end_times = append(end_times, end_time)
    }
    sort.Sort(sorted_times(times))
    fmt.Println(times)
    //fmt.Println(start_times)
    //fmt.Println(end_times)
    ctr := 0
    time_ctr := 1
    for i:=0; i<len(times)-1; i++ {
        if times[time_ctr].Start <= times[i].End && times[time_ctr].End > times[i].Start{
            
            ctr = ctr+1
        }
        time_ctr = time_ctr+1
    }
    fmt.Println(ctr)
}
