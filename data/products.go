package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"` //using struct tags (how the fields should appear in json output)
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` //remove it from the json output. (internal use only)
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Karak Tea",
		Description: "Tea with steamed milk, cardamom, and sugar.",
		Price:       5,
		SKU:         "KT5",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Whopper Burger meal",
		Description: "beef burger patty bun with cheese, tomatoes, lettuce, onions, and sauce",
		Price:       24,
		SKU:         "WBM24",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
