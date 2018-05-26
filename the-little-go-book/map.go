package main

import (
    "fmt"
)

func main() {
    lookup := make(map[string]int)

    lookup["gocool"] = 9001

    power, exists := lookup["gocool"]

    fmt.Println(power, exists)

    fmt.Println(len(lookup))

    delete(lookup, "gocool")

    fmt.Println(len(lookup))
}
