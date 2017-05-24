package main

import (
	"fmt"
	"gh-issue/gh-issue/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/dictyBase/gh-issue/commands"
	"github.com/dictyBase/gh-issue/resources"
	"github.com/dictyBase/go-middlewares/middlewares/chain"

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
			Name:  "run",
			Usage: "runs the server",
			Action: func(c *cli.Context) error {
				var logMw *middlewares.Logger
				if c.IsSet("log") {
					w, err := os.Create(c.String("log"))
					if err != nil {
						log.Fatalf("cannot open log file %q\n", err)
					}
					defer w.Close()
					logMw = middlewares.NewFileLogger(w)
				} else {
					logMw = middlewares.NewLogger()
				}
				mux := http.NewServeMux()

				baseChain := chain.NewChain(logMw.LoggerMiddleware).ThenFunc(handlers.Placeholder)
				//Chain := apollo.New(apollo.Wrap(logMw.LoggerMiddleware)).With(context.Background()).ThenFunc(handlers.Placeholder)
				mux.Handle("/dicty/order", baseChain)
				log.Printf("Starting web server on port %d\n", c.Int("port"))
				log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), mux))
				return nil
			},

			//commands.RunServer,
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
