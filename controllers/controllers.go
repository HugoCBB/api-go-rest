package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HugoCBB/api-go-rest/database"
	"github.com/HugoCBB/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidade(w http.ResponseWriter, r *http.Request) {
	/*
		database.DB.First(&p)
		Retorna todos os registro da tabela personalidades


	*/

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	/*
		database.DB.First(&p, 1)
		Retorna o primeiro registro da tabela personalidades atraves do ID

	*/

	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func CriarNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p) // Decodifica o JSON e armazena na variavel p
	database.DB.Create(&p)             // Cria um novo registro na tabela personalidades
	json.NewEncoder(w).Encode(p)       // Retorna o JSON da variavel p

}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)

}

func EditarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)

}
