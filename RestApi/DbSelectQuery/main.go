package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	log.Println("main...")
	connStr := getConnectionString()
	//log.Println("connection info:", connStr)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT productid ,productname ,inventory ,price  FROM  product p ")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)

		fmt.Println("Product:", p.id, p.name, p.price, p.inventory)
	}
}

func getConnectionString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

type Product struct {
	id        int
	name      string
	inventory int
	price     float64
}
