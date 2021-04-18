package service

import (
	"errors"
	"github.com/guilhermechaddad/transactions-golang/api"
	"github.com/guilhermechaddad/transactions-golang/repository"
)

type AccountService struct {
	repository repository.AccountRepository
}

func NewAccountService() AccountService {
	return AccountService{
		repository: repository.GetAccountRepository(),
	}
}

func (a AccountService) GetAccountById(id int64) *api.Account{
	accountFromDb := a.repository.GetAccountById(id)
	if accountFromDb == nil {
		return nil
	}
	return &api.Account{AccountId:accountFromDb.AccountId, Name:accountFromDb.Name, DocumentNumber: accountFromDb.DocumentNumber}
}

func (a AccountService) GetAllAccounts() []api.Account {
	var accountsToReturn []api.Account
	for _, accountFromDb := range a.repository.GetAllAccounts() {
		accountsToReturn = append(accountsToReturn, api.Account{AccountId:accountFromDb.AccountId, Name:accountFromDb.Name, DocumentNumber: accountFromDb.DocumentNumber})
	}
	return accountsToReturn
}

func (a AccountService) CreateAccount(documentNumber, name string) (int64, error) {
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

 	return a.repository.SaveAccount(name, documentNumber)
}

func (a AccountService) UpdateAccount(id int64, documentNumber, name string) error {
	return a.repository.UpdateAccount(id, name, documentNumber)
}

func (a AccountService) DeleteAccount(id int64) error {
	return a.repository.DeleteAccount(id)
}