package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	s "github.com/HMasataka/chi_survey"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	petStore := s.NewPets()
	h := s.Handler(petStore)

	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	log.Fatal(s.ListenAndServe())
}
