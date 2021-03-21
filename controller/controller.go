package controller

import (
	"net/http"
)

type CRUD interface {
	GetName() string
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllPath() string
	GetById(w http.ResponseWriter, r *http.Request)
	GetByIdPath() string
	Update(w http.ResponseWriter, r *http.Request)
	UpdatePath() string
	Create(w http.ResponseWriter, r *http.Request)
	CreatePath() string
	Delete(w http.ResponseWriter, r *http.Request)
	DeletePath() string
}
