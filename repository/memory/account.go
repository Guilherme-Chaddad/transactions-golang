package memory

import (
	"errors"
	"github.com/guilhermechaddad/transactions-golang/model"
	"time"
)

var dbAccounts = []model.Account{
	{
		AccountId:      1,
		DocumentNumber: "123456789",
		Name: "Name AccountController",
		CreationDate: time.Now(),
	},
	{
		AccountId:      2,
		DocumentNumber: "987654321",
		Name: "Name AccountController 2",
		CreationDate: time.Now(),
	},
}

type AccountRepo struct {}

func (ar AccountRepo) SaveAccount(name, documentNumber string) (int64, error) {
	newId := int64(len(dbAccounts)+1)
	dbAccounts = append(dbAccounts, model.Account{AccountId:newId, Name: name, DocumentNumber:documentNumber, CreationDate: time.Now()})

	return newId, nil
}

func (ar AccountRepo) GetAllAccounts() []model.Account {
	return dbAccounts
}

func (ar AccountRepo) GetAccountById(id int64) *model.Account {
	for _, account := range dbAccounts {
		if account.AccountId == id {
			return &account
		}
	}
	return nil
}

func (ar AccountRepo) DeleteAccount(id int64) error {
	delIndex := -1
	for index, account := range dbAccounts {
		if account.AccountId == id {
			delIndex = index
			break
		}
	}
	if delIndex < 0 {
		return errors.New("account was not found")
	}

	copy(dbAccounts[delIndex:], dbAccounts[delIndex+1:])
	dbAccounts = dbAccounts[:len(dbAccounts)-1]

	return nil
}

func (ar AccountRepo) UpdateAccount(id int64, name, documentNumber string) error {
	for index, account := range dbAccounts {
		if account.AccountId == id {
			if name != "" {
				dbAccounts[index].Name = name
			}

			if documentNumber != "" {
				dbAccounts[index].DocumentNumber = documentNumber
			}
			return nil
		}
	}
	return errors.New("account was not found")
}