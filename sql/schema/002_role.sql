-- +goose Up
CREATE TABLE Device (
    DeviceId SERIAL PRIMARY KEY,
    SerialNumber VARCHAR(255) UNIQUE,
    Port VARCHAR(255) NOT NULL UNIQUE,
    IpAddress VARCHAR(39) NOT NULL,  
    Name VARCHAR(100),
    Campus VARCHAR(100),
    BlockNumber VARCHAR(50),
    RegisteredBy VARCHAR(100) NOT NULL,
    RegisteredDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (Port ~ '^\d{4}$')
);

-- +goose Down
DROP TABLE Device;