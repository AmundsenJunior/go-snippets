// $ go run get_ip.go jenkins.coredev.cloud
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	hostname := os.Args[1]

	address := getIP(hostname)

	fmt.Println("IP address: ", address)
}

func getIP(hostname string) []net.IP {
	address, err := net.LookupIP(hostname)
	if err != nil {
		panic(err)
	} else {
		return address
	}
}

// References:
// http://ispycode.com/GO/Network/Get-ipaddress-from-hostname
