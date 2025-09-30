CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tasks (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       task VARCHAR(255) NOT NULL,
                       is_done BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMP DEFAULT NULL
);