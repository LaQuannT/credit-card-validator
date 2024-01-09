package main

import (
	"log"

	"github.com/LaQuannT/credit-card-validator/repository"
	"github.com/LaQuannT/credit-card-validator/server"
)

const cardIssuerFilePath = "./db.cardIssuers.json"

func main() {
	store := repository.NewInMemoryCardIssuerStore(cardIssuerFilePath)
	srv := server.NewCardServer(":8080", store)
	log.Printf("card server is listening on: http://www.localhost:8080/")
	log.Fatal(srv.ListenAndServe())
}
