package main

import (
	"fmt"
	"sync"
)

const numPhilos = 5

type Philo struct {
	leftCS, rightCS *sync.Mutex
	num, eatCnt int
}

func (p *Philo) initPhilo(num int, m1 *sync.Mutex, m2 *sync.Mutex) {
	p.leftCS = m1
	p.rightCS = m2
	p.num = num
	p.eatCnt = 0
	fmt.Println("Initialized Philosopher", p.num)
}

func (p *Philo) finished() bool {
	return p.eatCnt > 3
}

func (p *Philo) Eat () {
	if p.finished() {
		return
	}

	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Printf("Philosopher %d is starting to eat.\n", p.num)
	p.eatCnt++
	fmt.Printf("Philosopher %d is finished eating.\n", p.num)
	p.leftCS.Unlock()
	p.rightCS.Unlock()
}

func selectPhilo(philos []Philo, dinerChan chan *Philo) {
	// "Host" will keep adding philosophers to the dinerChan
	// since we know 2 adjacent philos can't eat at same time, start staggered
	a := 0
	b := 2
	for {
		if !philos[a].finished() {
			dinerChan <- &philos[a]
		}
		if !philos[b].finished() {
			dinerChan <- &philos[b]
		}
		a = (a + 1) % numPhilos
		b = (b + 1) % numPhilos
	}
}

func feedPhilo(dinerChan chan *Philo) {
	for {
		p := <-dinerChan
		p.Eat()
	}
}

func allDone(diners []Philo) bool {
	for _, p := range diners {
		if !p.finished() {
			return false
		}
	}
	return true
}

func main() {
	var chopsticks [numPhilos]sync.Mutex
	diners := make([]Philo, numPhilos)

	for i := 0; i < numPhilos; i++ {
		diners[i].initPhilo(i, &chopsticks[i], &chopsticks[(i + 1) % 5])
	}

	// Host allows 2 Philosophers to eat concurrently - start up 2 go routines
	dinerChan := make(chan *Philo, 2)
	for i := 0; i < 2; i++ {
		go feedPhilo(dinerChan)
	}

	// Start selecting diners
	go selectPhilo(diners, dinerChan)

	// ... and wait for them to all be done
	// This could be more efficient if we set up a channel, but it works.
	for {
		if allDone(diners) {
			break
		}
	}

	// Print final counts
	for i := 0; i < numPhilos; i++ {
		fmt.Printf("Philosopher %d ate a total of %d times\n", diners[i].num, diners[i].eatCnt - 1)
	}
}