package main

import (
	"log"
	"os"
	"testing"
)

var a App

//Creating a table
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func TestMain(m *testing.M) {
	a.Initialize(
		//set variables to connect
		os.Getenv("APP_DB_USERNAME"), //export APP_DB_USERNAME=postgres
		os.Getenv("APP_DB_PASSWORD"), //export APP_DB_PASSWORD=password1
		os.Getenv("APP_DB_NAME"),     //export APP_DB_NAME=postgres
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

//Function to make sure that the table we need for testing is available.
func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

//ClearTable to clean up the database
func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_sew RESTART WITH 1")
}
