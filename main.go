package main

import (
	"fmt"
	"github.com/guilhermechaddad/transactions-golang/api"
	"github.com/guilhermechaddad/transactions-golang/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	applicationPort = "7788"
)

func main() {
	fmt.Println("Starting Transactions project")
	api.Accounts = []api.Account{
		{
			AccountId:      1,
			DocumentNumber: "123456789",
			Name: "Name Account",
		},
		{
			AccountId:      2,
			DocumentNumber: "987654321",
			Name: "Name Account 2",
		},
	}
	router := createRouter()

	log.Fatal(http.ListenAndServe(":"+applicationPort, router))
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, crud := range getCRUDControllers() {
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

func getCRUDControllers() []controller.CRUD {
	var c []controller.CRUD

	c = append(c, controller.NewAccountController())

	return c
}