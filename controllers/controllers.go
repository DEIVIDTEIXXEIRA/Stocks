package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Saude(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	msgJson, erro := json.Marshal(
		map[string]string{"Saúde da API": "OK"},
	)

	if erro != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		fmt.Printf("Error controllers.Saude -> %s\n", erro)
	}

	w.WriteHeader(200)
	w.Write(msgJson)
	fmt.Printf("package controllers.Saude -> %s\n", "OK")


}
func Autorizacao(w http.ResponseWriter, r *http.Request)              {}
func TodosDisponiveis(w http.ResponseWriter, r *http.Request)         {}
func ObterMoeda(w http.ResponseWriter, r *http.Request)               {}
func ClassificarAcoesPorCetor(w http.ResponseWriter, r *http.Request) {}
