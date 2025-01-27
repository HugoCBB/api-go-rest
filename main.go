package main

import (
	"fmt"

	"github.com/HugoCBB/api-go-rest/database"
	"github.com/HugoCBB/api-go-rest/routers"
)

func main() {
	database.ConectDataBase()
	fmt.Println("Iniciando servidor rest em go....")
	routers.HandleRequest()
}
