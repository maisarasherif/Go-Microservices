package handlers

import (
	"log"
	"net/http"

	"github.com/maisarasherif/Go-Microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	//test with $ curl localhost:9090 | jq
}
