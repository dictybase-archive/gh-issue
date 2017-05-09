package handlers

import (
	"encoding/json"
	"net/http"
)

//OrderInfo : update later with real order fields
type OrderInfo struct {
	Email   string `json:"email"`
	Plasmid string `json:"plasmid"`
	Strain  string `json:"strain"`
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
