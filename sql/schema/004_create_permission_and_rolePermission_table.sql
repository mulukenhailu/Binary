-- +goose Up
CREATE TABLE Permission(
    PermissionId SERIAL PRIMARY KEY,
    PermissionName VARCHAR(100) NOT NULL UNIQUE,
    RegisteredBy VARCHAR(100) NOT NULL,
    RegisteredDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE RolePermission(
    RolePermissionId    SERIAL PRIMARY KEY,
    RoleId              INT REFERENCES Role(RoleId) ON DELETE CASCADE NOT NULL,
    PermissinoId        INT REFERENCES Permission(PermissionId) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE Permission;
DROP TABLE RolePermission;