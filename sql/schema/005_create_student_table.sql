-- +goose Up
CREATE TABLE Student(
    StudentInformationId  SERIAL PRIMARY KEY, 
    StudentId             VARCHAR(100) NOT NULL UNIQUE,
	FirstName             VARCHAR(100) NOT NULL,
	FatherName            VARCHAR(100) NOT NULL,
	GrandFatherName       VARCHAR(100) NOT NULL,
	YearOfRegistration    VARCHAR(50) NOT NULL,
	PhoneNumber           VARCHAR(100) NOT NULL UNIQUE,
	Religion              VARCHAR(100) NOT NULL,
	Sex                   VARCHAR(6) NOT NULL CHECK (Sex IN ('Male', 'Female')),
    Status                VARCHAR(10) NOT NULL CHECK (Status IN ('Active', 'Inactive')),
    CardNumber            VARCHAR(100) UNIQUE,
	Photo                 Text UNIQUE,
    RegisteredBy          VARCHAR(100) NOT NULL,
    RegisteredDate        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (PhoneNumber ~ '^\+?[0-9]+$')
);

-- +goose Down
DROP TABLE Student;