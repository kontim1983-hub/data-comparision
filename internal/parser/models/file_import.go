package models

import "time"

type FileImport struct {
	ID         string    `db:"id"`          // UUID импорта
	FileName   string    `db:"file_name"`   // имя файла
	UploadedBy string    `db:"uploaded_by"` // кто загрузил
	Hash       string    `db:"hash"`        // опционально: контрольная сумма
	Status     string    `db:"status"`      // PROCESSING / DONE / ERROR
	UploadedAt time.Time `db:"uploaded_at"` // время загрузки
}
