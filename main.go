package main

import (
	"fmt"

	"github.com/HugoCBB/api-go-rest/models"
	"github.com/HugoCBB/api-go-rest/routers"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Hugo", Historia: "Hugo é um cara legal"},
		{Id: 2, Nome: "Maria", Historia: "Maria é uma mulher legal"},
	}

	fmt.Println("Iniciando servidor rest em go....")
	routers.HandleRequest()
}
