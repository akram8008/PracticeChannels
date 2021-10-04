package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Key struct {
	Value int
	Name  string
}

var CH = make(chan Key)
var wg = &sync.WaitGroup{}

func main() {
	A := make([]int, 800)
	fillNumbers(A)
	
	go func () {
		
		wg.Add(4)
		
		go searchRange(wg, A, 0, 200, "Go1")
		go searchRange(wg, A, 200, 400, "Go2")
		go searchRange(wg, A, 400, 600, "Go3")
		go searchRange(wg, A, 600, 800, "Go4")
		
		wg.Wait()
		
		close(CH)
	}()

	for i := range CH {
		fmt.Println(i)
	}
}

func searchRange(wg *sync.WaitGroup, A []int, l int, r int, name string) {
	defer wg.Done()
	
	for i := l; i < r; i++ {
		if A[i]%2 == 1 {
			val := Key{
				Value: A[i],
				Name:  name,
			}
			CH <- val
		}
	}
}

func fillNumbers(A []int) {
	for i := 0; i < 800; i++ {
		A[i] = rand.Int()
	}
}

