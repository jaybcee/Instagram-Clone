package main

import (
	"log"

	"github.com/fvbock/endless"
)

func api() {
	r := createRoutes()

	if err := endless.ListenAndServe("localhost:3030", r); err != nil {
		log.Fatal(err)
	}
}
