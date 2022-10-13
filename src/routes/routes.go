package routes

import (
	"api-eco360/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/cliente", controllers.BuscaTodosClientes)
	r.HandleFunc("/cliente/{codigo}", controllers.ConsultaCliente)

	log.Fatal(http.ListenAndServe(":3000", r))
}
