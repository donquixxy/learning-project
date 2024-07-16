CREATE TABLE users (
    id INTEGER(255) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    login_token VARCHAR(255),
    password VARCHAR(255),
    phone VARCHAR(20),
    email VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);