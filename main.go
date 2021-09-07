package main

import (
	"log"
	"net/http"
	
	"github.com/luisalfredosv/golang-gorilla/routes"
	"github.com/luisalfredosv/golang-gorilla/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.MigrateDB()
	r := mux.NewRouter()
	routes.SetContactsRoutes(r)

	srv := http.Server{
		Addr: ":4000",
		Handler: r,
	}

	log.Println("Runing on port  :4000")
	log.Fatal(srv.ListenAndServe())
}