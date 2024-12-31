CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       telegram_id BIGINT UNIQUE NOT NULL,
                       first_name VARCHAR(255),
                       last_name VARCHAR(255),
                       role INTEGER NOT NULL,
                       is_authorized BOOLEAN NOT NULL DEFAULT FALSE,
                       created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_telegram_id ON users (telegram_id);