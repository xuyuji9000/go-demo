package main



import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOARCCH:",runtime.GOARCH)
	fmt.Println("GOOS:",runtime.GOOS)
}