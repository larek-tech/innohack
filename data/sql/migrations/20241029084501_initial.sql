-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS user_profiles(
    id BIGSERIAL PRIMARY KEY,
    email text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    otp_secret text NOT NULL,
    verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);
CREATE TABLE IF NOT EXISTS email_account(
    user_id BIGINT PRIMARY KEY,
    password text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY(user_id) REFERENCES user_profiles(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS yandex_oauth(
    user_id BIGINT PRIMARY KEY,
    access_token TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY(user_id) REFERENCES user_profiles(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS user_sessions(
    id BIGSERIAL PRIMARY KEY,
    session_id TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    is_valid BOOLEAN DEFAULT TRUE NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES user_profiles(id) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS user_sessions;
DROP TABLE IF EXISTS email_account;
DROP TABLE IF EXISTS yandex_oauth;
DROP TABLE IF EXISTS user_profiles;
-- +goose StatementEnd