CREATE TABLE IF NOT EXISTS credentials
(
    user_id       SERIAL REFERENCES users(id),
    password_hash VARCHAR(64),
    updated_at    timestamptz NOT NULL DEFAULT (now())
);