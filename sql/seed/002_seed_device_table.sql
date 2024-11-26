-- +goose Up
INSERT INTO Device (SerialNumber, Port, IpAddress, Name, Campus, BlockNumber, RegisteredBy)
VALUES 
    ('SN12345A', '8080', '192.168.1.1', 'Router1', 'Main Campus', 'Block A', 'Admin'),
    ('SN12346B', '9090', '192.168.1.2', 'Switch1', 'Main Campus', 'Block B', 'Admin'),
    ('SN12347C', '7070', '192.168.1.3', 'AccessPoint1', 'West Campus', 'Block C', 'System'),
    ('SN12348D', '6060', '192.168.1.4', 'Router2', 'East Campus', 'Block D', 'Manager'),
    ('SN12349E', '5050', '192.168.1.5', 'Switch2', 'North Campus', 'Block E', 'Admin'),
    ('SN12350F', '4040', '192.168.1.6', 'AccessPoint2', 'South Campus', 'Block F', 'System');

-- +goose Down
DELETE FROM Device;
