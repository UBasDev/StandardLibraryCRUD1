package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/goproject6/handlers"
)

func main() {

	logger := log.New(os.Stdout, "hello-api", log.LstdFlags)

	helloHandler := handlers.NewProducts(logger)
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serverCreated := &http.Server{
		Addr:         ":8081",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := serverCreated.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel1 := make(chan os.Signal)
	signal.Notify(signalChannel1, os.Interrupt)
	signal.Notify(signalChannel1, os.Kill)
	sig1 := <-signalChannel1
	logger.Println("Received terminate, graceful shutdown", sig1)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	serverCreated.Shutdown(timeoutContext)
}
