package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go/jsonapi"
)

type Orderinfo struct {
	ID                string `json:"-"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Courier           string `json:"courier"`
	CourierAccount    string `json:"courieraccount"`
	Comments          string `json:"comments"`
	Payment           string `json:"payment"`
	PurchaseOrderNumb string `json:"purchase_order_num"`
	Status            string `json:"status"`
}

type UnmarshalIdentifier interface {
	SetID(string) error
}

func (order *Orderinfo) SetID(id string) error {
	order.ID = id
	return nil
}

//MOVE EVERYTHING BELOW TO HANDLERS FILE WHEN READY
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Jsondecoder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var order Orderinfo
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
}
