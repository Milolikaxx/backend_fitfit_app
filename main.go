package main

import (
	"backend_fitfit_app/repository"
	"log"
)

func main() {
	_, err := repository.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	log.Println("Connection Success")
	StartServer()
}
