package service

import (
	"context"
	"data-comparision/internal/parser"
	"data-comparision/internal/parser/models"
	"data-comparision/internal/repository"
	"data-comparision/internal/utils"

	"time"
)

type ImportService struct {
	fileRepo *repository.FileImportRepository
	itemRepo *repository.LeasingItemRepository
	parser   parser.ExcelParser
}

func NewImportService(fRepo *repository.FileImportRepository,
	iRepo *repository.LeasingItemRepository,
	parser parser.ExcelParser) *ImportService {

	return &ImportService{fRepo, iRepo, parser}
}

func (s *ImportService) ImportFile(ctx context.Context, filePath, fileName string) error {
	importID := utils.NewUUID()

	// Сохраняем запись о файле
	s.fileRepo.Create(ctx, models.FileImport{
		ID:         importID,
		FileName:   fileName,
		UploadedAt: time.Now(),
		Status:     "PROCESSING",
	})

	// Парсим Excel
	items, err := s.parser.Parse(filePath, importID)
	if err != nil {
		return err
	}

	// Сохраняем данные
	err = s.itemRepo.SaveBatch(ctx, items)
	if err != nil {
		return err
	}

	return nil
}
