package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krouta1/go-sql-bookstore-api/pkg/routes"
	_ "gorm.io/driver/sqlite"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
