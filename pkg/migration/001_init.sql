CREATE TABLE users
(
    ID        serial PRIMARY KEY,
    full_name VARCHAR(50) UNIQUE  NOT NULL,
    email     VARCHAR(255) UNIQUE NOT NULL,
    password  VARCHAR(50)         NOT NULL,
    lock      BOOLEAN DEFAULT FALSE,
    admin     BOOLEAN DEFAULT FALSE

);

CREATE TABLE accounts
(
    id       serial PRIMARY KEY,
    user_id  INT                NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (ID),
    number   VARCHAR(50) UNIQUE NOT NULL,
    PRIMARY KEY (number),
    balance  INT                NOT NULL,
    currency VARCHAR(3)         NOT NULL,
    lock     BOOLEAN DEFAULT FALSE

);

CREATE TABLE cards
(
    id         serial PRIMARY KEY,
    account_id VARCHAR(50) UNIQUE NOT NULL,
    FOREIGN KEY (account_id)
        REFERENCES accounts (number),
    number     VARCHAR(16) UNIQUE NOT NULL

);

CREATE TABLE payments
(
    id         serial PRIMARY KEY,
    account_id VARCHAR(50) UNIQUE NOT NULL,
    FOREIGN KEY (account_id)
        REFERENCES accounts (number),
    date       TIMESTAMP          NOT NULL,
    sum        INT                NOT NULL,
    confirm    BOOLEAN DEFAULT FALSE
);

CREATE TABLE requests
(
    id         serial PRIMARY KEY,
    account_id VARCHAR(50) UNIQUE NOT NULL,
    FOREIGN KEY (account_id)
        REFERENCES accounts (number),
    date       TIMESTAMP          NOT NULL,
    status     VARCHAR(10)        NOT NULL
);

CREATE TABLE log
(
    id        serial PRIMARY KEY,
    user_id INT         NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (ID),
    date      TIMESTAMP   NOT NULL,
    action    VARCHAR(10) NOT NULL

)
