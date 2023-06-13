package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getRequest).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")
	http.Handle("/", r)
	fmt.Println("server started and listening on localhost:9003")
	http.ListenAndServe(":9003", nil)

}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Get request")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Post request")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Delete request")
}
