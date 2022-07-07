CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE weights (
	id uuid DEFAULT uuid_generate_v4(),
	"date" DATE NOT NULL,
    "max" int NOT NULL,
    "min" int NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);