package main

import "fmt"

type Saiyan struct {
    Name string
    Power int
}

func main() {
    goku := Saiyan{"Goku", 9000}
    fmt.Printf("%s's power over %d\n", goku.Name, goku.Power)
}