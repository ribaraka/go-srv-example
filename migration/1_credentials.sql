CREATE TABLE IF NOT EXISTS credentials
(
    user_id       SERIAL REFERENCES users(id),
    password_hash VARCHAR(64),
    email_token VARCHAR(64),
    created_at  TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at    TIMESTAMP NOT NULL,
);