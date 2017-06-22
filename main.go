package main

import (
	"os"

	"github.com/dictyBase/gh-issue/commands"
	//"gh-issue/gh-issue/commands"

	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "gh-issue"
	app.Usage = "Manage order requests"
	app.Commands = []cli.Command{

		{
			Name:   "run",
			Usage:  "Runs the server",
			Action: commands.RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "log,l",
					Usage: "Name of the web request log file(optional), default goes to stderr",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "port on which the server listen",
					Value: 9998,
				},
				cli.StringFlag{
					Name:  "gh-token, at",
					Usage: "Personal Access Token",
				},
				cli.StringFlag{
					Name:  "owner, o",
					Usage: "Repository owner",
				},
				cli.StringFlag{
					Name:  "repository, r",
					Usage: "repository",
				},
			},
		},
	}

	app.Run(os.Args)
}
