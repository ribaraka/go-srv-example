CREATE TABLE users
(
    id        SERIAL PRIMARY KEY,
    firstname VARCHAR(10) NOT NULL,
    lastname  VARCHAR(10) NOT NULL,
    email     VARCHAR(20) NOT NULL UNIQUE,
    password  VARCHAR(64) NOT NULL,
    verified  VARCHAR(5)  DEFAULT 'false'
);


