package main

import (
	"log"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("sin conexion a la db")
		return
	}
	handlers.Handlers()
}
