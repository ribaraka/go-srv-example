CREATE TABLE credentials
(
    user_id       INT REFERENCES users (id),
    password_hash VARCHAR(20),
    salt          VARCHAR(20),
    updated_at    DATE
);