CREATE TABLE file_imports (
                              id UUID PRIMARY KEY,
                              file_name TEXT NOT NULL,
                              uploaded_at TIMESTAMPTZ DEFAULT now(),
                              status TEXT
);

CREATE TABLE leasing_items (
                               id UUID PRIMARY KEY,
                               file_import_id UUID REFERENCES file_imports(id),
                               business_key TEXT,
                               vin TEXT,
                               leasing_contract TEXT,
                               leasing_subject TEXT,
                               approved_price NUMERIC,
                               initial_approved_price NUMERIC,
                               sale_price NUMERIC,
                               status TEXT,
                               location TEXT,
                               status_date TIMESTAMP,
                               data JSONB,
                               created_at TIMESTAMPTZ DEFAULT now()
);

CREATE UNIQUE INDEX idx_leasing_import_key ON leasing_items(business_key, file_import_id);
