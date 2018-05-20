package main

import (
    "fmt"
)

func main() {
    scores := []int{1,4,293,4,9}
    for index, element := range scores {
        fmt.Printf("The %d indexed number's value is %d \n", index, element)
    }
}