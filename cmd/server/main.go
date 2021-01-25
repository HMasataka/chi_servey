package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	s "github.com/HMasataka/chi_survey"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	petStore := s.NewPets()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := s.HandlerFromMux(petStore, r)
	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	log.Fatal(s.ListenAndServe())
}
