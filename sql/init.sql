CREATE DATABASE todoUser;

CREATE TABLE users (
    id VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE todo (
    id VARCHAR(255) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    user_id VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);