// https://www.coursera.org/learn/golang-concurrency/peer/MAemV/module-4-activity/submit
package main

import (
	"sync"
	"fmt"
	"time"
)

type Chop struct{ sync.Mutex }

type Diner struct{
	dinerNum int
	leftC *Chop
	rightC *Chop
}

// eat function let's a diner to eat if allowedDiner channel has true input.
// No more than two diners can eat at a time.
func ( d Diner ) eat( allowedDiner chan bool ) {
	isAllowed := <- allowedDiner
	if isAllowed {
		for i:=0; i<3; i++ {
			d.leftC.Lock()
			d.rightC.Lock()
			fmt.Println("starting to eat ", d.dinerNum)
			fmt.Println("finishing eating ", d.dinerNum)
			d.leftC.Unlock()
			d.rightC.Unlock()
		}
	}
}

// allow function, sets the allowed channel to be true
func ( d Diner ) allow( allowed chan bool ) {
	allowed <- true
}

func main() {
	CSticks := make([]*Chop, 5 )
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chop)
	}
	
	Diners := make([]*Diner, 5)
	for i:=0; i<5; i++ {
		Diners[i] = &Diner{ i+1, CSticks[i], CSticks[(i+1)%5] }
	}

	// Create a channel to take 2 inputs without blocking
	allowedDiner := make( chan bool, 2 )
	
	for i:=0; i<5; i++ {
		go Diners[i].allow( allowedDiner )
		go Diners[i].eat( allowedDiner)
	}
	time.Sleep( 5 * time.Second )
}
