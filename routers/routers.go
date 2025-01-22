package routers

import (
	"log"
	"net/http"

	"github.com/HugoCBB/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
