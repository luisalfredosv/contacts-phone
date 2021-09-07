package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"io"
	"github.com/luisalfredosv/golang-gorilla/routes"
	"github.com/luisalfredosv/golang-gorilla/utils"

)

func main(){

	utils.MigrateDB()

	r := mux.NewRouter()

	routes.SetContactsRoutes(r)

	srv:= &http.Server{
		Handler: r,
		Addr: "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Rining on port  :9100")

	log.Fatal(srv.ListenAndServe())
}