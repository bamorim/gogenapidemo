package main

import (
	"fmt"
	"log"
	"net/http"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/bamorim/gogenapidemo/api"
	"github.com/bamorim/gogenapidemo/service"
)

func main() {
	svc, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := api.NewServer(svc)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", srv)
	mux.HandleFunc("/docs/", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecContent: service.GetOpenAPISpec(),
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})
		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
