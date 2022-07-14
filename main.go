package main

import (
	"fmt"

	"github.com/ernanimusa/go-rest-api/database"
	"github.com/ernanimusa/go-rest-api/models"
	"github.com/ernanimusa/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Name 1", Historia: "hystory 1"},
		{Id: 2, Nome: "Segundo Nome", Historia: "Para uma segunda hystory"},
	}

	database.ConectaBD()

	fmt.Println("Requisição para o servidor")
	routes.HandleResquest()
}
