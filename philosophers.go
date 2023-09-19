package main

import (
	"fmt"
	"sync"
)

// anonymous struct field
type chopstick struct {
	mu     sync.Mutex
	number int
}

// nested struct
type philosopher struct {
	number          int
	eatcount        int
	leftcs, rightcs *chopstick
	phost           *host
}

type host struct {
	hostCount []int
}

var (
	wg      sync.WaitGroup
	ourHost = host{hostCount: make([]int, 0, 5)}
)

func (h *host) dinnerService(requestChan <-chan *philosopher, doneChan <-chan *philosopher, closeChan <-chan bool, responseChan chan<- bool) {

	hostAudit := []int{}

	for {
		select {
		case p := <-doneChan:
			tempnumber := -1
			for number, e := range h.hostCount {
				if e == p.number {
					tempnumber = number
					break
				}
			}
			if tempnumber == -1 {
				panic("finished philosopher not found in record.")
			}
			h.hostCount = append(h.hostCount[:tempnumber], h.hostCount[tempnumber+1:]...)
		case p := <-requestChan:
			if len(h.hostCount) >= 5 {
				if !intSlicesEqual(h.hostCount, hostAudit) {
					fmt.Println("hostque is full. Que status:", h.hostCount)
				}
				responseChan <- false
				hostAudit = append([]int{}, h.hostCount...)
				continue
			}
			p.rightcs.mu.Lock()
			p.leftcs.mu.Lock()

			responseChan <- true
			h.hostCount = append(h.hostCount, p.number)

		case <-closeChan:
			fmt.Println("hosting aborted")
			return
		}
	}
}

func (p *philosopher) eat(requestChan chan<- *philosopher, doneChan chan<- *philosopher, responseChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for p.eatcount < 3 {
		requestChan <- p
		response := <-responseChan
		if !response {
			continue
		}
		fmt.Println("Philosopher", p.number, "has started to eat.")
		p.eatcount++
		fmt.Println("Philosopher", p.number, "has finished eating.")

		p.leftcs.mu.Unlock()
		p.rightcs.mu.Unlock()
		doneChan <- p
	}
}

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	requestChan := make(chan *philosopher)
	responseChan := make(chan bool)
	doneChan := make(chan *philosopher)
	closeChan := make(chan bool)

	csticks := make([]*chopstick, 5)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		csticks[i] = &chopstick{number: i}
		wg.Done()
	}
	wg.Wait()
	philos := make([]*philosopher, 5)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		philos[i] = &philosopher{number: i + 1, eatcount: 0, leftcs: csticks[i], rightcs: csticks[(i+1)%5]}
		wg.Done()
	}

	wg.Wait()
	go ourHost.dinnerService(requestChan, doneChan, closeChan, responseChan)
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(requestChan, doneChan, responseChan, &wg)
	}
	wg.Wait()
	fmt.Println("Service Complete.\n philosopher1 eatcount:", philos[0].eatcount,
		"\n philosopher2 eatcount:", philos[1].eatcount, "\n philosopher3 eatcount:",
		philos[2].eatcount, "\n philosopher4 eatcount:", philos[3].eatcount,
		"\n philosopher5 eatcount:", philos[4].eatcount)
	closeChan <- true
}
