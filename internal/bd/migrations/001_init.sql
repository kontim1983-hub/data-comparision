-- 001_init.sql

-- Таблица загрузок Excel
CREATE TABLE file_imports (
                              id UUID PRIMARY KEY,
                              file_name TEXT NOT NULL,
                              uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
                              uploaded_by TEXT,
                              hash TEXT,
                              status TEXT DEFAULT 'PROCESSING'
);

-- Таблица объектов лизинга
CREATE TABLE leasing_items (
                               id UUID PRIMARY KEY,
                               file_import_id UUID REFERENCES file_imports(id),
                               business_key TEXT NOT NULL,
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
                               created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- Индекс на бизнес-ключ для быстрого diff
CREATE UNIQUE INDEX idx_leasing_items_key_import
    ON leasing_items (business_key, file_import_id);