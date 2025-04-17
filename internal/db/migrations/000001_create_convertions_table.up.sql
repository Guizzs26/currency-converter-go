CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE conversions ( 
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), 
  "from" VARCHAR(50) NOT NULL, 
  "to" VARCHAR(50) NOT NULL, 
  amount NUMERIC(20, 4) NOT NULL, 
  exchange_rate NUMERIC(20, 6) NOT NULL, 
  converted_amount NUMERIC(20, 4) NOT NULL, 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW() 
);