package config

import (
	"github.com/guilhermechaddad/transactions-golang/controller"
	"github.com/guilhermechaddad/transactions-golang/repository"
	"github.com/guilhermechaddad/transactions-golang/service"
)

type Infrastructure struct {
	accountController controller.CRUD
	accountService service.AccountService
	accountRepository repository.AccountRepository
}

type InfrastructureInterface interface {
	GetCRUDControllers() []controller.CRUD
}

func (i Infrastructure) GetCRUDControllers() []controller.CRUD {
	return []controller.CRUD {
		i.getAccountController(),
	}
}