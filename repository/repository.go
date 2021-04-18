package repository

import (
	"github.com/guilhermechaddad/transactions-golang/model"
	"github.com/guilhermechaddad/transactions-golang/repository/memory"
)

type AccountRepository interface {
	SaveAccount(name, documentNumber string) (int64, error)
	GetAllAccounts() []model.Account
	GetAccountById(id int64) *model.Account
	DeleteAccount(id int64) error
	UpdateAccount(id int64, name, documentNumber string) error
}

func GetAccountRepository() AccountRepository {
	return memory.AccountRepo{}
}
