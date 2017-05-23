package main

import (
	"fmt"
	"gh-issue/gh-issue/commands"
	"os"

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

			//commands.TestFunction,
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
	}

	app.Run(os.Args)
}
