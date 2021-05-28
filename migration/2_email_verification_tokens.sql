CREATE TABLE IF NOT EXISTS email_verification_tokens
(
    user_id            SERIAL REFERENCES users(id),
    verification_token VARCHAR(64),
    generated_at       timestamptz NOT NULL DEFAULT (now())
);