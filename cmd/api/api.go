package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"sallu.com/internal/store"
)

type application struct {
	config config
	store store.Storage
}

type dbConfig struct {
	addr string
}

type config struct {
	addr string
	db dbConfig
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