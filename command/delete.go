package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/api"
)

// NewDeleteCommand prepares the "delete" command
func NewDeleteCommand() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "delete a key",
		Action: func(c *cli.Context) {
			rawhandle(c, deleteCommandFunc)
		},
	}
}

func deleteCommandFunc(c *cli.Context, client *api.Client) {
	if len(c.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Key required")
		os.Exit(1)
	}

	key := c.Args()[0]
	kv := client.KV()

	_, err := kv.Delete(key, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
