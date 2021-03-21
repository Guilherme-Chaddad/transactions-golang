package api

type Account struct {
	AccountId int64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
	Name string `json:"name"`
}

var Accounts []Account



