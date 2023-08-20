package main

import (
	"github.com/edersonSouza02/gin-api-rest/database"
	"github.com/edersonSouza02/gin-api-rest/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
