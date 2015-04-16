package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/api"
)

// NewPutCommand prepares the "put" command
func NewPutCommand() cli.Command {
	return cli.Command{
		Name:  "put",
		Usage: "set the value of a key",
		Action: func(c *cli.Context) {
			rawhandle(c, putCommandFunc)
		},
	}
}

func putCommandFunc(c *cli.Context, client *api.Client) {
	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}
	key := c.Args()[0]

	if len(c.Args()) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Value required")
		os.Exit(1)
	}
	value := c.Args()[1]

	kv := client.KV()
	p := &api.KVPair{Key: key, Value: []byte(value)}

	_, err := kv.Put(p, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
