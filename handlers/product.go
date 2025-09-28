package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/maisarasherif/Go-Microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// GET method
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// POST method
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// PUT (update) method
	if r.Method == http.MethodPut {

		rx := regexp.MustCompile(`/([0-9]+)`)
		g := rx.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to number", idString)
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProducts(id, w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")

	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Error with GET", http.StatusInternalServerError)
	}
	//test with $ curl localhost:9090 | jq
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	pr := &data.Product{}
	err := pr.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Error with POST", http.StatusBadRequest)
	}

	data.AddProduct(pr)
	// test with $ $ curl localhost:9090 -v -d '{"name": "Big Baik", "description": "a long, crispy chicken fillet in long bun with pickles, lettuce, and garlic sauce", "price": 14.5}'
}

func (p Products) updateProducts(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	pr := &data.Product{}
	err := pr.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Error with PUT", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, pr)

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
