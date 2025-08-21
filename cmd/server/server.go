package main

import (
	"log"
	"net/http"

	"github.com/bamorim/gogenapidemo/api"
	"github.com/bamorim/gogenapidemo/service"
)

func main() {
	service, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := api.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
