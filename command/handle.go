package command

import (
	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/api"
)

type handlerFunc func(*cli.Context, *api.Client)

func rawhandle(c *cli.Context, fn handlerFunc) {
	address := c.GlobalString("address")
	consulConfig := api.Config{Address: address, Scheme: "http"}
	client, _ := api.NewClient(&consulConfig)

	fn(c, client)
}
