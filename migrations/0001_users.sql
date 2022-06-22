CREATE TABLE users (
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    password bytea NOT NULL,
    email VARCHAR(255) NOT NULL
);