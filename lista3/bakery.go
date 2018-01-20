package main

import (
	"fmt"
	"time"
)

const n = 10
var res int = -1


func max(a[] int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
		}
	}
	return m
}


func lock(k int, entering []bool, number []int) {
	entering[k] = true
	number[k] = 1 + max(number)
	entering[k] = false

	for i := 0; i < n; i++ {
		for entering[i] { 
			time.Sleep(10* time.Millisecond)		
		}
		for number[i] != 0 && (number[i] < number[k] || (number[i] == number[k] && i < k)) {
			time.Sleep(10* time.Millisecond)
		}
	}
}


func unlock(k int, number []int) {
	number[k] = 0
}


func use(k int, entering []bool, number []int) {
	for {
		select  {
			default :
				lock(k, entering, number)

				time.Sleep(1 * time.Millisecond)
				fmt.Printf("Proces %d wchodzi do sekcji krytycznej\n", k)
				/*
				if res != -1 {
					fmt.Printf("Zasób w użyciu, jednak proces %d przejmuje dostęp\n", k)
				} */
				res = k
				fmt.Printf("Proces %d wychodzi z sekcji krytycznej\n", k)
				res = -1

				unlock(k, number)
		}
	}
}

func main() {
	entering := make([]bool, n+1)
	number := make([]int, n+1)

	for i := 0; i < n; i++ {
		entering[i] = false
		number[i] = 0
	}

	for i := 0; i < n; i++ {
		go use(i, entering, number)
	}
	
	for {
		lock(n, entering, number)
		
		unlock(n, number)
	}
}
