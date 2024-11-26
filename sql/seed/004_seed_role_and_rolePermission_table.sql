-- +goose Up
INSERT INTO Permission (PermissionName, RegisteredBy)
VALUES 
    ('CreateUser', 'Admin'),
    ('DeleteUser', 'Admin'),
    ('UpdateUser', 'Admin'),
    ('ViewReports', 'Admin'),
    ('ManageRoles', 'Admin'),
    ('AccessDashboard', 'System'),
    ('ApproveRequests', 'Manager'),
    ('GenerateInvoices', 'Accountant'),
    ('ManageInventory', 'Warehouse'),
    ('ReviewFeedback', 'HR');


INSERT INTO RolePermission (RoleId, PermissinoId)
VALUES
   
    ((SELECT RoleId FROM Role WHERE RoleName = 'Admin'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'CreateUser')),
    ((SELECT RoleId FROM Role WHERE RoleName = 'Admin'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'DeleteUser')),
    ((SELECT RoleId FROM Role WHERE RoleName = 'Admin'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'UpdateUser')),
    ((SELECT RoleId FROM Role WHERE RoleName = 'Admin'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'ViewReports')),
    ((SELECT RoleId FROM Role WHERE RoleName = 'Admin'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'ManageRoles')),

    ((SELECT RoleId FROM Role WHERE RoleName = 'Manager'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'ApproveRequests')),
    ((SELECT RoleId FROM Role WHERE RoleName = 'Manager'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'ViewReports')),

   
    ((SELECT RoleId FROM Role WHERE RoleName = 'Accountant'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'GenerateInvoices')),


    ((SELECT RoleId FROM Role WHERE RoleName = 'Warehouse'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'ManageInventory')),


    ((SELECT RoleId FROM Role WHERE RoleName = 'HR'), (SELECT PermissionId FROM Permission WHERE PermissionName = 'ReviewFeedback'));


-- +goose Down
DELETE FROM Permission;
DELETE FROM RolePermission;
