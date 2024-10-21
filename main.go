package main

import (
	"log"

	"github.com/aszanky/newordsbe-digistar/internal/handler"
	"github.com/aszanky/newordsbe-digistar/internal/repository"
	"github.com/aszanky/newordsbe-digistar/internal/usecase"
	"github.com/aszanky/newordsbe-digistar/pkg/database"
)

func main() {
	database, err := database.NewDatabase(".env")
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Printf("Postgres connected, Status: %#v", database.Stats())
		log.Println()
	}

	//repository
	repository := repository.NewRepository(database)

	//usecase
	usecase := usecase.NewUsecase(repository)

	//handler
	handler := handler.NewHandler(usecase)

	//RUN SERVER
	log.Println("Starting newords services on PORT 8099")
	err = handler.Start(":8099")
	if err != nil {
		log.Fatal(err)
	}
}
