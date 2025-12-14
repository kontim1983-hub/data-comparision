package models

import "time"

type FileImport struct {
	ID         string    `db:"id"`
	FileName   string    `db:"file_name"`
	UploadedAt time.Time `db:"uploaded_at"`
	Status     string    `db:"status"`
}
