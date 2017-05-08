package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Name = "gh-issue"
	app.Usage = "Manage order requests"
	app.Commands = []cli.Command{
	//fill all this in later
	}
	app.Run(os.Args)
}
