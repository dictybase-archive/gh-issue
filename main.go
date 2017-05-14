package main

import (
	"os"

	"gh-issue/commands"

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
			Name:   "test",
			Usage:  "i'm testing out the server",
			Action: commands.TestFunction,
		},
	}

	app.Run(os.Args)
}
