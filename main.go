package main

import (
	"fmt"
	"math/rand"
)

type Key struct {
	Value int
	Name  string
}

var CH = make(chan Key)

func main() {
	A := make([]int, 800)
	fillNumbers(A)
	go searchRange(A, 0, 200, "Go1")
	go searchRange(A, 200, 400, "Go2")
	go searchRange(A, 400, 600, "Go3")
	go searchRange(A, 600, 800, "Go4")


	for i := range CH {
		fmt.Println(i)
	}
}

func searchRange(A []int, l int, r int, name string) {
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

