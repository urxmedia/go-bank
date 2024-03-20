package main

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	mathRand "math/rand"
)

type Account struct {
	ID        string  `json:"accountID"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Number    int64   `json:"accountNumber"`
	Balance   float64 `json:"accountBalance"`
}

func GenerateAccountID() (string, error) {
	idBytes := make([]byte, 16)
	_, err := cryptoRand.Read(idBytes)
	if err != nil {
		return "", fmt.Errorf("Error while generating Account ID: %s", err.Error())
	}
	return hex.EncodeToString(idBytes), nil
}

func NewAccount(firstName, lastName string) (*Account, error) {
	accountID, err := GenerateAccountID()
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:        accountID,
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(mathRand.Intn(1000000)),
	}, nil
}
