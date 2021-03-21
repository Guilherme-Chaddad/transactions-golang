package service

import (
	"errors"
	"github.com/guilhermechaddad/transactions-golang/api"
)

type Account struct {}

func NewAccountService() Account {
	return Account{}
}

func (a Account) GetAccountById(id int64) *api.Account{

	for _, account := range api.Accounts {
		if account.AccountId == id {
			return &account
		}
	}

	return nil
}

func (a Account) GetAllAccounts() []api.Account {
	return api.Accounts
}

func (a Account) CreateAccount(documentNumber, name string) (int64, error) {
	var errorMessage string
	if documentNumber == "" {
		errorMessage += "Document number is required. "
	}
	if name == "" {
		errorMessage += "Name is required."
	}

	if errorMessage != ""{
		return -1, errors.New(errorMessage)
	}

	newId := int64(len(api.Accounts)+1)
	api.Accounts = append(api.Accounts, api.Account{AccountId:newId, Name: name, DocumentNumber:documentNumber})

 	return newId, nil
}

func (a Account) UpdateAccount(id int64, documentNumber, name string) error {
	for index, account := range api.Accounts {
		if account.AccountId == id {
			if name != "" {
				api.Accounts[index].Name = name
			}

			if documentNumber != "" {
				api.Accounts[index].DocumentNumber = documentNumber
			}
			return nil
		}
	}
	return errors.New("account was not found")
}

func (a Account) DeleteAccount(id int64) error {
	delIndex := -1
	for index, account := range api.Accounts {
		if account.AccountId == id {
			delIndex = index
			break
		}
	}
	if delIndex < 0 {
		return errors.New("account was not found")
	}

	copy(api.Accounts[delIndex:], api.Accounts[delIndex+1:])
	api.Accounts = api.Accounts[:len(api.Accounts)-1]

	return nil
}