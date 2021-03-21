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
	name = "Account Handler"
	accountBasePath = "/accounts"
	accountIdPath = "/{id:[0-9]+}"
)

type Account struct {
	service service.Account
}

func NewAccountController() Account {
	return Account{
		service: service.NewAccountService(),
	}
}

func (a Account) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts := a.service.GetAllAccounts()

	_ = json.NewEncoder(w).Encode(accounts)
}

func (a Account) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	account := a.service.GetAccountById(int64(id))

	if account == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Account not found for id %d", id)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (a Account) Create(w http.ResponseWriter, r *http.Request) {
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

func (a Account) Update(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "Could not update Account: %s.", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Account updated Successfully")
}

func (a Account) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	err := a.service.DeleteAccount(int64(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Could not delete Account: %s.", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Account deleted successfully")
}

func (a Account) GetName() string {
	return name
}

func (a Account) GetAllPath() string {
	return accountBasePath
}

func (a Account) GetByIdPath() string {
	return accountBasePath + accountIdPath
}

func (a Account) CreatePath() string {
	return accountBasePath
}

func (a Account) UpdatePath() string {
	return accountBasePath + accountIdPath
}

func (a Account) DeletePath() string {
	return accountBasePath + accountIdPath
}