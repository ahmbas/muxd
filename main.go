package main

import (
	"os"

	"github.com/ahmbas/muxd/mux"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {

	var inputProtocol string
	var inputHost string
	var inputPort int
	var inputChannel string
	var outputProtocol string
	var outputHost string
	var outputPort int
	var outputChannel string
	var debugLevel bool

	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Usage = "Network multiplexer"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "input-protocol",
			Value:       "",
			Usage:       "Input protocol [tcp, udp, redis]",
			Destination: &inputProtocol,
		},
		cli.StringFlag{
			Name:        "input-host",
			Value:       "",
			Usage:       "Input host (IP)",
			Destination: &inputHost,
		},
		cli.IntFlag{
			Name:        "input-port",
			Value:       0,
			Usage:       "Input port",
			Destination: &inputPort,
		},
		cli.StringFlag{
			Name:        "input-channel",
			Value:       "",
			Usage:       "Input channel [redis]",
			Destination: &inputChannel,
		},
		cli.StringFlag{
			Name:        "output-protocol",
			Value:       "",
			Usage:       "Output protocol [tcp, udp, redis]",
			Destination: &outputProtocol,
		},
		cli.StringFlag{
			Name:        "output-host",
			Value:       "",
			Usage:       "Output host (IP)",
			Destination: &outputHost,
		},
		cli.IntFlag{
			Name:        "output-port",
			Value:       0,
			Usage:       "Output port",
			Destination: &outputPort,
		},
		cli.StringFlag{
			Name:        "output-channel",
			Value:       "",
			Usage:       "Output channel [redis]",
			Destination: &outputChannel,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Set debug log level",
			Destination: &debugLevel,
		},
	}
	app.Action = func(c *cli.Context) error {

		if debugLevel {
			log.SetLevel(log.DebugLevel)
		}

		inputOpts := mux.Opts{
			Protocol: inputProtocol,
			Host:     inputHost,
			Port:     inputPort,
			Channel:  inputChannel,
		}

		outputOpts := mux.Opts{
			Protocol: outputProtocol,
			Host:     outputHost,
			Port:     outputPort,
			Channel:  outputChannel,
		}

		mux := mux.Mux{
			Input:  mux.GetConnection(inputOpts),
			Output: mux.GetConnection(outputOpts),
		}
		mux.Run()
		return nil
	}

	app.Run(os.Args)

}
