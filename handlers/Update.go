package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maisarasherif/Go-Microservices/data"
)

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
