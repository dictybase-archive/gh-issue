package commands

import (
	"context"
	"fmt"

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
	body := c.String("body")
	owner := c.String("owner")
	repository := c.String("repository")
	ctx := context.Background()
	var issue = github.IssueRequest{
		Title: &title,
		Body:  &body}

	client.Issues.Create(ctx, owner, repository, &issue)
	return nil
}
