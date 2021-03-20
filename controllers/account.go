package controllers

import (
	"fmt"
	"net/http"
)

const (
	name = "Account Handler"
	accountBasePath = "/accounts"
	accountIdPath = "/{id}"
)

type AccountController struct {
}

func NewAccountController() AccountController {
	return AccountController{}
}

func (a AccountController) GetName() string {
	return name
}

func (a AccountController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Getting all accounts")
}

func (a AccountController) GetAllPath() string {
	return accountBasePath
}

func (a AccountController) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get account by Id")
}

func (a AccountController) GetByIdPath() string {
	return accountBasePath + accountIdPath
}

func (a AccountController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating account")
}

func (a AccountController) CreatePath() string {
	return accountBasePath
}

func (a AccountController) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Updating account")
}

func (a AccountController) UpdatePath() string {
	return accountBasePath + accountIdPath
}

func (a AccountController) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Deleting account")
}

func (a AccountController) DeletePath() string {
	return accountBasePath + accountIdPath
}