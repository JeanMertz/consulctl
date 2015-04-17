package command

import (
	"net/http"
	"net/url"
	"time"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/api"
)

type handlerFunc func(*cli.Context, *api.Client)

func rawhandle(c *cli.Context, fn handlerFunc) {
	address := c.GlobalString("address")
	timeout := c.GlobalInt("wait")
	consulConfig := api.Config{Address: address, Scheme: "http"}
	client, _ := api.NewClient(&consulConfig)

	if timeout != 0 {
		handleTimeout(timeout, consulConfig)
	}

	fn(c, client)
}

func handleTimeout(timeout int, conf api.Config) {
	req := url.URL{Scheme: conf.Scheme, Host: conf.Address}

	for i := 1; i <= timeout; i++ {
		_, err := http.Get(req.String())
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
