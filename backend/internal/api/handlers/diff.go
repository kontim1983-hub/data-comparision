package handlers

import (
	"data-comparision/internal/diff"
	"net/http"

	_ "data-comparision/internal/repository"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDiff(c *gin.Context) {
	// Берем последний и предыдущий импорт
	latest, err := h.fileRepo.FindLatest(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить последний импорт"})
		return
	}

	prev, err := h.fileRepo.FindPrevious(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить предыдущий импорт"})
		return
	}

	// Получаем элементы из БД
	prevItems, err := h.itemRepo.FindByImport(c, prev.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения предыдущих элементов"})
		return
	}

	currItems, err := h.itemRepo.FindByImport(c, latest.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения текущих элементов"})
		return
	}

	// Генерируем diff
	diffs := diff.DiffItems(prevItems, currItems)

	c.JSON(http.StatusOK, diffs)
}
