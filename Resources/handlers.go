package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

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
	ID   string `json:"id"`
	Name string `json:"name"`
}

//OrderInfo : update later with real order fields
type OrderInfo struct {
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Plasmid    string `json:"plasmid"`
	Strain     string `json:"strain"`
}

//JDecoder : returns struct with relevant order fields
func (order *OrderInfo) JDecoder(w http.ResponseWriter, r *http.Request) OrderInfo {
	var o OrderInfo
	if r.Body == nil {
		http.Error(w, "please send a request body", 400) //what number for the error?
	}
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, "decoding error!", 400) //what number
	}
	return o
}

//IssueParser : converts orderinfo struct to github issue request
//need to implement error handling
func (order *OrderInfo) IssueParser(o OrderInfo) (github.IssueRequest, error) {
	var issue github.IssueRequest
	if o.Plasmid != "" && o.Strain != "" {
		issue.Title = &o.Title
		issue.Body = &o.Body
		issue.Labels = &[]string{"Strain Order", "Plasmid Order"}
	} else if o.Plasmid == "" && o.Strain != "" {
		issue.Title = &o.Title
		issue.Body = &o.Body
		issue.Labels = &[]string{"Strain Order"}
	} else if o.Plasmid != "" && o.Strain == "" {
		issue.Title = &o.Title
		issue.Body = &o.Body
		issue.Labels = &[]string{"Plasmid Order"}
	} else {
		return issue, fmt.Errorf("no order in order")
	}

	return issue, nil
}
