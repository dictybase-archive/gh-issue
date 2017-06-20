package auth

import (
	"github.com/google/go-github/github"

	"golang.org/x/oauth2"
)

//GithubAuth takes the auth token as input and returns authorized github client
func GithubAuth(token string) *github.Client {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc) //Supposed to implement error handling but github.NewClient only returns 1 item..?

	return client
}
