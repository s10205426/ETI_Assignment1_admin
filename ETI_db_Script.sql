CREATE database ETI_db;
#DROP database ETI_db;
USE ETI_db;
CREATE TABLE Passenger
(PassengerID INT(5) NOT NULL AUTO_INCREMENT,
FirstName varchar(20), LastName varchar(20), PhoneNo INT, Email varchar(100),
PRIMARY KEY (PassengerID));
CREATE TABLE Driver
(DriverID INT(5) NOT NULL AUTO_INCREMENT,
FirstName varchar(20), LastName varchar(20), PhoneNo INT, Email varchar(100), IDNo varchar(9), LicenseNo varchar(10),
PRIMARY KEY (DriverID));

SELECT * from ETI_db.Passenger;
INSERT INTO Passenger (PassengerID, FirstName, LastName, PhoneNo, Email) VALUES
("1", "John", "Tan", 12121212, "john@gmail.com");
INSERT INTO Passenger (PassengerID, FirstName, LastName, PhoneNo, Email) VALUES
("2", "Tim", "Ho", 12355667, "tim@gmail.com");

SELECT * from ETI_db.Driver;
INSERT INTO Driver (DriverID, FirstName, LastName, PhoneNo, Email, IDNo, LicenseNo) VALUES
("1", "George", "Johnson", 90097007, "george@gmail.com", "T0122355K",  "SEA1234G");
INSERT INTO Driver (DriverID, FirstName, LastName, PhoneNo, Email, IDNo, LicenseNo) VALUES
("2", "Mike", "Miller", 78900789, "mike@gmail.com", "T0990555J",  "SBC1005F");


