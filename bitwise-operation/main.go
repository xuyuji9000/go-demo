package main

import (
    "fmt"
    // "math"
)

func main() {
    var a uint8   = 60	/* 60 = 0011 1100 */

    // var b uint = 13	/* 13 = 0000 1101 */

    // var c = a | b /* 61 = 0011 1101 */

    fmt.Println(^a)

}
