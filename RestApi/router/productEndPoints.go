package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Product struct {
	ProductId int     `json:"Id"`
	Name      string  `json:"Name"`
	Inventory int     `json:"Inventory"`
	Price     float64 `json:"Price"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", getProducts).Methods("GET")
	http.Handle("/", r)
	fmt.Println("server started and listening on localhost:9003")
	http.ListenAndServe(":9003", nil)

}

func getConnectionString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("getProducts...")
	connStr := getConnectionString()

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT productid ,productname ,inventory ,price  FROM  product p ")

	if err != nil {
		log.Fatal(err.Error())

	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		rows.Scan(&p.ProductId, &p.Name, &p.Inventory, &p.Price)
		products = append(products, p)
		fmt.Println("Product:", p.ProductId, p.Name, p.Price, p.Inventory)
	}
	response, _ := json.Marshal(products)
	w.Write(response)
	w.WriteHeader(http.StatusOK)

}
