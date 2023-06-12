package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go rest api\n")
}

func main() {
	http.HandleFunc("/", greeting)
	fmt.Println("server started and listening on localhost:9003")
	http.ListenAndServe(":9003", nil)
}
