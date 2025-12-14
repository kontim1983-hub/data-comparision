package repository

import (
	"context"
	"data-comparision/internal/parser/models"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type LeasingItemRepository struct {
	db *sqlx.DB
}

func NewLeasingItemRepository(db *sqlx.DB) *LeasingItemRepository {
	return &LeasingItemRepository{db: db}
}

// Batch insert
func (r *LeasingItemRepository) SaveBatch(ctx context.Context, items []models.LeasingItem) error {
	if len(items) == 0 {
		return nil
	}

	vals := []interface{}{}
	valueStrings := []string{}

	for i, item := range items {
		valueStrings = append(valueStrings, fmt.Sprintf(
			"($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			i*13+1, i*13+2, i*13+3, i*13+4, i*13+5, i*13+6, i*13+7,
			i*13+8, i*13+9, i*13+10, i*13+11, i*13+12, i*13+13,
		))

		vals = append(vals,
			item.ID,
			item.FileImportID,
			item.BusinessKey,
			item.VIN,
			item.LeasingContract,
			item.LeasingSubject,
			item.ApprovedPrice,
			item.InitialApprovedPrice,
			item.SalePrice,
			item.Status,
			item.Location,
			item.StatusDate,
			item.Data,
		)
	}

	query := fmt.Sprintf(`INSERT INTO leasing_items
		(id, file_import_id, business_key, vin, leasing_contract, leasing_subject,
		approved_price, initial_approved_price, sale_price, status, location, status_date, data)
		VALUES %s ON CONFLICT (business_key, file_import_id) DO NOTHING`, strings.Join(valueStrings, ","))

	_, err := r.db.ExecContext(ctx, query, vals...)
	return err
}

// Получить все строки по импорту
func (r *LeasingItemRepository) FindByImport(ctx context.Context, importID string) ([]models.LeasingItem, error) {
	var items []models.LeasingItem
	err := r.db.SelectContext(ctx, &items, `SELECT * FROM leasing_items WHERE file_import_id=$1`, importID)
	return items, err
}
