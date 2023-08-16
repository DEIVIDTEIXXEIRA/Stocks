package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Porta aonde a API vai rodar
var (
	Porta = ""
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta = os.Getenv("PORTA")
}