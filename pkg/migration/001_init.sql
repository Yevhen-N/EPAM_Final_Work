CREATE TABLE users
(
    id        serial PRIMARY KEY,
    full_name VARCHAR(50)         NOT NULL,
    email     VARCHAR(255) UNIQUE NOT NULL,
    password  VARCHAR(50)         NOT NULL,
    lock      BOOLEAN DEFAULT FALSE,
    admin     BOOLEAN DEFAULT FALSE

);

CREATE TABLE accounts
(
    id       serial PRIMARY KEY,
    user_id  INTEGER REFERENCES users (id) ON DELETE CASCADE,
    number   VARCHAR(50) UNIQUE NOT NULL,
    balance  INTEGER            NOT NULL,
    currency VARCHAR(3)         NOT NULL,
    lock     BOOLEAN DEFAULT FALSE

);

CREATE TABLE cards
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    number     VARCHAR(16) UNIQUE NOT NULL
);

CREATE TABLE payments
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    date       TIMESTAMP NOT NULL DEFAULT NOW(),
    sum        INTEGER   NOT NULL,
    confirm    BOOLEAN            DEFAULT FALSE
);

CREATE TABLE requests
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    date       TIMESTAMP   NOT NULL DEFAULT NOW(),
    status     VARCHAR(10) NOT NULL
);

CREATE TABLE log
(
    id      serial PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    date    TIMESTAMP   NOT NULL DEFAULT NOW(),
    action  VARCHAR(10) NOT NULL
)
