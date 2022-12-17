CREATE database my_db;
#DROP database my_db;
USE my_db;
CREATE TABLE Passenger
(Username varchar(10),
FirstName varchar(20), LastName varchar(20), PhoneNo INT, Email varchar(100), Password varchar(20),
PRIMARY KEY (Username));
CREATE TABLE Driver
(Username varchar(10),
FirstName varchar(20), LastName varchar(20), PhoneNo INT, Email varchar(100), Password varchar(20), IDNo varchar(9), LicenseNo varchar(10),
PRIMARY KEY (Username));

INSERT INTO Passenger (Username, FirstName, LastName, PhoneNo, Email, Password) VALUES
("jonn123", "John", "Tan", 12121212, "john@gmail.com", "pass123");
INSERT INTO Passenger (Username, FirstName, LastName, PhoneNo, Email, Password) VALUES
("timho1", "Tim", "Ho", 12355667, "tim@gmail.com", "password");
SELECT * from my_db.Passenger;

INSERT INTO Driver (Username, FirstName, LastName, PhoneNo, Email, Password, IDNo, LicenseNo) VALUES
("george7", "George", "Johnson", 90097007, "george@gmail.com", "password", "T0122355K",  "SEA1234G");
INSERT INTO Driver (Username, FirstName, LastName, PhoneNo, Email, Password, IDNo, LicenseNo) VALUES
("mikey33", "Mike", "Miller", 78900789, "mike@gmail.com", "qwerty", "T0990555J",  "SBC1005F");
SELECT * from my_db.Driver;