-- +goose Up
CREATE TABLE Permission(
    PermissionId SERIAL PRIMARY KEY,
    PermissionName VARCHAR(100) NOT NULL UNIQUE,
    RegisteredBy VARCHAR(100) NOT NULL,
    RegisteredDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE RolePermission(
    RolePermissionId    SERIAL NOT NULL,
    RoleId              INT REFERENCES Role(RoleId) ON DELETE CASCADE NOT NULL,
    PermissinoId        INT REFERENCES Permission(PermissionId) ON DELETE CASCADE NOT NULL,
    PRIMARY KEY(RolePermissionId, RoleId, PermissinoId),
    UNIQUE (RoleId, PermissinoId)
);

-- +goose Down
DROP TABLE RolePermission;
DROP TABLE Permission;
