package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/guilhermechaddad/transactions-golang/api"
	"github.com/guilhermechaddad/transactions-golang/service"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	name = "AccountService Handler"
	accountBasePath = "/accounts"
	accountIdPath = "/{id:[0-9]+}"
)

type AccountController struct {
	service service.AccountService
}

func NewAccountController() AccountController {
	return AccountController{
		service: service.NewAccountService(),
	}
}

func (a AccountController) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts := a.service.GetAllAccounts()

	_ = json.NewEncoder(w).Encode(accounts)
}

func (a AccountController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	account := a.service.GetAccountById(int64(id))

	if account == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "AccountService not found for id %d", id)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (a AccountController) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var account api.Account
	err := json.Unmarshal(reqBody, &account)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid payload. Error: %s", err.Error())
		return
	}

	newId, err := a.service.CreateAccount(account.DocumentNumber, account.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to create account. %s", err.Error())
		return
	}

	account.AccountId = newId
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)

}

func (a AccountController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	var account api.Account
	err := json.Unmarshal(reqBody, &account)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid payload. Error: %s", err.Error())
		return
	}

	err = a.service.UpdateAccount(int64(id), account.DocumentNumber, account.Name)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not update AccountService: %s.", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "AccountService updated Successfully")
}

func (a AccountController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := a.service.DeleteAccount(int64(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not delete AccountService: %s.", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "AccountService deleted successfully")
}

func (a AccountController) GetName() string {
	return name
}

func (a AccountController) GetAllPath() string {
	return accountBasePath
}

func (a AccountController) GetByIdPath() string {
	return accountBasePath + accountIdPath
}

func (a AccountController) CreatePath() string {
	return accountBasePath
}

func (a AccountController) UpdatePath() string {
	return accountBasePath + accountIdPath
}

func (a AccountController) DeletePath() string {
	return accountBasePath + accountIdPath
}