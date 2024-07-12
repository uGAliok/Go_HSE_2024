package models

import (
	"errors"
	"fmt"
)

type Account struct {
	Name   string
	Amount int
}

type BankDatabase struct {
	data map[string]*Account
}

func TheDatabase() *BankDatabase {
	return &BankDatabase{make(map[string]*Account)}
}

// Аналог handler No2, почему-то в отдельном файле эта часть не хотела работать
func (db BankDatabase) CreateAccount(name string, amount int64) error {
	if len(name) == 0 {
		return errors.New("Invalid account name")
	}
	if amount < 0 {
		return errors.New("Invalid account amount")
	}
	if _, ok := db.data[name]; ok {
		return errors.New(fmt.Sprintf("Account '%s' already exists", name))
	}
	db.data[name] = &Account{name, 0}
	db.data[name].Amount = int(amount)
	return nil
}

func (db *BankDatabase) DeleteAccount(name string) error {
	if len(name) == 0 {
		return errors.New("Invalid account name")
	}
	if _, ok := db.data[name]; !ok {
		return errors.New(fmt.Sprintf("Account '%s' does not exist", name))
	}
	delete(db.data, name)
	return nil
}

func (db *BankDatabase) Patch(name string, amount int64) error {
	if len(name) == 0 {
		return errors.New("Invalid account name")
	}
	if amount < 0 {
		return errors.New("Invalid account amount")
	}
	if _, ok := db.data[name]; !ok {
		return errors.New(fmt.Sprintf("Account '%s' does not exist", name))
	}
	db.data[name].Amount = int(amount)
	return nil
}

func (db *BankDatabase) UpdateName(name string, newName string) error {
	if len(name) == 0 {
		return errors.New("Invalid account name")
	}
	if len(newName) == 0 {
		return errors.New("Invalid account name")
	}
	if _, ok := db.data[name]; !ok {
		return errors.New(fmt.Sprintf("Account '%s' does not exist", name))
	}
	account := db.data[name]
	account.Name = newName
	delete(db.data, name)
	db.data[newName] = account
	return nil
}

func (db *BankDatabase) GetAccount(name string) (*Account, error) {
	if len(name) == 0 {
		return nil, errors.New("Invalid account name")
	}
	if _, ok := db.data[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Account '%s' does not exist", name))
	}
	return db.data[name], nil
}
