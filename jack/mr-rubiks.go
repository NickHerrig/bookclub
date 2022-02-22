package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Diary struct {
  Walk int
  Sort int
  Clean int
}

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)
    diary := Diary {
      Walk: 0,
      Sort: 0,
      Clean: 0,
    }

    for fileScanner.Scan() {
      line := fileScanner.Text()
      task := strings.Split(line, " ")[1]
      value, _ := strconv.Atoi(strings.Split(line, " ")[3])
      if task == "Clean" {
        diary.Clean += value
        if diary.Clean < 0 {
          diary.Clean = 0
        }
      }
      if task == "Walk" {
        diary.Walk += value
        if diary.Walk < 0 {
          diary.Walk = 0
        }
      }
      if task == "Sort" {
        diary.Sort += value
        if diary.Sort < 0 {
          diary.Sort = 0
        }
      }
    }

    walks := strings.Repeat("*", diary.Walk)
    cleans := strings.Repeat("*", diary.Clean)
    sorts := strings.Repeat("*", diary.Sort)
    // These are not lined up correctly. 
    fmt.Println("walk, clean, sort")
    fmt.Printf("[%s] [%s] [%s]\n", walks, cleans, sorts)
}
