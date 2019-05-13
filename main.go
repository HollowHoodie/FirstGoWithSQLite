package main

import (
	"database/sql"
	"fmt"

	"./helpers"

	_ "github.com/mattn/go-sqlite3"
)

type people struct {
	firstname string
	lastname  string
}

func initDatabase() bool {
	database, _ := sql.Open("sqlite3", "./database.db")
	sqlMessage := helpers.ReadFile("./migrations/database.sql")
	statement, _ := database.Prepare(sqlMessage)
	statement.Exec()
	return true
}

func seedData() bool {
	var peoples = []people{
		people{"Jon", "Snow"},
		people{"Darth", "Vader"},
	}
	database, _ := sql.Open("sqlite3", "./database.db")
	statement, _ := database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	for _, people := range peoples {
		statement.Exec(people.firstname, people.lastname)
	}
	return true
}

func getData() *sql.Rows {
	database, _ := sql.Open("sqlite3", "./database.db")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	return rows
}

func main() {
	initDatabase()
	seedData()
	data := getData()

	var id int
	var firstname string
	var lastname string
	for data.Next() {
		data.Scan(&id, &firstname, &lastname)
		fmt.Println(firstname, lastname)
	}

}
