package main

import (
	"container/ring"
	"fmt"
)

func main() {
	RingFunc()
}

func RingFunc() {
	r := ring.New(10) 

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
}
