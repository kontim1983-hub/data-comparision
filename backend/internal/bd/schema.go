package bd

const Schema = `
CREATE TABLE IF NOT EXISTS file_imports (
    id VARCHAR(255) PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status VARCHAR(50) NOT NULL DEFAULT 'pending'
);

CREATE TABLE IF NOT EXISTS leasing_items (
    id VARCHAR(255) PRIMARY KEY,
    file_import_id VARCHAR(255) NOT NULL REFERENCES file_imports(id) ON DELETE CASCADE,
    business_key VARCHAR(255) NOT NULL,
    vin VARCHAR(255),
    leasing_contract VARCHAR(255),
    leasing_subject TEXT,
    approved_price NUMERIC,
    initial_approved_price NUMERIC,
    sale_price NUMERIC,
    status VARCHAR(255),
    location VARCHAR(255),
    status_date TIMESTAMP,
    data JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(business_key, file_import_id)
);

CREATE INDEX IF NOT EXISTS idx_leasing_items_file_import ON leasing_items(file_import_id);
CREATE INDEX IF NOT EXISTS idx_leasing_items_business_key ON leasing_items(business_key);
CREATE INDEX IF NOT EXISTS idx_leasing_items_vin ON leasing_items(vin);
`
