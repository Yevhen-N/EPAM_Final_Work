CREATE TYPE users_status_type AS ENUM ('active', 'blocked')
CREATE TYPE users_role_type AS ENUM ('admin', 'user')

CREATE TABLE users
(
    id        serial PRIMARY KEY,
    full_name VARCHAR(50)         NOT NULL,
    email     VARCHAR(255) UNIQUE NOT NULL,
    password  VARCHAR(50)         NOT NULL,
    status    users_status_type DEFAULT 'active',
    role      users_role_type   DEFAULT 'user'

);

CREATE TYPE accounts_currency_type AS ENUM ('usd', 'uah', 'eur')
CREATE TYPE accounts_status_type AS ENUM ('active', 'blocked')

CREATE TABLE accounts
(
    id       serial PRIMARY KEY,
    user_id  INTEGER REFERENCES users (id) ON DELETE CASCADE,
    number   VARCHAR(50) UNIQUE NOT NULL,
    balance  INTEGER            NOT NULL,
    currency accounts_currency_type DEFAULT 'uah',
    status   accounts_status_type   DEFAULT 'active'

);

CREATE TABLE cards
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    number     VARCHAR(16) UNIQUE NOT NULL
);

CREATE TYPE payments_status_type AS ENUM ('prepared', 'sent')

CREATE TABLE payments
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    date       TIMESTAMP            NOT NULL DEFAULT NOW(),
    sum        INTEGER              NOT NULL,
    status     payments_status_type NOT NULL
);

CREATE TYPE requests_status_type AS ENUM ('new', 'approved')

CREATE TABLE requests
(
    id         serial PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id) ON DELETE CASCADE,
    date       TIMESTAMP NOT NULL   DEFAULT NOW(),
    status     requests_status_type DEFAULT 'new'
);

CREATE TABLE log
(
    id      serial PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    date    TIMESTAMP    NOT NULL DEFAULT NOW(),
    action  VARCHAR(250) NOT NULL
)