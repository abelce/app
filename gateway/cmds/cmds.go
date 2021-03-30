package cmds

import "github.com/urfave/cli"

//NewCliApp 创建命令行程序
func NewCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gateway"
	app.Usage = "the commands of gateway"
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
			Usage:  "start a invoice server",
			Action: serve,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "80",
					Usage: "开启80端口，内网调用端口",
				},
				cli.BoolFlag{
					Name:  "8443",
					Usage: "开启8443端口，外网调用端口",
				},
			},
		},
	}

	return app
}
