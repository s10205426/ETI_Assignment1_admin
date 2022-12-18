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
CREATE TABLE CarTrip
(ID INT(5) NOT NULL AUTO_INCREMENT,
PassengerUsername varchar(10), DriverUsername varchar(10),
Pickup char(6), Dropoff char(6), PickupTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP, IsCompleted char(1),
PRIMARY KEY (ID));

INSERT INTO CarTrip (PassengerUsername, DriverUsername, Pickup, Dropoff, PickupTime, IsCompleted) VALUES
("jonn123", "limmy", "334059", "455098", "2020-12-18 12:52:40", "1");
INSERT INTO CarTrip (PassengerUsername, DriverUsername, Pickup, Dropoff, PickupTime, IsCompleted) VALUES
("jonn123", "mikey33", "102333", "859444", "2021-12-18 12:52:40", "1");
INSERT INTO CarTrip (PassengerUsername, DriverUsername, Pickup, Dropoff, PickupTime, IsCompleted) VALUES
("timho1", "mikey33", "102553", "859404", "2022-12-18 12:52:40", "1");
INSERT INTO CarTrip (PassengerUsername, DriverUsername, Pickup, Dropoff, PickupTime, IsCompleted) VALUES
("jonn123", "george7", "107763", "733444", "2021-10-20 12:52:40", "0");
SELECT * from my_db.CarTrip;
 
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


