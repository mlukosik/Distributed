package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

const n = 5

type processStruct struct {
		processID int //ID procesu
		leaderID int //maxID otrzymane od sąsiadów 
		index int //index w slice ring
		neighbourID int //ID sąsiada
}

func msg(i int, ch chan processStruct, process processStruct, ring []processStruct, chLeader chan processStruct) {
	for i != process.processID {
		if i > process.leaderID {
			process.leaderID = i
			ch <- process
		} else {
			ch <- process
		}
		k := search(i, ring)
		i = ring[k].neighbourID
	}
	chLeader <- process
}

func search(i int, ring []processStruct) int {
	index := 0
	for j := 0; j < n; j++ {
		if ring[j].processID == i {
			index = j
		}
	}
	return index
}

func main() {

	ring := make([]processStruct, n)
	rand.Seed(time.Now().UTC().UnixNano())
	s := rand.Perm(n)
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		ring[i].processID = s[i]
		ring[i].leaderID = s[i]
		ring[i].index = i
	}


	for i := 0; i < len(s)-1; i++ {
		ring[i].neighbourID = ring[i+1].processID
	}
	ring[n-1].neighbourID = ring[0].processID
	fmt.Println(ring)


	ch := make(chan processStruct)
	chLeader := make(chan processStruct)

	for i := 0; i < n; i++ {
		time.Sleep(100* time.Millisecond)
		go msg(ring[i].neighbourID, ch, ring[i], ring, chLeader)
	}

	y := float64(n)
	x := int(math.Pow(y, 2))
	fmt.Println(x)
	for i := 0; i < x; i++ {
		time.Sleep(100* time.Millisecond)
	
		select {
			case msg := <- ch:
				fmt.Printf("Nowy lider %d dla procesu %d\n", msg.leaderID, msg.processID)
			case msg := <- chLeader:
				fmt.Printf("================> Proces: %d Lider: %d\n", msg.processID, msg.leaderID)
				ring[msg.index].leaderID = msg.leaderID
			default:
				fmt.Println("ok")
		}
	}	

	fmt.Println()
	for i := 0; i < len(ring); i++ {
		fmt.Printf("Proces: %d Lider %d\n", ring[i].processID, ring[i].leaderID)
	}
}
