package main

import (
	"log"
	"net/http"

	PGConfig "github.com/cocoasterr/gowarungapp/infra/db/postgres"
	PGGormRepository "github.com/cocoasterr/gowarungapp/infra/db/postgres/repository/gorm"
	"github.com/cocoasterr/gowarungapp/service"
)

func sad(dsa interface{}) {

}

func main() {
	//database
	db_conn := PGConfig.PGGormConfig()
	sad(db_conn)
	personRepo := PGGormRepository.NewPGGormRepository(db_conn)
	personService := service.CreateNewPersonService(*personRepo)
	// repo:= PGGormRepository.
	//http
	port := ":8080"
	log.Println("Connected port", port)
	http.ListenAndServe(port, nil)
}
