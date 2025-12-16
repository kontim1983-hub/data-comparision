package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type LeasingItem struct {
	ID                   string
	FileImportID         string
	BusinessKey          string
	VIN                  string
	LeasingContract      string
	LeasingSubject       string
	ApprovedPrice        *float64 // ✅ Теперь указатель (может быть NULL)
	InitialApprovedPrice *float64 // ✅ Указатель
	SalePrice            *float64 // ✅ Указатель
	Status               string
	Location             string
	StatusDate           *time.Time // ✅ Используем time.Time вместо string
	Data                 JSONMap
}

// JSONMap остаётся как было
type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan JSONMap: value is not []byte")
	}
	return json.Unmarshal(bytes, j)
}
