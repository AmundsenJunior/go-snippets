// copied from ispycode.com/GO/Network/Get-ipaddress-from-hostname
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupIP("jenkins.coredev.cloud")
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		fmt.Println("IP address: ", addr)
	}
}
