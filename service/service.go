package service

import (
	"github.com/guilhermechaddad/transactions-golang/api"
)

type AccountService interface {
	GetAccountById(id int64) *api.Account
	GetAllAccounts() []api.Account
	CreateAccount(documentNumber, name string) (int64, error)
	UpdateAccount(id int64, documentNumber, name string) error
	DeleteAccount(id int64) error
}


