package c4

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	str := ""
	for idx, val := range i {
		s := strconv.Itoa(int(val))

		if idx+1 == len(i) {
			str += s
		} else {
			str += s + ","
		}
	}
	return fmt.Sprintf("\"%v\"", str)
}

// TODO: Add a "String() string" method to IPAddr.

func Str() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
