package main

import (
    "fmt"
)

func main() {
    scores := make([]int, 0, 10)
    fmt.Printf("length of socres: %d \n", len(scores))
    scores = scores[0:8]
    scores[7] = 1024
    fmt.Println(scores)
}