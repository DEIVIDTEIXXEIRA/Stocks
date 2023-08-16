package router

import (
	"fmt"
	"log"
	"net/http"
	"stoks/controllers"

	"github.com/gorilla/mux"
)

func Router(porta string) {
	r := mux.NewRouter()


	r.HandleFunc("/", controllers.Saude).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/autenticacao", controllers.Autorizacao).Methods(http.MethodGet)
	r.HandleFunc("/acoes", controllers.TodosDisponiveis).Methods(http.MethodGet)
	r.HandleFunc("/moeda", controllers.ObterMoeda).Methods(http.MethodGet)
	r.HandleFunc("/relatorio", controllers.ClassificarAcoesPorCetor).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", porta), r))
}
