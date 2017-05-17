package handlers

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dictyBase/gh-issue/models"
	"github.com/google/go-github/github"
	"github.com/manyminds/api2go/jsonapi"
)

type Client struct {
	Github     *github.Client
	Label      string
	Repository string
	Owner      string
	//Logger     *log.Logger
}

/*
type GeneralInfo struct {
	Date    string `json:"date"`
	OrderId string `json:"orderid"`
}

type ShippingInfo struct {
	Name       string `json:"name"`
	University string `json:"university"`
	Lab        string `json:"lab"`
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	Address3   string `json:"address3"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Tracking   string `json:"tracking"`
}

type BillingInfo struct {
	Name       string `json:"name"`
	University string `json:"university"`
	Lab        string `json:"lab"`
	Address    string `json:"address"`
	Address2   string `json:"address2"`
	Address3   string `json:"address3"`
	Country    string `json:"country"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Payment    string `json:"payment"`
}
type Strains struct {
	ID             string `json:"id"`
	Descriptor     string `json:"descriptor"`
	Name           string `json:"name"`
	SystematicName string `json:"SystematicName"`
	Storage        []Storage
	Citations      []Citations
}

type Storage struct {
	StoredAs string `json:"storedas"`
	Location string `json:"location"`
	NoVials  string `json:"novials"`
	Color    string `json:"color"`
}
type Citations struct {
	Citations string `json:"string"`
}

type Plasmids struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Storage   []Storage
	Citations []Citations
}
*/
/*
//OrderInfo : update later with real order fields
type OrderInfo struct {
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
	Title      string `json:"title"`
	GeneralInfo
	ShippingInfo
	BillingInfo
	Strains  []Strains
	Plasmids []Plasmids
}*/

//Jdecoder : returns struct with relevant order fields
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Jsondecoder(w http.ResponseWriter, r *http.Request) models.Orderinfo {
	var order models.Orderinfo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err) //HANDLE THIS BETTER
	}
	err = jsonapi.Unmarshal(body, &order)
	if err != nil {
		log.Print("unmarshal bad")
		panic(err) //HANDLE THIS BETTER

	}
	fmt.Printf("%+v\n", order)
	return order
}

func (client *Client) GithubHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	data := Jsondecoder(w, r)
	var issue = github.IssueRequest{
		Title: &data.CreatedAt,
		Body:  &data.Comments}

	client.Github.Issues.Create(ctx, client.Owner, client.Repository, &issue)
}

//IssueParser : converts orderinfo struct to github issue request
//need to implement error handling
/*
func (order *OrderInfo) IssueParser(o OrderInfo) (github.IssueRequest, error) {
	var issue github.IssueRequest
	issue.Title = &o.Title
	//set the labels
	if o.Plasmids != nil && o.Strains != nil {
		issue.Labels = &[]string{"Strain Order", "Plasmid Order"}
	} else if o.Plasmids == nil && o.Strains != nil {
		issue.Labels = &[]string{"Strain Order"}
	} else if o.Plasmids != nil && o.Strains == nil {
		issue.Labels = &[]string{"Plasmid Order"}
	} else {
		return issue, fmt.Errorf("no order in order")
	}
	//generate markdown and put it in the body
	//maybe try to use text/template package
	/*"**Date:** {{.GeneralInfo.Date}}
	  **Order ID:** {{GeneralInfo.OrderId}}"
	  #Shipping and Billing Info#
	| Shipping                  | Billing/Payer            |
	|---------------------------|--------------------------|
	| o.ShippingInfo.Name       | o.BillingInfo.Name       |
	| o.ShippingInfo.University | o.BillingInfo.University |
	| o.ShippingInfo.Lab        | o.BillingInfo.Lab        |
	| o.ShippingInfo.Address    | o.BillingInfo.Address    |
	| o.ShippingInfo.Address2   | o.BillingInfo.Address2   |
	| o.ShippingInfo.Address3   | o.BillingInfo.Address3   |
	| o.ShippingInfo.Country    | o.BillingInfo.Country    |
	| o.ShippingInfo.Phone      | o.BillingInfo.Phone      |
	| o.ShippingInfo.Email      | o.BillingInfo.Email      |
	| o.ShippingInfo.Tracking   | o.BillingInfo.Payment    |


	return issue, nil
}*/
