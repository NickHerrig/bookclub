/*
Bookclub Example Program #3:

Mr. Rubiks is in a pickle. He has a 3x3 grid that he uses to track his goals, but his dog (Fruruf) knocked it off the wall! But luckily Mr. Rubiks is a data hoarder, and keeps his goal progress recorded in his diary.

Mr. Rubiks has 3 goals, and each time he successfully works on one, he adds a star to the column under that goal, and when he gets 3 stars under a goal he has mastered it! Being a critic though, when he fails to make progress on a goal in a day, he removes a star (up to empty under the goal). He's a busy man though, and days where he doesn't have time for working on himself, he doesn't penalize himself (no updates to the grid).

In addition to updating the chart, each night Mr. Rubiks updates his diary with the date, the goal to progress, and if he succeeded in making progress or not.

Given Mr. Rubiks diary, and his current three goals:

    Walk Fruruf right, left, right, right, left, then reverse around the block
    Sort his colored cubes according to color
    Clean up his puzzling station

Help Mr. Rubiks reassemble his goal chart and report back the finished chart.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csv := flag.String("csv", "goals.csv", "csv file")
	flag.Parse()

	// load in csv
	file, err := os.Open(*csv)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var goals [][]string
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ",")
		goals = append(goals, res)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// loop over entries
	var chart [3][3]string
	for _, goal := range goals {

		// look for different goals
		switch strings.Trim(goal[1], " ") {
		case "Walk Fruruf":
			for i, progress := range chart[0] {
				if strings.Trim(goal[2], " ") == "1" {
					if progress != "" {
						continue
					} else {
						chart[0][i] = "x"
						break
					}
				} else {
					if progress != "" {
						continue
					} else {
						if i == 0 {
							break
						}
						chart[0][i-1] = ""
						break
					}
				}
			}
		case "Clean Puzzles":
			for i, progress := range chart[1] {
				if strings.Trim(goal[2], " ") == "1" {
					if progress != "" {
						continue
					} else {
						chart[1][i] = "x"
						break
					}
				} else {
					if progress != "" {
						continue
					} else {
						if i == 0 {
							break
						}
						chart[1][i-1] = ""
						break
					}
				}
			}
		case "Sort Cubes":
			for i, progress := range chart[2] {
				if strings.Trim(goal[2], " ") == "1" {
					if progress != "" {
						continue
					} else {
						chart[2][i] = "x"
						break
					}
				} else {
					if progress != "" {
						continue
					} else {
						if i == 0 {
							break
						}
						chart[2][i-1] = ""
						break
					}
				}
			}
		}
	}
	fmt.Println("walk clean sort")
	for i := 0; i < 3; i++ {
		fmt.Printf(" [%v]  [%v]  [%v]\n", chart[0][i], chart[1][i], chart[2][i])
	}
}

// Evaluate does not work
func Evaluate(goal int, entry []string, chart [3][3]string) {
	for i, progress := range chart[goal] {
		if strings.Trim(entry[2], " ") == "1" {
			if progress != "" {
				continue
			} else {
				chart[goal][i] = "x"
				break
			}
		} else {
			if progress != "" {
				continue
			} else {
				if i == 0 {
					break
				}
				chart[goal][i-1] = ""
				break
			}
		}
	}
}
