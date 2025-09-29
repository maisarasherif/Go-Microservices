package handlers

import (
	"net/http"

	"github.com/maisarasherif/Go-Microservices/data"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	pr := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&pr)
	// test with $ curl localhost:9090 -v -d '{"name": "Big Baik", "description": "a long, crispy chicken fillet in long bun with pickles, lettuce, and garlic sauce", "price": 14.5, "sku": "s-b-b"}'
}
