-- name: CreateDevice :one
INSERT INTO Device(
    SerialNumber,
    Port,
    IpAddress,
    Name,
    Campus,
    BlockNumber,
    RegisteredBy
)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: DeleteDevice :one
DELETE FROM Device
WHERE DeviceId = $1
RETURNING *;

-- name: UpdateDevice :one
UPDATE Device
SET SerialNumber        = $2,
    Port                = $3,
    IpAddress           = $4,
    Name                = $5,
    Campus              = $6,
    BlockNumber         = $7,
    RegisteredBy        = $8
WHERE DeviceId          = $1
RETURNING *;


-- name: FetchDevices :many
SELECT * FROM Device;

-- name: FetchByCampus :many
SELECT * FROM Device
WHERE Campus = $1;


