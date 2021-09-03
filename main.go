package main

import (
	"fmt"
	"net"

	"github.com/woshidama323/LearningGolang/patterns/options"
)

func main() {
	fmt.Println("test the code")
	ips, _ := LocalIPv4s()
	fmt.Printf("ips:%v", ips)

	//这里就是使用了option的模式来设置一些struct 一些默认值设置在里面
	helloperson := options.NewPerson("harry", options.Country("ChinaHello"))
	fmt.Println("helloperson is:", helloperson)
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
