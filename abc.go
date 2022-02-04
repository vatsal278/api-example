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
	router.HandleFunc("/passed", printpassedtest).Methods("GET")
	http.Handle("/", router)


	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

func PrintHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "helloworld")
}

func printpassedtest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "PassedTest")
}