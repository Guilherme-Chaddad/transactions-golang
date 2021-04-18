package model

import "time"

type Account struct {
	AccountId int64
	DocumentNumber string
	Name string
	CreationDate time.Time
}
