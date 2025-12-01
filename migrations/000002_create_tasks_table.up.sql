CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tasks (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        task VARCHAR(255) NOT NULL,
                        is_done BOOLEAN DEFAULT FALSE,
                        user_id UUID REFERENCES users(id) ON DELETE CASCADE
);