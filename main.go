package main

import (
	"context"
	"fmt"

	"log"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/woshidama323/LearningGolang/filecoin"
	"github.com/woshidama323/LearningGolang/patterns/options"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "practicegolang",
		Usage:                "practice golang",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			filecoinCmd,
			optionsCmd,
			RawMethodCmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

var optionsCmd = &cli.Command{
	Name:  "optionspattern",
	Usage: "optionspattern test",
	Action: func(c *cli.Context) error {
		helloperson := options.NewPerson("harry", options.Country("ChinaHello"))
		fmt.Println("helloperson is:", helloperson)
		return nil
	},
}

var filecoinCmd = &cli.Command{
	Name:  "filecoin",
	Usage: "filecoin test",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "urlstr",
			Usage: "url for update the code",
		},

		&cli.StringFlag{
			Name:  "minerid",
			Usage: "target miner for get addresses",
		},
	},

	Action: func(c *cli.Context) error {

		geturl := "ws://127.0.0.1:1234/rpc/v0" //
		if c.IsSet("urlstr") {
			geturl = c.String("urlstr")
		}

		minerID := "f0143858" //
		if c.IsSet("minerid") {
			minerID = c.String("minerid")
		}

		api, err := filecoin.NewMiner(geturl)
		if err != nil {
			fmt.Println("failed to new miner,err:", err)
			return err
		}

		defer api.ChainProvider.Closer()
		minerAddress, err := address.NewFromString(minerID)
		if err != nil {
			fmt.Println("failed to get miner addressL:", err)
			return err
		}
		api.ChainProvider.Api.StateMinerInfo(context.Background(), minerAddress, types.EmptyTSK)
		return nil
	},
}

var RawMethodCmd = &cli.Command{
	Name:  "filecoinraw",
	Usage: "filecoinraw test",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "urlstr",
			Usage: "url for update the code",
		},

		&cli.StringFlag{
			Name:  "minerid",
			Usage: "target miner for get addresses",
		},
	},

	Action: func(c *cli.Context) error {
		geturl := "http://192.168.1.102:1234:1234/rpc/v0" //
		if c.IsSet("urlstr") {
			geturl = c.String("urlstr")
		}

		minerID := "f0143858" //
		if c.IsSet("minerid") {
			minerID = c.String("minerid")
		}

		out, _ := filecoin.GetMinerInfo(geturl, minerID)
		test := out.(*miner.MinerInfo)
		fmt.Println("test...", test)
		return nil
	},
}

// func main() {

// 	fmt.Println("test the code")
// 	ips, _ := LocalIPv4s()
// 	fmt.Printf("ips:%v", ips)

// 	//这里就是使用了option的模式来设置一些struct 一些默认值设置在里面
// 	helloperson := options.NewPerson("harry", options.Country("ChinaHello"))
// 	fmt.Println("helloperson is:", helloperson)
// }

// func LocalIPv4s() ([]string, error) {
// 	var ips []string
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return ips, err
// 	}

// 	for _, a := range addrs {
// 		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
// 			ips = append(ips, ipnet.IP.String())
// 		}
// 	}

// 	return ips, nil
// }
