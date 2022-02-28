# Schedule
This doc holds the reading schedule, related blogs/docs and summaries for each week

## Week One 01/12
### Leader
Nick Herrig

### Readings
- Read Through [commentary](https://go.dev/doc/effective_go#commentary)

### Todos
- Download [Go Programming Language](https://go.dev/dl/)

### Related Blogs/Docs
- https://benhoyt.com/writings/go-intro/
- [Tour of Go](https://go.dev/tour/welcome/1)


## Week Two 01/26
### Leader
Andrew Dailey

### Readings
- Read through [Control Structures](https://go.dev/doc/effective_go#control-structures)

### Problem 
Did you just fall into an Escher painting?
Are you in Egypt?
Things have definitely gotten sideways and pointy.
There is clearly a pyramid off in the distance, but it appears to be flipped!
Also, its size appears to vary based on how hard you squint (don't wanna get too much sand in your eyes).

Write a program that accepts a single integer argument (must be greater than zero).
Based on this value, print out a sideways pyramid of asterisks that peaks at the provided size.

Here are a few examples:
```
go run pyramids.go 1
*

go run pyramids.go 2
*
**
*

go run pyramids.go 5
*
**
***
****
*****
****
***
**
*
```

## Week Three 02/09
### Leader
Nora Mayse

### Readings
- Read [Functions](https://go.dev/doc/effective_go#functions) and [Data](https://go.dev/doc/effective_go#data)

### Problem
Modernizing the world of meeting spaces sure is hard work. 
Our drone delivery tech can bring our eco friendly MeetHubs, made from international shipping crates, 
anywhere in the continental United States! We're kind of the Uber of the shipped temporary meeting room world. 
However, we have hit a snag. We currently guess the correct number of MeetHubs to drop off! 
It would be much better if we used our clients meeting schedules to determine exactly how many MeetHubs 
that they need for a given day.

Given a csv in which the 0th column is start time (inclusive),  the 1th column is end time (exclusive), 
and the 2th column is meeting name, please provide the number of meeting rooms required for the time period.

*Example:*

```csv
9,12,mob session
8,11,better team's mob session
13,16,afternoon nap
```

```shell
go run schedules.go data.csv
2
```

## Week Four 02/23
### Leader
Wesley Brueland

### Readings
- Read [Initialization](https://go.dev/doc/effective_go#initialization) and [Methods](https://go.dev/doc/effective_go#methods)

### Problem
Mr. Rubiks is in a pickle. He has a 3x3 grid that he uses to track his goals, but his dog (Fruruf) knocked it off the wall! But luckily Mr. Rubiks is a data hoarder, and keeps his goal progress recorded in his diary. 

Mr. Rubiks has 3 goals, and each time he successfully works on one, he adds a star to the column under that goal, and when he gets 3 stars under a goal he has mastered it! Being a critic though, when he fails to make progress on a goal in a day, he removes a star (up to empty under the goal). He's a busy man though, and days where he doesn't have time for working on himself, he doesn't penalize himself (no updates to the grid). 

In addition to updating the chart, each night Mr. Rubiks updates his diary with the date, the goal to progress, and if he succeeded in making progress or not. 

Given Mr. Rubiks diary, and his current three goals:

1. Walk Fruruf right, left, right, right, left, then reverse around the block
2. Sort his colored cubes according to color
3. Clean up his puzzling station

Help Mr. Rubiks reassemble his goal chart and report back the finished chart.

TEST INPUT:
```
05-11-21, Walk Fruruf, 1
05-13-21, Walk Fruruf, -1
05-14-21, Clean Puzzles, -1
05-14-21, Sort Cubes, 1
05-15-21, Walk Fruruf, 1
05-17-21, Walk Fruruf, 1
05-17-21, Clean Puzzles, 1
05-19-21, Sort Cubes, 1
05-20-21, Clean Puzzles, 1
05-21-21, Sort Cubes, 1
```

As an example:
```
05-11-21, Walk Fruruf, 1
```
 
causes the state:
```
walk, clean, sort
[*]   []     []
```

and losing progress looks like:
```
05-12-21, Walk Fruruf, -1

walk, clean, sort
[]    []     []
```

## Week Five 03/9
### Leader
Jack Fordyce 

### Readings
- Read [Interfaces](https://go.dev/doc/effective_go#interfaces_and_types)

### Problem
Mr. S is a terrible sleeper.

For years, Mr. S has struggled maintaining a healthy and regular sleep schedule. In a last ditch effort, he has started recording his “to bed” times and “wake up” times every day to try to get an idea of how much sleep he is getting and when he typically goes to bed and wakes up. Help Mr. S improve his sleep by calculating his average hours (round to nearest tenth) slept over n number of days. His entries look as follows:

```
To Bed, Wake Up
9:30 PM, 6:15 AM
9 PM, 4 AM
12:15 AM, 6:20 AM
4 PM, 3 AM
```

Example Output:
```
8.2 hours
```
