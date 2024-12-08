-- +goose Up
INSERT INTO Student (
    StudentId, FirstName, FatherName, GrandFatherName, 
    YearOfRegistration, PhoneNumber, Religion, Sex, Status, 
    CardNumber, Photo, RegisteredBy
) 
VALUES
('S10001', 'Abebe', 'Kebede', 'Gebre', 
 2024, '+251911223344', 'Christianity', 'Male', 'active', 
 'CARD001', 'photo1.jpg', 'Admin'),

('S10002', 'Emebet', 'Bekele', 'Mamo', 
 2023, '+251922334455', 'Islam', 'Female', 'not_active', 
 'CARD002', 'photo2.jpg', 'Admin'),

('S10003', 'Tesfaye', 'Admasu', 'Melaku', 
 2022, '+251933445566', 'Judaism', 'Male', 'active', 
 'CARD003', 'photo3.jpg', 'Admin'),

('S10004', 'Selam', 'Yohannes', 'Hailu', 
 2024, '+251944556677', 'Orthodox', 'Female', 'not_active', 
 'CARD004', 'photo4.jpg', 'Admin'),

('S10005', 'Tewodros', 'Haile', 'Gebremedhin', 
 2021, '+251955667788', 'Protestant', 'Male', 'active', 
 NULL, NULL, 'Admin');


-- +goose Down
DELETE FROM Student;
