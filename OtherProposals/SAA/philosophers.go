package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	c1, c2 *ChopStick
	name   string
}

func (p Philosopher) eat(m *sync.Mutex, add, remove chan string, wg *sync.WaitGroup, meals int) {
	for i := 0; i < meals; i++ {
		add <- p.name
		p.c1.Lock()
		p.c2.Lock()
		fmt.Printf("starting to eat %s\n", p.name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("finishing eating %s\n", p.name)
		remove <- p.name
		p.c1.Unlock()
		p.c2.Unlock()
		m.Lock()
	}
	wg.Done()
}

func finishEating(abort chan string, wg *sync.WaitGroup) {
	wg.Wait()
	abort <- "done"
}

func host(persons, limit, meals int) {
	var (
		input     string
		count     = 0
		waitList  = make([]string, 0)
		mutexList = make([]sync.Mutex, persons)
		people    = make([]Philosopher, 0)
		sticks    = make([]ChopStick, persons)
		add       = make(chan string)
		remove    = make(chan string)
		abort     = make(chan string)
		wg        sync.WaitGroup
	)
	wg.Add(persons)
	go finishEating(abort, &wg)
	for i := 0; i < persons; i++ {
		people = append(people, Philosopher{&sticks[i], &sticks[(i+1)%persons], strconv.Itoa(i)})
		mutexList = append(mutexList, sync.Mutex{})
		mutexList[i].Lock()
		go people[i].eat(&mutexList[i], add, remove, &wg, meals)
	}
	for {
		select {
		case input = <-add:
			if count < limit {
				count += 1
				num, _ := strconv.Atoi(input)
				mutexList[num].Unlock()
			} else {
				waitList = append(waitList, input)
			}
		case input = <-remove:
			if len(waitList) != 0 {
				num, _ := strconv.Atoi(waitList[0])
				mutexList[num].Unlock()
				waitList = waitList[1:]
			} else {
				count -= 1
			}
		case <-abort:
			return
		default:
		}
	}
}

func main() {
	host(5, 2, 3)
}
