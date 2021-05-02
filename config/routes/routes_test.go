package routes

import (
	"github.com/guilhermechaddad/transactions-golang/config"
	"github.com/guilhermechaddad/transactions-golang/controller"
	"net/http"
	"testing"
)

const (
	mockControllerName = "/mock-controller-name"
	getAllPath = "/get-all-path"
	getByIdPath = "/get-by-id-path"
	createPath = "/"
	updatePath = "/update-path"
	deletePath = "/delete-path"
)

type mockController struct {}

func (c mockController) GetAll(w http.ResponseWriter, r *http.Request) {}

func (c mockController) GetById(w http.ResponseWriter, r *http.Request) {}

func (c mockController) Create(w http.ResponseWriter, r *http.Request) {}

func (c mockController) Update(w http.ResponseWriter, r *http.Request) {}

func (c mockController) Delete(w http.ResponseWriter, r *http.Request) {}

func (c mockController) GetName() string {
	return mockControllerName
}

func (c mockController) GetAllPath() string {
	return getAllPath
}

func (c mockController) GetByIdPath() string {
	return getByIdPath
}

func (c mockController) CreatePath() string {
	return createPath
}

func (c mockController) UpdatePath() string {
	return updatePath
}

func (c mockController) DeletePath() string {
	return deletePath
}

type mockInfrastructure struct {
	controller controller.CRUD
}

func (m mockInfrastructure) GetCRUDControllers() []controller.CRUD {
	return []controller.CRUD {
		m.controller,
	}
}

func TestCreateRouter(t *testing.T) {
	mock := mockController{}
	tests := []struct {
		name  string
		infra config.InfrastructureInterface
	}{
		{
			name:  "Validate Route Creation",
			infra: mockInfrastructure{mock},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateRouter(tt.infra)

			getAllRoute := got.GetRoute(mockControllerName + getAllSuffix)
			if getAllRoute == nil {
				t.Error("Expected getAllRoute created but it was nil")
			} else {
				pathTemplate, _ := getAllRoute.GetPathTemplate()
				if pathTemplate != mock.GetAllPath() {
					t.Errorf("Expected path: '%s' but it was '%s'", mock.GetAllPath(), pathTemplate)
				}
			}
			getByIdRoute := got.GetRoute(mockControllerName + getByIdSuffix)
			if getByIdRoute == nil {
				t.Error("Expected getByIdRoute created but it was nil")
			} else {
				pathTemplate, _ := getByIdRoute.GetPathTemplate()
				if pathTemplate != mock.GetByIdPath() {
					t.Errorf("Expected path: '%s' but it was '%s'", mock.GetByIdPath(), pathTemplate)
				}
			}
			createRoute := got.GetRoute(mockControllerName + createSuffix)
			if createRoute == nil {
				t.Error("Expected createRoute created but it was nil")
			} else {
				pathTemplate, _ := createRoute.GetPathTemplate()
				if pathTemplate != mock.CreatePath() {
					t.Errorf("Expected path: '%s' but it was '%s'", mock.CreatePath(), pathTemplate)
				}
			}
			updateRoute := got.GetRoute(mockControllerName + updateSuffix)
			if updateRoute == nil {
				t.Error("Expected updateRoute created but it was nil")
			} else {
				pathTemplate, _ := updateRoute.GetPathTemplate()
				if pathTemplate != mock.UpdatePath() {
					t.Errorf("Expected path: '%s' but it was '%s'", mock.UpdatePath(), pathTemplate)
				}
			}
			deleteRoute := got.GetRoute(mockControllerName + deleteSuffix)
			if deleteRoute == nil {
				t.Error("Expected deleteRoute created but it was nil")
			} else {
				pathTemplate, _ := deleteRoute.GetPathTemplate()
				if pathTemplate != mock.DeletePath() {
					t.Errorf("Expected path: '%s' but it was '%s'", mock.DeletePath(), pathTemplate)
				}
			}

		})
	}
}