package main

import (
	"github.com/felixlambertv/nextalent/config"
	"github.com/felixlambertv/nextalent/routes"
	"log"
)

func main() {
	db := config.InitDB()
	r := routes.SetupRouter(db)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
