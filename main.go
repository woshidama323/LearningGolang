package main

import (
	"context"
	"fmt"

	"log"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/multiformats/go-multiaddr"
	"github.com/woshidama323/LearningGolang/filecoin"
	"github.com/woshidama323/LearningGolang/markdown"
	"github.com/woshidama323/LearningGolang/patterns/options"
	"github.com/woshidama323/LearningGolang/rpcserver"

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
			RPCServerTestCmd,
			MarkDownPocCmd,
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

		addrlist, err := filecoin.GetMinerInfo(geturl, minerID)
		if err != nil {
			fmt.Println("failed to get miner info,err:", err)
			return err
		}

		for _, addr := range addrlist {

			getaddr, err := filecoin.GetAddressInfo(geturl, addr)
			if err != nil {
				fmt.Println("failed to get the address info,err:", err)
				continue
			}
			fmt.Println("get current addresslist :", getaddr)

		}

		return nil
	},
}

var RPCServerTestCmd = &cli.Command{
	Name:  "runrpc",
	Usage: "hello rpc server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "xxx",
			Usage: "xxx",
		},
	},
	Action: func(c *cli.Context) error {

		fn := &rpcserver.ImplementFullNode{
			Test: "helo",
		}
		wfn, err := rpcserver.FlutterHandler(fn)
		if err != nil {
			return err
		}
		am, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/11111")
		if err != nil {
			return err
		}

		_, err = rpcserver.ServerRPC(wfn, "rpcserver", am)
		if err != nil {
			fmt.Println("failed to start rpc server:", err)
			return err
		}
		// rpcStopper()
		select {}
		// return nil
	},
}

var MarkDownPocCmd = &cli.Command{
	Name:  "markdown",
	Usage: "command test for generating markdown file",
	Action: func(c *cli.Context) error {
		// mk := markdown.NewMarkDownTemplate()
		mk := markdown.NewMarkDownTemplate()
		mk.Indicator = []string{
			"昨日出块奖励",
			"昨日出块奖励",
			"昨日出块奖励",
		}

		mk.MinerList = []string{
			"f02301",
			"f02301",
			"f02301",
		}

		mk.MinerFeeInfo()
		return nil
	},
}

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
