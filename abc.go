package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/helloworld", PrintHelloWorld).Methods("GET")
	router.HandleFunc("/letsroll", letsroll).Methods("GET")
	http.Handle("/", router)


	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

func PrintHelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "helloworld")
}

func letsroll(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "letsroll")
}