-- +goose Up
CREATE TABLE Role (
    RoleId SERIAL PRIMARY KEY,
    RoleName VARCHAR(255) NOT NULL UNIQUE,
    RegisteredBy VARCHAR(255) NOT NULL,
    RegisteredDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE Role;