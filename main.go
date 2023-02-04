package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
	)
	flag.Parse()

	fmt.Println("Repository: In progress")

	productRepo, err := NewProductRepo("products.csv", "partners.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Repository: Ready")

	fmt.Println("Endpoints and handlers: In progress")

	pricingService := NewPricingService(productRepo)
	totalRetailPriceHandler := MakeTotalRetailPriceHttpHandler(pricingService)

	// use chi
	rtr := chi.NewRouter()
	rtr.Post("/retail", totalRetailPriceHandler.ServeHTTP)

	fmt.Println("Endpoints and handlers: Ready")

	fmt.Printf("Hosting on %s\n", *listen)

	http.ListenAndServe(*listen, rtr)
}
