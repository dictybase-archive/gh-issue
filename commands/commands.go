package commands

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	//"github.com/dictyBase/gh-issue/middlewares"
	//"github.com/dictyBase/gh-issue/resources"
	//"github.com/dictyBase/gh-issue/routes"
	"gh-issue/gh-issue/middlewares"
	"gh-issue/gh-issue/resources"
	"gh-issue/gh-issue/routes"

	"github.com/dictyBase/go-middlewares/middlewares/chain"
	"github.com/dictybase/go-middlewares/middlewares/logrus"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"gopkg.in/urfave/cli.v1"
)

//CreateIssue creates githubclient and posts an issue to specified
//repository/owner (can probably be deleted)
func CreateIssue(c *cli.Context) error {
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
		GhClient:   handlers.GithubAuth(c.String("gh-token")),
	}

	baseChain := chain.NewChain(logMw.MiddlewareFn, validate.JsonValidator).ThenFunc(ghInfo.OrderHandler)
	router.Post("/dicty/order", baseChain)
	log.Printf("Starting web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), router.Router))
	return nil
}
