package routes

import (
	"github.com/gorilla/mux"
	"github.com/guilhermechaddad/transactions-golang/config"
	"github.com/guilhermechaddad/transactions-golang/controller"
	"net/http"
)

const (
	methodGet = "GET"
	methodPost = "POST"
	methodPut = "PUT"
	methodDelete = "DELETE"

	getAllSuffix = "_GetAllMethod"
	getByIdSuffix = "_GetByIdMethod"
	createSuffix = "_CreateMethod"
	deleteSuffix = "_DeleteMethod"
	updateSuffix = "_UpdateMethod"
)

func CreateRouter(i *config.Infrastructure) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, crud := range getCRUDControllers(i) {
		router.Name(crud.GetName() + getAllSuffix).
			Path(crud.GetAllPath()).
			Methods(methodGet).
			Handler(http.HandlerFunc(crud.GetAll))


		router.Name(crud.GetName() + getByIdSuffix).
			Path(crud.GetByIdPath()).
			Methods(methodGet).
			Handler(http.HandlerFunc(crud.GetById))


		router.Name(crud.GetName() + createSuffix).
			Path(crud.CreatePath()).
			Methods(methodPost).
			Handler(http.HandlerFunc(crud.Create))


		router.Name(crud.GetName() + updateSuffix).
			Path(crud.UpdatePath()).
			Methods(methodPut).
			Handler(http.HandlerFunc(crud.Update))


		router.Name(crud.GetName() + deleteSuffix).
			Path(crud.DeletePath()).
			Methods(methodDelete).
			Handler(http.HandlerFunc(crud.Delete))
	}

	return router
}

func getCRUDControllers(i *config.Infrastructure) []controller.CRUD {
	var c []controller.CRUD

	c = append(c, i.GetAccountController())

	return c
}
