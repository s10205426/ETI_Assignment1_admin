CREATE database my_db;
#DROP database my_db;
USE my_db;
CREATE TABLE Passenger
(Username varchar(10),
FirstName varchar(20), LastName varchar(20), PhoneNo INT, Email varchar(100),
PRIMARY KEY (Username));
CREATE TABLE Driver
(Username varchar(10),
FirstName varchar(20), LastName varchar(20), PhoneNo INT, Email varchar(100), IDNo varchar(9), LicenseNo varchar(10),
PRIMARY KEY (Username));

INSERT INTO Passenger (Username, FirstName, LastName, PhoneNo, Email) VALUES
("jonn123", "John", "Tan", 12121212, "john@gmail.com");
INSERT INTO Passenger (Username, FirstName, LastName, PhoneNo, Email) VALUES
("timho1", "Tim", "Ho", 12355667, "tim@gmail.com");
SELECT * from my_db.Passenger;

INSERT INTO Driver (Username, FirstName, LastName, PhoneNo, Email, IDNo, LicenseNo) VALUES
("george7", "George", "Johnson", 90097007, "george@gmail.com", "T0122355K",  "SEA1234G");
INSERT INTO Driver (Username, FirstName, LastName, PhoneNo, Email, IDNo, LicenseNo) VALUES
("mikey33", "Mike", "Miller", 78900789, "mike@gmail.com", "T0990555J",  "SBC1005F");
SELECT * from my_db.Driver;