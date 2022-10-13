package main

import (
	"api-eco360/routes"
	"fmt"
)

func main() {
	fmt.Print("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}
