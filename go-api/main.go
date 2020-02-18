package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/serdarkalayci/prometheus-works/go-api/handlers"
	"github.com/serdarkalayci/prometheus-works/go-api/middleware"

	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":6543", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "go-api ", log.LstdFlags)

	// create the handlers
	value := handlers.NewValue(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	sm.Use(middleware.MonitoringMiddleware)
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/values", value.GetValues)
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/values/{id:[0-9]+}", value.PutValue)
	//putRouter.Use(value.MiddlewareValidateProduct)
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/values", value.PostValue)
	//postRouter.Use(value.MiddlewareValidateProduct)

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	sm.PathPrefix("/metrics").Handler(promhttp.Handler())
	prometheus.MustRegister(middleware.RequestCounterVec)
	prometheus.MustRegister(middleware.RequestDurationGauge)

	// start the server
	go func() {
		l.Println("Starting server on port 6543")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
