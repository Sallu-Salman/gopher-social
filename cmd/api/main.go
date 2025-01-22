package main

import (
	"log"
	_ "github.com/lib/pq"

	"sallu.com/internal"
	"sallu.com/internal/env"
	"sallu.com/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: "postgres://postgres:password@localhost/gopherdb?sslmode=disable",
		},
	}
	db, err := internal.New(cfg.db.addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	store:= store.NewStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	err = app.run(mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}