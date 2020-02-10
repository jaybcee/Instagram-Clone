package main

import (
	"log"

	"github.com/fvbock/endless"
)

func api() {
	r := createRoutes()
	r.run(":3030")
	//if err := endless.ListenAndServe("localhost:3030", r); err != nil {
	//	log.Fatal(err)
	//}
}
