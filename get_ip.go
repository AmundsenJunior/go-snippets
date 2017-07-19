package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	hostname := os.Args[1]

	address, err := net.LookupIP(hostname)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("IP address: ", address)
	}
}

// References:
// http://ispycode.com/GO/Network/Get-ipaddress-from-hostname
