package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)
	})

	return r
}

func (a *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr: a.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}
	
	fmt.Println("Server started at port", a.config.addr)
	return srv.ListenAndServe()
}