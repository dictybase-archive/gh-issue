package handlers

import (
	"gh-issue/gh-issue/auth"
	"io/ioutil"
	"testing"
)

func TestHandlerSuccess(t *testing.T) {

	//first need to create the client struct
	repo, err := ioutil.ReadFile("repository.txt")
	if err != nil {
		t.Fatal(err)
	}
	repostring := string(repo)
	owner, err := ioutil.ReadFile("owner.txt")
	if err != nil {
		t.Fatal(err)
	}
	authtoken, err := ioutil.ReadFile("auth_token.txt")
	if err != nil {
		t.Fatal(err)
	}
	authtokenstring := string(authtoken)
	ownerstring := string(owner)
	testClient := Client{
		Repository: repostring,
		Owner:      ownerstring,
		GhClient:   auth.GithubAuth(authtokenstring),
	}
}
