CREATE TABLE tokens (
    username VARCHAR(255) NOT NULL,
    token_string VARCHAR(500) NOT NULL,
    PRIMARY KEY (username, token_string),
    FOREIGN KEY (username) REFERENCES users (username)
);