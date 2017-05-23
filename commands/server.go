package commands

import (
	"context"
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/dictyBase/gh-issue/resources"

	"github.com/cyclopsci/apollo"
	"github.com/dictyBase/gh-issue/middlewares"
	"github.com/urfave/cli"
)

func RunServer(c *cli.Context) {
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

	Chain := apollo.New(apollo.Wrap(logMw.LoggerMiddleware)).With(context.Background()).ThenFunc(handlers.Placeholder)
	mux.Handle("/dicty/order", Chain)
	log.Printf("Starting web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), mux))
}
