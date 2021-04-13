CREATE TABLE credentials
(
    user_id       SERIAL REFERENCES users(id),
    password_hash VARCHAR(100),
    updated_at    DATE
);