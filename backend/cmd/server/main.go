package main

import (
	"data-comparision/internal/api"
	db "data-comparision/internal/bd/migrations"
	"data-comparision/internal/parser"
	"data-comparision/internal/repository"
	"data-comparision/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	dbConn, err := db.InitPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}

	fileRepo := repository.NewFileImportRepository(dbConn)
	itemRepo := repository.NewLeasingItemRepository(dbConn)
	parserService := parser.NewExcelParser()

	importService := service.NewImportService(fileRepo, itemRepo, parserService)
	router := gin.Default()

	api.InitRouter(router, fileRepo, itemRepo, importService)

	router.Run(":8080")
}
