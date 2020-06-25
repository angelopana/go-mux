package main

import (
	"log"
	"net/http"
	"net/http/httptest"
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

//TestEmpyTable testing the response to the /products endpoint with an empty table.
//This test deletes all records from the products table and sends a GET request to the /products endpoint.
func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

//ExecuteRequest this function executes the request using the application’s router and returns the response.
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

//CheckResponseCode function to check the response
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
