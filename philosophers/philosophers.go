package main

// =============================================================================================================================
// Implement the dining philosopher’s problem with the following constraints/modifications.
// =============================================================================================================================
//
//     1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
//     2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
//     3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
//     4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
//     5. The host allows no more than 2 philosophers to eat concurrently.
//     6. Each philosopher is numbered, 1 through 5.
//     7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
//         on a line by itself, where <number> is the number of the philosopher.
//     8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>”
//         on a line by itself, where <number> is the number of the philosopher.
// =============================================================================================================================

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var separator = "------------------------------------------------------------------------------------------------------"

// ------------------------------------------------------------------------
// A Chop Stick structure
// ------------------------------------------------------------------------
// Note:
// Due to the introduction of a regulator agent ("the host") only the
// Chop Stick identifier is included as it will not be necessary
// to implement any mutex in this structure.
// ------------------------------------------------------------------------
type ChopS struct {
	id int
}

// ------------------------------------------------------------------------
// A Host structure (The "regulator agent")
// ------------------------------------------------------------------------
type Host struct {
	activePhilos int          // Number of currently Philosopher still candidates to have a eat turn
	philos       []*Philo     // The list of Philosophers (to gain access to their info)
	signalBack   *chan string // The shared signal channel where notifications will be received
}

func (h Host) giveTurn(wg *sync.WaitGroup, mutex *sync.Mutex) {
	i := 0
	defer wg.Done()

	for {

		turn1 := (i) % 5
		turn2 := (i + 2) % 5

		turn1Consumed := 0
		turn2Consumed := 0

		mutex.Lock()
		turn1Consumed = h.philos[turn1].eatingTurns
		turn2Consumed = h.philos[turn2].eatingTurns
		mutex.Unlock()

		if turn1Consumed < 3 {
			h.philos[turn1].signalin <- "Proceed"
		} else {
			h.activePhilos--
		}

		if turn2Consumed < 3 {
			h.philos[turn2].signalin <- "Proceed"
		} else {
			h.activePhilos--
		}

		if turn1Consumed < 3 {
			<-*h.philos[turn1].signalout
		}

		if turn2Consumed < 3 {
			<-*h.philos[turn2].signalout
		}

		if h.activePhilos <= 0 {
			break
		} else {
			i++
			fmt.Println(separator)
		}

	}
}

// ------------------------------------------------------------------------
// A Philosopher structure
// ------------------------------------------------------------------------
type Philo struct {
	id              int          // Its identificator
	leftCS, rightCS *ChopS       // The left and right hands chop sticks assigned
	signalin        chan string  // The chanel to receive activation signal/permission
	signalout       *chan string // The chanel to send back ending to the HOST
	eatingTime      int          // The time units that this philosopher consumes to eat
	eatingTurns     int          // The eating turns already consumed
}

func (philo *Philo) eat(wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	for {
		<-philo.signalin
		fmt.Printf("starting to eat  <%d> \t\t( Using Chopsticks: #%d and #%d)\n", philo.id, philo.leftCS.id, philo.rightCS.id)
		mutex.Lock()
		philo.eatingTurns++
		mutex.Unlock()
		time.Sleep(time.Duration(philo.eatingTime) * time.Millisecond)
		fmt.Printf("finishing eating <%d> \t\t( Eating turns consumed: %d. Units of time required: %d )\n", philo.id, philo.eatingTurns, philo.eatingTime)
		*philo.signalout <- "done"

		if philo.eatingTurns >= 3 {
			break
		}
	}
}

// ------------------------------------------------------------------------
// MAIN execution entry point
// ------------------------------------------------------------------------
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	// ------------------------------------------------------------------------
	// Chop Sticks
	// ------------------------------------------------------------------------
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = &ChopS{i}
	}

	// ------------------------------------------------------------------------
	// A channel to receive a signal back when a Philosopher ends eating
	// ------------------------------------------------------------------------
	signalBack := make(chan string)

	// ------------------------------------------------------------------------
	// The Philosopher list configuration
	// ------------------------------------------------------------------------
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			i + 1,
			CSticks[i],
			CSticks[(i+1)%5],
			make(chan string),
			&signalBack,
			rand.Intn(1000),
			0,
		}
	}

	// ------------------------------------------------------------------------
	// The Host configuration that intervene give turn permissions
	// ------------------------------------------------------------------------
	host := Host{5, philos, &signalBack}

	// ------------------------------------------------------------------------
	// The MAIN routine
	// ------------------------------------------------------------------------
	fmt.Println(separator)
	fmt.Println("MISSION BEGINS")
	fmt.Println(separator)

	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup

	wg.Add(6)
	go host.giveTurn(&wg, mutex)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, mutex)
	}
	wg.Wait()

	fmt.Println(separator)
	fmt.Println("MISSION ACOMPLISHED")
	fmt.Println(separator)

}
