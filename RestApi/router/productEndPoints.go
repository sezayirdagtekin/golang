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
	ProductId   int     `json:"id"`
	ProductName string  `json:"name"`
	Inventory   int     `json:"inventory"`
	Price       float64 `json:"price"`
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
	connStr := getConnectionString()
	products, err := getAllProducts(connStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusOK)

}

func getAllProducts(connStr string) ([]Product, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	rows, err := db.Query("SELECT productid ,productname ,inventory ,price  FROM  product p ")

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		rows.Scan(&p.ProductId, &p.ProductName, &p.Inventory, &p.Price)
		products = append(products, p)
		fmt.Println("Product:", p.ProductId, p.ProductName, p.Price, p.Inventory)
	}
	return products, nil
}
