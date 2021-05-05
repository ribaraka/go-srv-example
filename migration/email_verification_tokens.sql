CREATE TABLE IF NOT EXISTS email_verification_tokens
(
    user_id           INT  REFERENCES users (id),
    verificationToken VARCHAR(20),
    generated_at      timestamptz NOT NULL DEFAULT (now())
);