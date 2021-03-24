package cmds

import (
	"github.com/urfave/cli"
)

func NewCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "comment"
	app.Usage = "the command of comment"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Value: "./config.json",
			Usage: "the config file path",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "start a comment server",
			Action: Serve,
		},
	}

	return app
}
