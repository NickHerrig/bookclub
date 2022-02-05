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
