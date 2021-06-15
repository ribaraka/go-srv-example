CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    firstname VARCHAR(10) NOT NULL,
    lastname  VARCHAR(10) NOT NULL,
    email     VARCHAR(100) NOT NULL UNIQUE,
    verified  BOOLEAN DEFAULT FALSE,
    location    VARCHAR(10),
    number    VARCHAR(10),
    nickname    VARCHAR(10),
    avatar    VARCHAR(10) NOT NULL
);