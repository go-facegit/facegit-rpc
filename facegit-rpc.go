package main

import (
	// "log"
	"os"

	"github.com/go-facegit/facegit-rpc/internal/cmd"
	"github.com/go-facegit/facegit-rpc/internal/conf"
	"github.com/go-facegit/facegit-rpc/internal/log"
	"github.com/urfave/cli"
)

func init() {
	conf.App.Version = "0.0.1"
}

func main() {

	app := cli.NewApp()
	app.Name = "FaceGit RPC"
	app.Usage = "A Git RPC service"
	app.Version = conf.App.Version
	app.Commands = []cli.Command{
		cmd.Run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Infof("Failed to start application: %v", err)
	}

	cmd.RpcRunDebug()
}
