package models

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

/*
type UnmarshalIdentifier interface {
	SetID(string) error
}*/

func (order *Orderinfo) SetID(id string) error {
	order.ID = id
	return nil
}
