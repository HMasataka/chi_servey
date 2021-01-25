package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	s "github.com/HMasataka/chi_survey"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func customMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	petStore := s.NewPets()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware)

	h := s.HandlerFromMux(petStore, r)
	s := &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	log.Fatal(s.ListenAndServe())
}
