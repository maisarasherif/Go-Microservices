package main

import (
	"context"
	"log" //http package for creating web servers
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/maisarasherif/Go-Microservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "\nproduct api: ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting the server on port 9090.........")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Receieved Terminate, graceful shutdown........... \n", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
