package api

import (
	"data-comparision/internal/api/handlers"
	"data-comparision/internal/repository"
	"data-comparision/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine,
	fileRepo *repository.FileImportRepository,
	itemRepo *repository.LeasingItemRepository,
	importService *service.ImportService) {

	h := handlers.NewHandler(fileRepo, itemRepo, importService)
	r.POST("/upload", h.UploadExcel)
	r.GET("/diff", h.GetDiff)
}
