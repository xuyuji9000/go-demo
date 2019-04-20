package main

import (
	"fmt"
	"net/http"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
    	log.Fatal("ListenAndServe: ", err)
	}
}