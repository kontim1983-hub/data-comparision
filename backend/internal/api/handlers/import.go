package handlers

import (
	"data-comparision/internal/repository"
	"data-comparision/internal/service"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	fileRepo      *repository.FileImportRepository
	itemRepo      *repository.LeasingItemRepository
	importService *service.ImportService
}

func NewHandler(fRepo *repository.FileImportRepository, iRepo *repository.LeasingItemRepository, importService *service.ImportService) *Handler {
	return &Handler{
		fileRepo:      fRepo,
		itemRepo:      iRepo,
		importService: importService,
	}
}

// UploadExcel - обработка загрузки файла Excel
func (h *Handler) UploadExcel(c *gin.Context) {
	file, err := c.FormFile("excel")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не выбран файл"})
		return
	}

	// Сохраняем файл во временную папку
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}

	path := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить файл"})
		return
	}

	// Вызываем сервис импорта
	if err := h.importService.ImportFile(c.Request.Context(), path, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
