package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/israelowusu/bookstore_management.git/pkg/docs"
	"github.com/israelowusu/bookstore_management.git/pkg/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := mux.NewRouter()

	// Swagger documentation
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
