package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	h.l.Println("Customer Registered")

	d, err := io.ReadAll(r.Body)
	// ERROR HANDLING
	if err != nil {
		http.Error(w, "An Error Encountered", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Welcome %s \n", d)
	// test with $ curl -d "Maysara" localhost:9090
}
