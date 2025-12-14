package parser

import (
	"data-comparision/internal/parser/models"
	"data-comparision/internal/utils"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

//
// Интерфейс
//

type ExcelParser interface {
	Parse(filePath string, fileImportID string) ([]models.LeasingItem, error)
}

//
// Реализация
//

type excelParser struct{}

func NewExcelParser() ExcelParser {
	return &excelParser{}
}

//
// Маппинг колонок Excel → поля модели
//

var columnMap = map[string]string{
	"Предмет лизинга":             "leasing_subject",
	"VIN / Зав.№":                 "vin",
	"Договор лизинга":             "leasing_contract",
	"Утвержденная цена":           "approved_price",
	"Утвержденная цена начальная": "initial_approved_price",
	"Цена реализации":             "sale_price",
	"Статус":                      "status",
	"Местонахождение":             "location",
	"Дата статуса":                "status_date",
}

//
// Основной метод парсинга
//

func (p *excelParser) Parse(filePath string, fileImportID string) ([]models.LeasingItem, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("open excel file: %w", err)
	}

	sheetName := file.GetSheetName(0)
	if sheetName == "" {
		return nil, fmt.Errorf("excel has no sheets")
	}

	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("read rows: %w", err)
	}

	if len(rows) < 2 {
		return nil, fmt.Errorf("excel file has no data rows")
	}

	// Заголовки
	headers := rows[0]
	colIndex := make(map[int]string)

	for idx, h := range headers {
		h = strings.TrimSpace(h)
		if h != "" {
			colIndex[idx] = h
		}
	}

	var items []models.LeasingItem

	// Данные
	for rowIdx, row := range rows[1:] {
		item := models.LeasingItem{
			ID:           utils.NewUUID(),
			FileImportID: fileImportID,
			Data:         make(map[string]interface{}),
		}

		for colIdx, cell := range row {
			header, ok := colIndex[colIdx]
			if !ok {
				continue
			}

			value := strings.TrimSpace(cell)
			if value == "" {
				continue
			}

			p.applyCell(&item, header, value)
		}

		// Формируем business key
		item.BusinessKey = utils.BuildBusinessKey(
			item.VIN,
			item.LeasingContract,
			item.LeasingSubject,
			utils.GetString(item.Data["Предмет лизинга. Марка"]),
			utils.GetInt(item.Data["Год выпуска"]),
		)

		// Если ключ пустой — пропускаем строку
		if item.BusinessKey == "" {
			continue
		}

		items = append(items, item)

		_ = rowIdx // можно использовать для логирования
	}

	return items, nil
}

//
// Обработка одной ячейки
//

func (p *excelParser) applyCell(item *models.LeasingItem, header string, value string) {
	switch columnMap[header] {

	case "leasing_subject":
		item.LeasingSubject = value

	case "vin":
		item.VIN = value

	case "leasing_contract":
		item.LeasingContract = value

	case "approved_price":
		item.ApprovedPrice = utils.ParseFloat(value)

	case "initial_approved_price":
		item.InitialApprovedPrice = utils.ParseFloat(value)

	case "sale_price":
		item.SalePrice = utils.ParseFloat(value)

	case "status":
		item.Status = value

	case "location":
		item.Location = value

	case "status_date":
		item.StatusDate = utils.ParseDate(value)

	default:
		// Все остальные колонки сохраняем как есть
		item.Data[header] = value
	}
}
