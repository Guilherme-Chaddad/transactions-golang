package main

import (
	"fmt"
	"github.com/guilhermechaddad/transactions-golang/controllers"
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

	router := createRouter()

	http.ListenAndServe(":"+applicationPort, router)
}

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range getCRUDControllers() {
		router.Name(route.GetName() + getAllSuffix).
			Path(route.GetAllPath()).
			Methods(methodGet).
			Handler(http.HandlerFunc(route.GetAll))


		router.Name(route.GetName() + getByIdSuffix).
			Path(route.GetByIdPath()).
			Methods(methodGet).
			Handler(http.HandlerFunc(route.GetById))


		router.Name(route.GetName() + createSuffix).
			Path(route.CreatePath()).
			Methods(methodPost).
			Handler(http.HandlerFunc(route.Create))


		router.Name(route.GetName() + updateSuffix).
			Path(route.UpdatePath()).
			Methods(methodPut).
			Handler(http.HandlerFunc(route.Update))


		router.Name(route.GetName() + deleteSuffix).
			Path(route.DeletePath()).
			Methods(methodDelete).
			Handler(http.HandlerFunc(route.Delete))
	}

	return router
}

func getCRUDControllers() []controllers.CRUDController {
	var c []controllers.CRUDController

	c = append(c, controllers.NewAccountController())

	return c
}