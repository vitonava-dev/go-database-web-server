package main

import (
	"fmt"
	"go-database-webserver/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Starting application...")
	r := mux.NewRouter()
	fmt.Println("Registering routes...")
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on the port 9010...")
	log.Fatal(http.ListenAndServe(":9010", nil))

}
