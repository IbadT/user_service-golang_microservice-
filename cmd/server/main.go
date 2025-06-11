package main

import (
	"log"

	"github.com/IbadT/user_service-golang_microservice-.git/internal/database"
	transportgrpc "github.com/IbadT/user_service-golang_microservice-.git/internal/transport/grpc"
	"github.com/IbadT/user_service-golang_microservice-.git/internal/user"
)

func main() {
	// Инициализация базы данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Создание репозитория и сервиса
	repository := user.NewRepository(db)
	service := user.NewService(repository)

	// Запуск gRPC сервера
	if err := transportgrpc.RunGRPC(service); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
