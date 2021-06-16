package main

import (
	"fmt"

	"net"
)

func main() {
	fmt.Println("test the code")

	// la := netlink.NewLinkAttrs()
	// la.Name = "foo"
	// mybridge := &netlink.Bridge{LinkAttrs: la}
	// err := netlink.LinkAdd(mybridge)
	// if err != nil  {
	//     fmt.Printf("could not add %s: %v\n", la.Name, err)
	// }
	// link, err := netlink.LinkByName("lo0")
	// if err != nil {
	// 	fmt.Printf("could not get lo %v", err)
	// }
	// eth1, _ := netlink.AddrList(link, 2)
	// fmt.Printf("eth1 is %v", eth1)

	ips, _ := LocalIPv4s()
	fmt.Printf("ips:%v", ips)
}

func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}
