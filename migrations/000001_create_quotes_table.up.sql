-- Filename: migrations/000001_create_quotes_table.up.sql
CREATE TABLE IF NOT EXISTS quotes (
  quotation_id bigserial PRIMARY KEY,
  quote text NOT NULL,
  quote_author text NOT NULL,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);