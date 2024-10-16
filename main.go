package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mircroservice/m/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	eh := handlers.NewEngine(l)

	sm := http.NewServeMux()
	sm.Handle("/", eh)

	http.ListenAndServe(":9090", sm)
}
