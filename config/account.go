package config

import (
	"github.com/guilhermechaddad/transactions-golang/controller"
	"github.com/guilhermechaddad/transactions-golang/repository"
	"github.com/guilhermechaddad/transactions-golang/service"
)

func (i Infrastructure) GetAccountController() controller.CRUD {
	if i.accountController == nil {
		i.accountController = controller.NewAccountController(i.GetAccountService())
	}
	return i.accountController
}

func (i Infrastructure) GetAccountService() service.AccountService {
	if i.accountService == nil {
		i.accountService = service.NewAccountService(i.GetAccountRepository())
	}
	return i.accountService
}

func (i Infrastructure) GetAccountRepository() repository.AccountRepository {
	if i.accountRepository == nil {
		i.accountRepository = repository.GetAccountRepository()
	}
	return i.accountRepository
}
