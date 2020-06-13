package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//Functions thats dealing with a single product as methods
func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implements")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implements")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not Implements")
}

func getProduct(db *sql.DB, stat int, count int) ([]product, error) {
	return nil, errors.New("Not Implements")
}
