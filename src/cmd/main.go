package main

import (
	"SCORESABER--New/src/route"
	"log"
	"net/http"
)

func main() {
	route.LoadTemplates()
	route.Routes(http.DefaultServeMux)

	log.Println("Serveur lancé sur :8080")
	http.ListenAndServe(":8080", nil)
}
