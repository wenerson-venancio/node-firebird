package controllers

import (
	"api-eco360/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pagina de teste")
}

func BuscaTodosClientes(w http.ResponseWriter, r *http.Request) {
	TodosOsClientes := models.BuscaTodosClientes()
	json.NewEncoder(w).Encode(TodosOsClientes)
}

func ConsultaCliente(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	codigo := params["codigo"]
	ConsultaCliente := models.ClienteCodigo(codigo)
	json.NewEncoder(w).Encode(ConsultaCliente)

}
