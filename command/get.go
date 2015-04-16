package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/api"
)

// NewGetCommand prepares the "get" command
func NewGetCommand() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "retrieve the value of a key",
		Action: func(c *cli.Context) {
			rawhandle(c, getCommandFunc)
		},
	}
}

func getCommandFunc(c *cli.Context, client *api.Client) {
	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}

	key := c.Args()[0]
	kv := client.KV()

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		panic(err)
	}

	if pair != nil {
		fmt.Print(string(pair.Value))
	}
}
