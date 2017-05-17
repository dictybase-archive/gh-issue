package main

import (
	"gh-issue/commands"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	commands.Hello()
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "gh-issue"
	app.Usage = "Manage order requests"
	app.Commands = []cli.Command{
		//fill all this in later
		{
			Name:   "test",
			Usage:  "i'm testing out the server",
			Action: commands.TestFunction,
		},
		{
			Name:   "run",
			Usage:  "starts the webhook server for gmail push notifications",
			Action: commands.RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "subscription, s",
					Usage: "Name of the subscription",
				},
				cli.StringFlag{
					Name:  "project, p",
					Usage: "Name of the project",
				},
				cli.StringFlag{
					Name:   "cache-file, cf",
					Usage:  "location of cached gmail token file, defaults to ~/.credentials/gmail.json",
					EnvVar: "CACHE_TOKEN_FILE",
				},
				cli.StringFlag{
					Name:  "gmail-secret, gs",
					Usage: "gmail client secret json file",
				},
				cli.StringFlag{
					Name:  "gh-token, ght",
					Usage: "github personal access token file, defaults to ~/.credentials/github.json",
				},
				cli.StringFlag{
					Name:  "log,l",
					Usage: "Name of the web request log file(optional), default goes to stderr",
				},
				cli.StringFlag{
					Name:  "app-log,apl",
					Usage: "Name of the application log file(optional), default goes to stderr",
				},
				cli.IntFlag{
					Name:  "port",
					Usage: "port on which the server listen",
					Value: 9998,
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Gmail label which will be filtered for messages",
				},
				cli.StringFlag{
					Name:  "repository, r",
					Usage: "Github repository",
				},
				cli.StringFlag{
					Name:  "owner",
					Usage: "Github repository owner",
				},
				cli.StringFlag{
					Name:  "redis-address",
					Usage: "IP address of redis-server",
					Value: "redis",
				},
				cli.IntFlag{
					Name:  "redis-port",
					Usage: "Port of redis server",
					Value: 6379,
				},
			},
		},
	}

	app.Run(os.Args)
}
