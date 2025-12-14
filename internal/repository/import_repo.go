package repository

import (
	"context"
	"data-comparision/internal/parser/models"

	"github.com/jmoiron/sqlx"
)

type FileImportRepository struct {
	db *sqlx.DB
}

func NewFileImportRepository(db *sqlx.DB) *FileImportRepository {
	return &FileImportRepository{db: db}
}

func (r *FileImportRepository) Create(ctx context.Context, f models.FileImport) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO file_imports (id, file_name, uploaded_by, hash, status)
		VALUES ($1,$2,$3,$4,$5)`,
		f.ID, f.FileName, f.UploadedBy, f.Hash, f.Status)
	return err
}

func (r *FileImportRepository) FindLatest(ctx context.Context) (*models.FileImport, error) {
	var f models.FileImport
	err := r.db.GetContext(ctx, &f, `SELECT * FROM file_imports ORDER BY uploaded_at DESC LIMIT 1`)
	if err != nil {
		return nil, err
	}
	return &f, nil
}
