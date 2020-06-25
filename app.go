package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App the struct of app
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize some thing
func (a *App) Initialize(user string, password string, dbname string) {

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error

	//connecting to postgres
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()

}

//Run this
func (a *App) Run(addr string) {}
