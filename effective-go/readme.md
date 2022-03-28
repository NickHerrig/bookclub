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

## Week Six 03/23
### Leader
Nick Herrig

### Readings
- Read [The blank identifier](https://go.dev/doc/effective_go#blank) and [Embedding](https://go.dev/doc/effective_go#embedding)

### Problem
Slick Steve has an awesome idea for a web application!
His first step, is to get a simple HTML page hosted on his server.
Steve has a problem though, he works at a pretty locked down organization.
This organization only gives him the ability to deploy a single binary!

Good thing you know about embed!
Help steve deploy a single binary that serves this basic HTML page.

Bonus Points: for embedding types where possible for this problem.

Example input:
```html
<!DOCTYPE html>
<html>

  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE-edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>TTDSM - Book Club</title>
  </head>

  <body>
    <main>
    <h1>Hello from TTDSM</h1>
    <button type="button">Do Not Press This!</button>
    </main>
  </body>

</html>
```

Successful Problem:

Visiting {localhost}:{port} returns the html file, logs the date, and IP address of visitor
```
go build prog.go
./prog
APP: 2022/03/12 192.168.0.36:54259 - New friend visited site.
```

## Week Seven 04/06
### Leader
Andrew Dailey

### Readings
- Read [Concurrency](https://go.dev/doc/effective_go#concurrency)

### Problem
O Python, Where Art Thou? It is with constant sorrow that I often attempt to write Python constructs in Go. Simply out of habit, I expect to be able to write clean `for i in range(5)` loops. Sadly, this ain't no big rock candy mountain. Best we can do is an explicit for loop... right?

I want bring back the glory of Python's `range` builtin. This week, write a function called `xrange` (since `range` is a Go keyword) that accepts an integer N and returns a channel that produces sequential integers up to (but not including) N. Here is the signature:
```go
func xrange(n int) chan int {
    // TODO: implement this (hint, it'll require a goroutine)
    return nil
}
```

Here is an example:
```go
package main

import (
	"fmt"
)

func main() {
	for i := range xrange(5) {
		fmt.Println(i)
	}
}
```

And the output:
```
0
1
2
3
4
```
