package models

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var DB *sql.DB

type Order struct {
	Name    string
	Mobile  string
	Ordered bool
	items   map[string]int
}

func ReadCSV() [][]string {
	fmt.Println()
	file, err := os.Open("./babra.csv")
	if err != nil {
		log.Print("Error while opening file...")
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ItemsTable(items []string) {
	DB.Exec(`CREATE TABLE IF NOT EXISTS items (name TEXT NOT NULL)`)
	stmt := []string{}
	for _, item := range items {
		stmt = append(stmt, fmt.Sprintf(`('%v')`, item)) // '%v' single quotes are important here
	}
	finalstmt := fmt.Sprintf(`INSERT INTO items (name) VALUES %v;`, strings.Join(stmt, ", "))
	_, err := DB.Exec(finalstmt)
	if err != nil {
		fmt.Println("Error creating items table")
		fmt.Println(err)
	}
}

// this currently does not used but probably we can fetch the items to be
// dynamic
func GetItems() (any, error) {
	rows, err := DB.Query("SELECT * FROM orders;")
	if err != nil {
		fmt.Println("error while getting rows")
		return nil, err
	}
	_, err = rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	return nil, nil
}

func ParseHeader(header []string) {
	stringProps := []string{"name", "mobile"}
	boolProps := []string{"ordered"}
	cols := ""
	items := []string{}
	exc := []string{"name", "mobile", "ordered"}
	for _, v := range header {
		if !slices.Contains(exc, v) {
			items = append(items, v)
		}
		if slices.Contains(stringProps, v) {
			cols += fmt.Sprintf(`%v TEXT DEFAULT '', `, v)
		} else if slices.Contains(boolProps, v) {
			cols += fmt.Sprintf(`%v INT DEFAULT 0`, v)
		} else {
			cols += fmt.Sprintf(`%v REAL DEFAULT 0.0, `, v)
		}
	}
	ItemsTable(items)
	stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS orders (%v)`, cols)
	DB.Exec(stmt)
}

// add data to database
func AddData(header [][]string, data [][]string) (any, error) {
	head := header[0]
	stmt := fmt.Sprintf("INSERT INTO orders(%v) VALUES ", strings.Join(head, ", "))
	places := make([]string, len(head))
	for idx := range places {
		places[idx] = "?"
	}
	placeholder := "(" + strings.Join(places, ", ") + "), "

	for range data {
		stmt += placeholder
	}

	flattenData := flatten(data)

	stmt = stmt[0 : len(stmt)-2]

	prepared, err := DB.Prepare(stmt)
	if err != nil {
		log.Println("Error while preparing statement")
		log.Println(err)
		return nil, err
	}

	_, err = prepared.Exec(flattenData...)
	if err != nil {
		log.Println(err)
	}
	return nil, nil
}

func flatten(data [][]string) []any {
	var out []any
	for _, row := range data {
		for _, v := range row {
			out = append(out, v) // v is string → any
		}
	}
	return out
}
