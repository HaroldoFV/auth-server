CREATE TABLE users
(
    id         UUID PRIMARY KEY,
    name       VARCHAR(50)        NOT NULL,
    email      VARCHAR(50) UNIQUE NOT NULL,
    password   VARCHAR(255)       NOT NULL,
    status     VARCHAR(15)        NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);