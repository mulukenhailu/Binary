-- +goose Up
INSERT INTO AppUser (UserId, RoleId, UserName, FirstName, FatherName, GrandFatherName, Password, PhoneNumber, Address, Email, RegisteredBy)
VALUES 
    (nextval('appuser_id_seq'), (SELECT RoleId FROM Role WHERE RoleName = 'Admin'), 'admin_user', 'John', 'Doe', 'Smith', 'Admin@123', '+1234567890', '123 Admin St', 'admin@example.com', 'System'),
    (nextval('appuser_id_seq'), (SELECT RoleId FROM Role WHERE RoleName = 'Manager'), 'manager_user', 'Jane', 'Doe', 'Smith', 'Manager@123', '+9876543210', '456 Manager Ave', 'manager@example.com', 'System'),
    (nextval('appuser_id_seq'), (SELECT RoleId FROM Role WHERE RoleName = 'Accountant'), 'accountant_user', 'Alice', 'Doe', 'Smith', 'Accountant@123', '+1122334455', '789 Accountant Blvd', 'accountant@example.com', 'System'),
    (nextval('appuser_id_seq'), (SELECT RoleId FROM Role WHERE RoleName = 'Warehouse'), 'warehouse_user', 'Bob', 'Doe', 'Smith', 'Warehouse@123', '+9988776655', '123 Warehouse Dr', 'warehouse@example.com', 'System'),
    (nextval('appuser_id_seq'), (SELECT RoleId FROM Role WHERE RoleName = 'HR'), 'hr_user', 'Charlie', 'Doe', 'Smith', 'HR@123', '+2233445566', '789 HR Rd', 'hr@example.com', 'System');

-- +goose Down
DELETE From AppUser;
