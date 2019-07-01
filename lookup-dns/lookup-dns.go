package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	ips, err := net.LookupIP("google.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}

	for _, ip := range ips {
		fmt.Printf("google.com. IN A %s\n", ip.String())
	}

}
