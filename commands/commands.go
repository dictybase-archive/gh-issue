package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"gopkg.in/urfave/cli.v1"
)

func Hello() {
	fmt.Println("hello")
}
func TestFunction(c *cli.Context) {
	fmt.Println("testing out the server")
}

func CreateIssue(c *cli.Context) {
	tok, err := ioutil.ReadFile(c.String("gh-token"))
	if err != nil {
		fmt.Errorf("error cannot open token")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(tok)},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	title := c.String("title")
	body := c.String("body")
	owner := c.String("owner")
	repository := c.String("repository")
	var issue = github.IssueRequest{
		Title: &title,
		Body:  &body}

	client.Issues.Create(context.Context, owner, repository, &issue)
}
