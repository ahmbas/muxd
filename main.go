package main

import (
	"os"

	"github.com/ahmbas/muxd/muxd"
	"github.com/ahmbas/muxd/protocols/ip/tcp"
	"github.com/ahmbas/muxd/protocols/ip/udp"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input-protocol",
			Value: "",
			Usage: "Input protocol [tcp, udp]",
		},
		cli.StringFlag{
			Name:  "input-host",
			Value: "",
			Usage: "Input host (IP)",
		},
		cli.IntFlag{
			Name:  "input-port",
			Value: 0,
			Usage: "Input port",
		},
		cli.StringFlag{
			Name:  "output-protocol",
			Value: "",
			Usage: "Output protocol [tcp, udp]",
		},
		cli.StringFlag{
			Name:  "output-host",
			Value: "",
			Usage: "Output host (IP)",
		},
		cli.IntFlag{
			Name:  "output-port",
			Value: 0,
			Usage: "Output port",
		},
	}
	app.Action = func(c *cli.Context) error {
		mux := muxd.Mux{udp.Connection{"udp", "127.0.0.1", 4567}, tcp.Connection{"tcp", "127.0.0.1", 12345}}
		mux.Run()
		return nil
	}

	app.Run(os.Args)

}
