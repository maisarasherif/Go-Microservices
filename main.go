package main

import (
	"fmt"
	"io"
	"log" //http package for creating web servers
	"net/http"
)

func main() {
	//HandleFunc registers a function to a path in the DefaultServeMux (an http handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Customer Registered")
		d, err := io.ReadAll(r.Body)
		// ERROR HANDLING
		if err != nil {
			http.Error(w, "An Error Encountered", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Welcome %s \n", d)
		// test with $ curl -d "Maysara" localhost:9090

	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("BonApetit")
	})

	// when the path matches "/goodbye" or any other endpoint, it executes the endpoint function.
	// when the path is any thing other than "/goodbye" or any other endpoint even if it doesn't exist, it executes the function with ("/") path.

	http.ListenAndServe(":9090", nil) // surprise! a web server
	// $ go run main.go
	// test with
	// $ curl -v localhost:9090 or $ curl -v localhost:9090/goodbye
}
