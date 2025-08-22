package main

import (
	"database/sql"
	"fmt"
	"log"

	"billing.chaitanya.observer/internals/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:billing.db")
	if err != nil {
		log.Println("Error while opening DB")
		log.Fatal(err)
	}
	err = db.Ping()
	defer db.Close()
	if err != nil {
		db.Close()
		log.Println("Error while connecting to DB")
		log.Fatal(err)
	}

	models.DB = db

	data := models.ReadCSV()
	models.ParseHeader(data[:1][0])
	models.AddData(data[:1], data[1:])
	data = data[1:]

	_, err = models.GetItems()
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
}
