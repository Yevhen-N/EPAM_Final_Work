CREATE TABLE users
(
    id        serial PRIMARY KEY,
    full_name VARCHAR(50)         NOT NULL,
    email     VARCHAR(255) UNIQUE NOT NULL,
    password  VARCHAR(50)         NOT NULL,
    lock      BOOLEAN DEFAULT FALSE,
    admin     BOOLEAN DEFAULT FALSE

);

CREATE TYPE currency_type AS ENUM ('usd', 'uah', 'eur')
CREATE TABLE accounts
(
    id       serial PRIMARY KEY,
    user_id  INTEGER REFERENCES users (id) ON DELETE CASCADE,
    number   VARCHAR(50) UNIQUE NOT NULL,
    balance  INTEGER            NOT NULL,
    currency currency_type      NOT NULL,
    lock     BOOLEAN DEFAULT FALSE

);

CREATE TABLE cards
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    number     VARCHAR(16) UNIQUE NOT NULL
);

CREATE TYPE payment_status AS ENUM ('prepared', 'sent')
CREATE TABLE payments
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    date       TIMESTAMP      NOT NULL DEFAULT NOW(),
    sum        INTEGER        NOT NULL,
    status     payment_status NOT NULL
);

CREATE TYPE request_status AS ENUM ('new', 'approved')
CREATE TABLE requests
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    date       TIMESTAMP NOT NULL DEFAULT NOW(),
    status     request_status     DEFAULT 'new'
);

CREATE TABLE log
(
    id      serial PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    date    TIMESTAMP   NOT NULL DEFAULT NOW(),
    action  VARCHAR(10) NOT NULL
)
