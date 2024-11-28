-- +goose Up
CREATE SEQUENCE appuser_id_seq;

CREATE TABLE AppUser(
    UserId          INT PRIMARY KEY,
	RoleId          INT REFERENCES Role(RoleId) ON DELETE CASCADE NOT NULL,
	UserName        VARCHAR(255) NOT NULL UNIQUE,
	FirstName       VARCHAR(255) NOT NULL,
	FatherName      VARCHAR(255) NOT NULL,
	GrandFatherName VARCHAR(255) NOT NULL,
	Password        VARCHAR(255) NOT NULL,
	PhoneNumber     VARCHAR(255) NOT NULL UNIQUE,
	Address         VARCHAR(255) NOT NULL,
	Email           VARCHAR(255) UNIQUE,
	RegisteredBy 	VARCHAR(255) NOT NULL,
	RegisteredDate 	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (PhoneNumber ~ '^\+?[0-9]+$'),
    CHECK (Email ~* '^[^\s@]+@[^\s@]+\.[^\s@]+$')
);

-- +goose Down
DROP TABLE AppUser;
DROP SEQUENCE appuser_id_seq;






















