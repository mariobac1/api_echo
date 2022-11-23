package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/mariobac1/api_/authorization"
	"github.com/mariobac1/api_/handler/login"
	"github.com/mariobac1/api_/handler/person"
	"github.com/mariobac1/api_/storage/postgres"
	postCommunity "github.com/mariobac1/api_/storage/postgres/community"
	postPerson "github.com/mariobac1/api_/storage/postgres/person"
	postUser "github.com/mariobac1/api_/storage/postgres/user"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("We can't load the certificates: %v", err)
	}

	connection, _ := postgres.NewPostgresDB()
	store := postPerson.New(connection)
	commu := postCommunity.New(connection)
	usr := postUser.New(connection)

	if err := commu.Migrate(); err != nil {
		log.Fatalf("commu.Migrate: %v", err)
	}

	if err = store.Migrate(); err != nil {
		log.Fatalf("store.Migrate: %v", err)
	}

	if err = usr.Migrate(); err != nil {
		log.Fatalf("store.Migrate: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	person.RoutePerson(*e, store)
	// community.RouteCommunity(e, commu)
	login.RouteUser(e, usr)

	log.Println("Server in port 8080 start")

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error in server %v\n", err)
	}
}
