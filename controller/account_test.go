package controller

import (
	"github.com/guilhermechaddad/transactions-golang/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountControllerPaths(t *testing.T) {
	acc := NewAccountController(service.Service{})

	assert.NotNil(t, acc.service)
	assert.Equal(t, accountBasePath, acc.GetAllPath())
	assert.Equal(t, accountBasePath, acc.CreatePath())
	assert.Equal(t, accountBasePath + accountIdPath, acc.GetByIdPath())
	assert.Equal(t, accountBasePath + accountIdPath, acc.UpdatePath())
	assert.Equal(t, accountBasePath + accountIdPath, acc.DeletePath())
	assert.Equal(t, name, acc.GetName())
}
