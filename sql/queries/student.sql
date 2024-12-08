-- name: CreateStudent :one
INSERT INTO Student (
    StudentId, 
    FirstName, 
    FatherName, 
    GrandFatherName, 
    YearOfRegistration, 
    PhoneNumber, 
    Religion, 
    Sex, 
    Status, 
    CardNumber, 
    Photo, 
    RegisteredBy
) 
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: UpdateStudent :one
UPDATE Student 
SET
    StudentId = $2, 
    FirstName = $3, 
    FatherName = $4, 
    GrandFatherName = $5, 
    YearOfRegistration = $6, 
    PhoneNumber = $7, 
    Religion = $8, 
    Sex = $9, 
    Status = $10, 
    CardNumber = $11,  
    Photo = $12,  
    RegisteredBy = $13
WHERE StudentInformationId = $1
RETURNING *;

-- name: DeleteStudent :one
DELETE FROM Student
WHERE StudentInformationId = $1
RETURNING *;

-- name: FetchStudents :many
SELECT * FROM Student
ORDER BY StudentInformationId;


-- name: FetchByStudentId :one
SELECT * FROM Student
WHERE StudentId = $1;

