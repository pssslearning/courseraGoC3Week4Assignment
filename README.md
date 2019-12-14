# courseraGoC3Week4Assignment

Assigment Week #4- Concurrency in Go - THREADS IN GO - Course https://www.coursera.org/learn/golang-concurrency/home/week/4

## Assigment Week #4 - Concurrency in Go - THREADS IN GO  

- [Assignment](https://www.coursera.org/learn/golang-concurrency/peer/MAemV/module-4-activity)


- [Course: Concurrency in Go](https://www.coursera.org/learn/golang-concurrency/home/welcome)
  
## Instructions

The goals of this activity are to explore the design of concurrent algorithms and resulting synchronization issues.

Implement the dining philosopher’s problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

## My assignment

Source code at `philosophers/philosophers.go`

## Sample compilation and first test execution

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoC3Week4Assignment/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week4Assignment$ cd philosophers/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week4Assignment/philosophers$ go build philosophers.go 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week4Assignment/philosophers$ ./philosophers 
------------------------------------------------------------------------------------------------------
MISSION BEGINS
------------------------------------------------------------------------------------------------------
starting to eat  <1> 		( Using Chopsticks: #0 and #1)
starting to eat  <3> 		( Using Chopsticks: #2 and #3)
finishing eating <3> 		( Eating turns consumed: 1. Units of time required: 611 )
finishing eating <1> 		( Eating turns consumed: 1. Units of time required: 818 )
------------------------------------------------------------------------------------------------------
starting to eat  <4> 		( Using Chopsticks: #3 and #4)
starting to eat  <2> 		( Using Chopsticks: #1 and #2)
finishing eating <2> 		( Eating turns consumed: 1. Units of time required: 273 )
finishing eating <4> 		( Eating turns consumed: 1. Units of time required: 967 )
------------------------------------------------------------------------------------------------------
starting to eat  <5> 		( Using Chopsticks: #4 and #0)
starting to eat  <3> 		( Using Chopsticks: #2 and #3)
finishing eating <5> 		( Eating turns consumed: 1. Units of time required: 547 )
finishing eating <3> 		( Eating turns consumed: 2. Units of time required: 611 )
------------------------------------------------------------------------------------------------------
starting to eat  <1> 		( Using Chopsticks: #0 and #1)
starting to eat  <4> 		( Using Chopsticks: #3 and #4)
finishing eating <1> 		( Eating turns consumed: 2. Units of time required: 818 )
finishing eating <4> 		( Eating turns consumed: 2. Units of time required: 967 )
------------------------------------------------------------------------------------------------------
starting to eat  <2> 		( Using Chopsticks: #1 and #2)
starting to eat  <5> 		( Using Chopsticks: #4 and #0)
finishing eating <2> 		( Eating turns consumed: 2. Units of time required: 273 )
finishing eating <5> 		( Eating turns consumed: 2. Units of time required: 547 )
------------------------------------------------------------------------------------------------------
starting to eat  <3> 		( Using Chopsticks: #2 and #3)
starting to eat  <1> 		( Using Chopsticks: #0 and #1)
finishing eating <3> 		( Eating turns consumed: 3. Units of time required: 611 )
finishing eating <1> 		( Eating turns consumed: 3. Units of time required: 818 )
------------------------------------------------------------------------------------------------------
starting to eat  <4> 		( Using Chopsticks: #3 and #4)
starting to eat  <2> 		( Using Chopsticks: #1 and #2)
finishing eating <2> 		( Eating turns consumed: 3. Units of time required: 273 )
finishing eating <4> 		( Eating turns consumed: 3. Units of time required: 967 )
------------------------------------------------------------------------------------------------------
starting to eat  <5> 		( Using Chopsticks: #4 and #0)
finishing eating <5> 		( Eating turns consumed: 3. Units of time required: 547 )
------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------
MISSION ACOMPLISHED
------------------------------------------------------------------------------------------------------
```