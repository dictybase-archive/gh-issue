package commands

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/dictyBase/gh-issue/middlewares"
	"github.com/dictyBase/gh-issue/resources"
	"github.com/dictyBase/go-middlewares/middlewares/chain"
	"gopkg.in/urfave/cli.v1"
)

func RunServer(c *cli.Context) error {
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
}
