package cmd

import (
	// "fmt"
	"github.com/urfave/cli"

	"github.com/go-facegit/facegit-rpc/internal/app"
	"github.com/go-facegit/facegit-rpc/internal/app/router"
	"github.com/go-facegit/facegit-rpc/internal/conf"
)

var Run = cli.Command{
	Name:        "run",
	Usage:       "This command starts all git rpc service",
	Description: `Start Git RPC Service`,
	Action:      RpcRun,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func RpcRun(c *cli.Context) error {
	err := router.Init(c.String("config"))
	if err != nil {
		return err
	}

	app.Start(conf.Rpc.HttpPort)
	return nil
}

func RpcRunDebug() error {
	err := router.Init("")
	if err != nil {
		return err
	}

	app.Start(conf.Rpc.HttpPort)
	return nil
}
