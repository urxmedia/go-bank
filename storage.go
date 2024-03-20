package main

import "database/sql"

type Storage interface {
	GetAccountByID(string) (*Account, error)
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	UpdateAccountByID(string) error
	DeleteAccountByID(string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	
}