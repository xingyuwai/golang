package main

import (
	"fmt"
)

type IPAddr [4]byte

// print ip in .... style.
func print_ip() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%-10v: \"%3v.%3v.%3v.%3v\"\n", name, ip[0], ip[1], ip[2], ip[3])
	}
}
