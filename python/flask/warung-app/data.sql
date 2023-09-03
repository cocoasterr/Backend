CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    status VARCHAR(10) NOT NULL
);

CREATE TABLE person (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users(id),
    fullname VARCHAR(255),
    address VARCHAR(255),
    birth DATE,
    gender VARCHAR(10),
    phone_number VARCHAR(20)
);

CREATE INDEX idx_users_id ON users(id);
CREATE INDEX idx_person_id ON person(id);
