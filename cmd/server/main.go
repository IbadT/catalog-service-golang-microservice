package main

import (
	"log"

	"github.com/IbadT/catalog-service-golang-microservice.git/internal/catalog"
	"github.com/IbadT/catalog-service-golang-microservice.git/internal/database"
	transportgrpc "github.com/IbadT/catalog-service-golang-microservice.git/internal/transport"
)

func main() {
	database, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	repository := catalog.NewRepository(database)
	service := catalog.NewService(repository)

	if err := transportgrpc.RunGRPC(service); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
