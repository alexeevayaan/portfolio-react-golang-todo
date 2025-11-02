BEGIN;

CREATE TABLE IF NOT EXISTS todo
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    title      TEXT,
    description       TEXT,
    completed  BOOLEAN
);

COMMIT;