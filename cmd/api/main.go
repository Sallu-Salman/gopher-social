package main

import (
	"log"

	"sallu.com/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	err := app.run(mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}