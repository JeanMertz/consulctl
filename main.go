package main

import (
	"os"

	"github.com/JeanMertz/consulctl/command"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "consulctl"
	app.Version = "0.3.0"
	app.Author = "Jean Mertz"
	app.Email = "jean@mertz.fm"
	app.Usage = "A simple command line client for Consul k/v store."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "address, A",
			Value:  "127.0.0.1:8500",
			Usage:  "the remote endpoint for the Consul cluster",
			EnvVar: "CONSUL_KV_ADDRESS",
		},
		cli.IntFlag{
			Name:   "wait, w",
			Usage:  "request timeout value in seconds",
			EnvVar: "CONSUL_KV_WAIT",
		},
	}
	app.Commands = []cli.Command{
		command.NewGetCommand(),
		command.NewPutCommand(),
		command.NewDeleteCommand(),
	}

	app.Run(os.Args)
}
