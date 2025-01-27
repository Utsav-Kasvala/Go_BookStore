package main


import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Utsav_Kasvala/Go_BookStore/pkg/routes"
	"github.com/jinzhu/gorm/dialects/mysql"
)


func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
