package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	elliot := &Person{
		Name: "Alice",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)

	if nil != err {
		log.Fatal("marshaling error:", err)
	}
	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
}
