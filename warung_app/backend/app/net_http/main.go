package main

import (
	"log"
	"net/http"

	Controller "github.com/cocoasterr/warung_app/app/net_http/controller"
	PgConfig "github.com/cocoasterr/warung_app/infra/db/postgre"
	Repo "github.com/cocoasterr/warung_app/infra/db/postgre/repository"
	Service "github.com/cocoasterr/warung_app/service"
)

func main() {
	db, err := PgConfig.PGinit()
	if err != nil {
		log.Fatal(err)
	}
	authRepo := Repo.NewAuthRepository(db)
	authService := Service.NewAuthService(authRepo.Repository)
	authController := Controller.NewAuthController(*authService)

	http.HandleFunc("/api/register", authController.Register)

	http.ListenAndServe(":8080", nil)

}
