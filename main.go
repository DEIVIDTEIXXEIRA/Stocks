package main

import (
	"fmt"
	"stoks/config"
	"stoks/router"
)

func main() {

	fmt.Println("API OK")
	config.Carregar()
	PORTA := config.Porta
	router.Router(PORTA)

}
