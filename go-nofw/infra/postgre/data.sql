CREATE TABLE users (
    id VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE,
    fullName VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    
    CONSTRAINT user_id PRIMARY KEY (id)
);

CREATE INDEX idx_user_id ON users (id);
CREATE INDEX idx_user_username ON users (username);


CREATE TABLE product (
    id VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    stock INT,
    createdAt TIMESTAMP,
    createdBy VARCHAR(255),
    updatedAt TIMESTAMP,
    updatedBy VARCHAR(255),
    
    CONSTRAINT product_id PRIMARY KEY (id),
    CONSTRAINT fk_createdBy FOREIGN KEY (createdBy) REFERENCES users (username),
    CONSTRAINT fk_updatedBy FOREIGN KEY (updatedBy) REFERENCES users (username)
);

CREATE INDEX idx_product_id ON product (id);
CREATE INDEX idx_product_name ON product (name);

