package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("mrrubiks.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	diary := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, ",")
		task := strings.TrimSpace(lineSlice[1])
		progress, err := strconv.Atoi(strings.TrimSpace(lineSlice[2]))
		if err != nil {
			log.Fatal(err)
		}
		diary[task] += progress
	}

	for task, _ := range diary {
		fmt.Printf("%s\t", task)
	}
	fmt.Print("\n")
	for _, result := range diary {
		stars := strings.Repeat("*", result)
		fmt.Printf("[%s]\t\t", stars)
	}
	fmt.Print("\n")

}
