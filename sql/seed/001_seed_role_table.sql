-- +goose Up
INSERT INTO Role (RoleName, RegisteredBy)
VALUES 
    ('Admin', 'System'),
    ('Manager', 'System'),
    ('Accountant', 'System'),
    ('Warehouse', 'System'),
    ('HR', 'System');

-- +goose Down
DELETE From Role;
