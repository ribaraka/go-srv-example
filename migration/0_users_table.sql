CREATE TABLE IF NOT EXISTS users
(
    id            SERIAL PRIMARY KEY,
    firstname     VARCHAR(252) NOT NULL,
    lastname      VARCHAR(252) NOT NULL,
    email         VARCHAR(252) NOT NULL UNIQUE,
    mobile_number INTEGER,
    location      VARCHAR(252) DEFAULT '',
    nickname      VARCHAR(252),
    avatar        VARCHAR(252) NOT NULL,
    verified      BOOLEAN      DEFAULT FALSE,
    created_at    timestamptz  NOT NULL DEFAULT NULL,
    update_at     timestamptz  DEFAULT NULL
);