package main

import (
	"fmt"
	"github.com/guilhermechaddad/transactions-golang/config"
	"github.com/guilhermechaddad/transactions-golang/config/routes"
	"log"
	"net/http"
)

const (
	applicationPort = "7788"
)

func main() {
	fmt.Println("Starting Transactions project")
	infra := new(config.Infrastructure)

	router := routes.CreateRouter(infra)

	log.Fatal(http.ListenAndServe(":"+applicationPort, router))
}