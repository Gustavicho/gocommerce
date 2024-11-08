package main

import (
	"log"

	"github.com/Gustavicho/gocommerce/cmd/api"
)

func main() {
	server := api.NewAPIService(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}