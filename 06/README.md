# 06

```sh
go mod init github.com/kangkyu/go-kit-fundamentals/06
```

Write endpoint.go main.go repo.go service.go transport.go

```go
// main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
	)
	flag.Parse()

	fmt.Println(*listen)
}
```

```sh
go run main.go -listen=:8081
# :8081
go run main.go
# :8080
```

```sh
# use go-kit & chi
go get github.com/go-kit/kit
go get -u github.com/go-chi/chi/v5
```

And then update `main.go`

```go
// main.go
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
```

```sh
go build -o six .
./six
# Hosting on :8080
```
