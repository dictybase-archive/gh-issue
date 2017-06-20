package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//"github.com/dictyBase/gh-issue/middlewares"
	//"github.com/dictyBase/gh-issue/resources"
	//"github.com/dictyBase/gh-issue/routes"
	"gh-issue/gh-issue/auth"
	"gh-issue/gh-issue/middlewares"
	"gh-issue/gh-issue/resources"
	"gh-issue/gh-issue/routes"

	"github.com/dictyBase/go-middlewares/middlewares/chain"
	"github.com/dictybase/go-middlewares/middlewares/logrus"

	"gopkg.in/urfave/cli.v1"
)

//RunServer starts server and sets up default route
func RunServer(c *cli.Context) error {
	var logMw *logrus.Logger
	if c.IsSet("log") {
		w, err := os.Create(c.String("log"))
		if err != nil {
			log.Fatalf("cannot open log file %q\n", err)
		}
		defer w.Close()
		logMw = logrus.NewFileLogger(w)
	} else {
		logMw = logrus.NewLogger()
	}
	router := routes.NewRouter()

	ghInfo := &handlers.Client{
		Repository: c.String("repository"),
		Owner:      c.String("owner"),
		GhClient:   auth.GithubAuth(c.String("gh-token")),
	}

	baseChain := chain.NewChain(logMw.MiddlewareFn, validate.JsonValidator).ThenFunc(ghInfo.OrderHandler)
	router.Post("/dicty/order", baseChain)
	log.Printf("Starting web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), router.Router))
	return nil
}
