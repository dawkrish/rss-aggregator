package main

import (
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

	srv := http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err",err500)
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
