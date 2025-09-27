package main

import (
	"handlers"
	"log" //http package for creating web servers
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)

}
