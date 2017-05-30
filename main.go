package main

import (
	"fmt"
	"os"

	//"github.com/dictyBase/gh-issue/commands"
	"gh-issue/gh-issue/commands"

	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "gh-issue"
	app.Usage = "Manage order requests"
	app.Commands = []cli.Command{
		//fill all this in later
		{
			Name:  "test",
			Usage: "i'm testing out the server",
			Action: func(c *cli.Context) error {
				fmt.Println("testing out the server")
				return nil
			},
		},
		{
			Name:   "create",
			Usage:  "Makes issue in Github",
			Action: commands.CreateIssue,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "gh-token, at",
					Usage: "Personal Access Token",
				},
				cli.StringFlag{
					Name:  "title, t",
					Usage: "Issue Title",
				},
				cli.StringFlag{
					Name:  "body, b",
					Usage: "Body text",
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
		{
			Name:   "run",
			Usage:  "runs the server",
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
			},
		},
	}

	app.Run(os.Args)
}
