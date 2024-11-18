-- name: CreateUser :one
INSERT INTO AppUser(
    RoleId, 
    UserName, 
    FirstName, 
    FatherName, 
    GrandFatherName, 
    Password, 
    PhoneNumber, 
    Address, 
    Email, 
    RegisteredBy
    )
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: UpdateUser :one
UPDATE AppUser
SET 
    RoleId = $2, 
    UserName= $3, 
    FirstName = $4, 
    FatherName = $5, 
    GrandFatherName = $6, 
    Password= $7, 
    PhoneNumber = $8, 
    Address = $9, 
    Email = $10, 
    RegisteredBy = $11
WHERE UserId = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM AppUser
where UserId = $1
RETURNING *;

-- name: FetchByUserName :one
SELECT * FROM AppUser
WHERE UserName = $1;


-- name: FetchUsers :many
SELECT * FROM AppUser
ORDER BY UserId;

-- name: FetchByRoleName :many
SELECT 
AppUser.UserId, 
AppUser.RoleId, 
AppUser.UserName, 
AppUser.FirstName, 
AppUser.FatherName, 
AppUser.GrandFatherName, 
AppUser.Password,
AppUser.PhoneNumber, 
AppUser.Address, 
AppUser.Email, 
AppUser.RegisteredBy, 
AppUser.RegisteredDate 
FROM AppUser
INNER JOIN Role ON AppUser.RoleId = Role.RoleId
WHERE Role.RoleId = $1;

