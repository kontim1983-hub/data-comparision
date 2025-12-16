package main

import (
	"data-comparision/internal/api"
	"data-comparision/internal/bd"
	_ "data-comparision/internal/bd"
	"data-comparision/internal/parser"
	"data-comparision/internal/repository"
	"data-comparision/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к БД
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/leasing?sslmode=disable"
	}

	database, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// ✅ Применяем схему (создаём таблицы)
	if _, err := database.Exec(bd.Schema); err != nil {
		log.Fatalf("Failed to apply schema: %v", err)
	}
	log.Println("Database schema applied successfully")

	// Инициализация repositories
	fileRepo := repository.NewFileImportRepository(database)
	itemRepo := repository.NewLeasingItemRepository(database)

	excelParser := parser.NewExcelParser()

	importService := service.NewImportService(
		fileRepo,
		itemRepo,
		excelParser,
	)

	// Инициализация роутера
	r := gin.Default()
	api.InitRouter(r, fileRepo, itemRepo, importService)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
