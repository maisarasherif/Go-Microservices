package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maisarasherif/Go-Microservices/data"
)

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product Not Found", http.StatusInternalServerError)
		return
	}
	p.l.Println("Handle DELETE Product")
}
