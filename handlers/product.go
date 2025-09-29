package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maisarasherif/Go-Microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")

	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Error with GET", http.StatusInternalServerError)
	}
	//test with $ curl localhost:9090 | jq
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	pr := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&pr)
	// test with $ curl localhost:9090 -v -d '{"name": "Big Baik", "description": "a long, crispy chicken fillet in long bun with pickles, lettuce, and garlic sauce", "price": 14.5}'
}

func (p Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT product", id)
	pr := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &pr)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product Not Found", http.StatusInternalServerError)
		return
	}
	// test with $ curl localhost:9090/3 -XPUT -d '{"name": "Spicy Big Baik", "description": "a long, crispy spicy chicken fillet in long bun with jalapeno, lettuce, and garlic sauce", "price": 14.5, "sku": "SBB14.5"}'
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pr := data.Product{}

		err := pr.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = pr.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, pr)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
