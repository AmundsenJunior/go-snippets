// $ go run get_ip.go jenkins.coredev.cloud
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	hostname := os.Args[1]

	address, err := net.LookupIP(hostname)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("IP address: ", address)
	}
}

// References:
// http://ispycode.com/GO/Network/Get-ipaddress-from-hostname
