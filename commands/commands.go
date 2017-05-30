package commands

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gh-issue/gh-issue/resources"

	"github.com/dictyBase/go-middlewares/middlewares/chain"
	"github.com/dictybase/go-middlewares/middlewares/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"gopkg.in/urfave/cli.v1"
)

func TestFunction(c *cli.Context) error {
	fmt.Println("testing out cli")
	return nil
}

func CreateIssue(c *cli.Context) error {
	//tok, err := ioutil.ReadFile(c.String("gh-token"))
	//if err != nil {
	//return cli.NewExitError(err.Error(), 2)
	//}
	//fmt.Println("token accepted")
	tok := c.String("gh-token")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(tok)},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	title := c.String("title")
	body, err := ioutil.ReadFile("jsontest.txt")
	if err != nil {
		log.Fatalf("cannot read file %q\n", err)
	}
	bodystring := string(body)
	log.Printf(bodystring)
	owner := c.String("owner")
	repository := c.String("repository")
	ctx := context.Background()
	var issue = github.IssueRequest{
		Title: &title,
		Body:  &bodystring}

	client.Issues.Create(ctx, owner, repository, &issue)
	return nil
}

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
	mux := http.NewServeMux()

	ghInfo := &handlers.Client{
		Token:      c.String("gh-token"),
		Repository: c.String("repository"),
		Owner:      c.String("owner"),
	}

	baseChain := chain.NewChain(logMw.MiddlewareFn).ThenFunc(handlers.Placeholder)
	//Chain := apollo.New(apollo.Wrap(logMw.LoggerMiddleware)).With(context.Background()).ThenFunc(handlers.Placeholder)
	mux.Handle("/dicty/order", baseChain)
	log.Printf("Starting web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), mux))
	return nil
}
