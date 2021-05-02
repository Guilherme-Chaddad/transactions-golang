package config

import (
	"github.com/guilhermechaddad/transactions-golang/controller"
	"github.com/guilhermechaddad/transactions-golang/repository"
	"github.com/guilhermechaddad/transactions-golang/service"
)

func (i Infrastructure) getAccountController() controller.CRUD {
	if i.accountController == nil {
		i.accountController = controller.NewAccountController(i.getAccountService())
	}
	return i.accountController
}

func (i Infrastructure) getAccountService() service.AccountService {
	if i.accountService == nil {
		i.accountService = service.NewAccountService(i.getAccountRepository())
	}
	return i.accountService
}

func (i Infrastructure) getAccountRepository() repository.AccountRepository {
	if i.accountRepository == nil {
		i.accountRepository = repository.GetAccountRepository()
	}
	return i.accountRepository
}
