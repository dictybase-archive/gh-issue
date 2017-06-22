package handlers

import (
	"bytes"
	"gh-issue/gh-issue/auth"
	"gh-issue/gh-issue/middlewares"
	"gh-issue/gh-issue/models"

	"github.com/manyminds/api2go/jsonapi"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//IMPORTANT: These tests equire github owner, repository and auth token to be located
//in seperate txt files in resources folder

//TestHandlerSuccess checks that OrderHandler is working.
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

	testHandlerFn := http.HandlerFunc(testClient.OrderHandler)
	w := httptest.NewRecorder()

	//get JSON into buffer to be put in reuest body
	postBody := models.Orderinfo{"1223", "Date1", "Date2", "Fedex", "FedexAccount", "No comment", "Fake payment", "Num 3", "OK status"}
	body, _ := jsonapi.Marshal(postBody)
	b := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/handler-test", b)
	if err != nil {
		t.Fatal(err)
	}
	validate.JSONValidator(testHandlerFn).ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//TestHandlerFailure does not call JSONValidator so OrderHandler is unable to get JSON info from request context
func TestHandlerFailure(t *testing.T) {

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

	testHandlerFn := http.HandlerFunc(testClient.OrderHandler)
	w := httptest.NewRecorder()

	//get JSON into buffer to be put in reuest body
	postBody := models.Orderinfo{"1223", "Date1", "Date2", "Fedex", "FedexAccount", "No comment", "Fake payment", "Num 3", "OK status"}
	body, _ := jsonapi.Marshal(postBody)
	b := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/handler-test", b)
	if err != nil {
		t.Fatal(err)
	}
	testHandlerFn.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
