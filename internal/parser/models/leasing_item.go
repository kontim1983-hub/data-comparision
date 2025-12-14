package models

import "time"

type LeasingItem struct {
	ID           string `db:"id"`
	FileImportID string `db:"file_import_id"`

	// --- Business key ---
	BusinessKey     string `db:"business_key"`
	VIN             string `db:"vin"`
	LeasingContract string `db:"leasing_contract"`

	// --- Core fields ---
	LeasingSubject string `db:"leasing_subject"` // 👈 НОВОЕ ПОЛЕ

	ApprovedPrice        *float64   `db:"approved_price"`
	InitialApprovedPrice *float64   `db:"initial_approved_price"`
	SalePrice            *float64   `db:"sale_price"`
	Status               string     `db:"status"`
	Location             string     `db:"location"`
	StatusDate           *time.Time `db:"status_date"`

	// --- Other fields ---
	Data map[string]interface{} `db:"data" json:"data"`

	CreatedAt time.Time `db:"created_at"`
}
