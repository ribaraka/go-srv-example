--in case of emergency.
DROP TABLE users;

CREATE TABLE users
(
    id        SERIAL PRIMARY KEY,
    firstname VARCHAR(10) NOT NULL,
    lastname  VARCHAR(10) NOT NULL,
    email     VARCHAR(20) NOT NULL UNIQUE,
    password  VARCHAR(64) NOT NULL,
    verified  VARCHAR(5) DEFAULT 'false'
);


CREATE TABLE email_verification_tokens
(
    user_id integer REFERENCES users (id),
    verificationToken VARCHAR(20),
    generated_at      VARCHAR(20)

);

CREATE TABLE credentials
(
        user_id integer REFERENCES users (id),
        password_hash VARCHAR(20),
        salt VARCHAR(20),
        updated_at DATE
);


