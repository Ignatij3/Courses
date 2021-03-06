#CREATE DATABASE `MySchool`;
USE `MySchool`;

DROP TABLE IF EXISTS `classes`;
DROP TABLE IF EXISTS `lessons`;
DROP TABLE IF EXISTS `rooms`;

CREATE TABLE `teachers` (
  `TName` VARCHAR(32) NOT NULL,
  `TSurname` VARCHAR(32) NOT NULL,
  PRIMARY KEY(`TName`, `TSurname`)
);

CREATE TABLE `classes` (
  `cNum` SMALLINT NOT NULL,
  `abc` CHARACTER(1) NOT NULL,
  `tutor_TName` VARCHAR(32) NOT NULL,
  `tutor_TSurname` VARCHAR(32) NOT NULL,
  PRIMARY KEY(`cNum`, `abc`),
  FOREIGN KEY(`tutor_TName`, `tutor_TSurname`) REFERENCES teachers(`tName`, `tSurname`)
);

CREATE TABLE `lessons` (
  `cNum` SMALLINT NOT NULL,
  `abc` CHARACTER(1) NOT NULL,
  `time` DATETIME NOT NULL,
  `lesson` VARCHAR(16) NOT NULL,
  PRIMARY KEY(`time`, `lesson`)
);

CREATE TABLE `rooms` (
  `rNum` INTEGER NOT NULL,
  `tutor_TName` VARCHAR(32) DEFAULT NULL,
  `tutor_TSurname` VARCHAR(32) DEFAULT NULL,
  PRIMARY KEY(`rNum`),
  FOREIGN KEY(`tutor_TName`, `tutor_TSurname`) REFERENCES teachers(`tName`, `tSurname`)
);

INSERT INTO `classes` (`cNum`, `abc`, `tutor_TName`, `tutor_TSurname`) VALUES
  (1, 'a', 'Riley', 'Burrowes'),
  (1, 'b', 'Mohammad', 'Lewis'),
  (1, 'c', 'Lexi', 'Tapson'),
  (1, 'd', 'Lily', 'O\'Reilly'),
  (2, 'a', 'Emily', 'Webster'),
  (2, 'b', 'Alexandra', 'Richards'),
  (2, 'c', 'Anna', 'Lee'),
  (2, 'd', 'Elena', 'Harris'),
  (3, 'a', 'Georgia', 'Nhira'),
  (3, 'b', 'Hannah', 'Tully'),
  (3, 'c', 'Harvey', 'Davis'),
  (3, 'd', 'Isabelle', 'Russell'),
  (4, 'a', 'Jackson', 'Hoaton'),
  (4, 'b', 'Jasmin', 'Cooper'),
  (4, 'c', 'Joe', 'Mandiveyi'),
  (4, 'd', 'Kiera', 'Roy'),
  (5, 'a', 'Laura', 'Dobel'),
  (5, 'b', 'Leo', 'Hernandez'),
  (5, 'c', 'Liam', 'Walker'),
  (5, 'd', 'Logan', 'Smit'),
  (6, 'a', 'Luke', 'Walsh'),
  (6, 'b', 'Madison', 'Haukozi'),
  (6, 'c', 'Maria', 'Cason'),
  (6, 'd', 'Martha', 'Gonzalez'),
  (7, 'a', 'Megan', 'Rose'),
  (7, 'b', 'Molly', 'Watson'),
  (7, 'c', 'Nicole', 'Mudarikwa'),
  (7, 'd', 'Noah', 'Li'),
  (8, 'a', 'Oliver', 'Commerford'),
  (8, 'b', 'Oscar', 'Gutierrez'),
  (8, 'c', 'Owen', 'Wratten'),
  (8, 'd', 'Penelope', 'Lopez'),
  (9, 'a', 'Poppy', 'Hughes'),
  (9, 'b', 'Rhys', 'Bismark'),
  (9, 'c', 'Rose', 'Pamberi'),
  (9, 'd', 'Ryan', 'George'),
  (10, 'a', 'Sam', 'Venter'),
  (10, 'b', 'Samuel', 'Chitsiga'),
  (10, 'c', 'Sarah', 'Reid'),
  (10, 'd', 'Scarlett', 'Bennett'),
  (11, 'a', 'Sebastian', 'Morrison'),
  (11, 'b', 'Sienna', 'Thomson'),
  (11, 'c', 'Sophia', 'Khan'),
  (11, 'd', 'Tia', 'Williams'),
  (12, 'a', 'Victoria', 'Cote'),
  (12, 'b', 'William', 'Easterbrook'),
  (12, 'c', 'Zachary', 'Netshamulivho'),
  (12, 'd', 'Zoe', 'Byrne');

INSERT INTO `teachers` (`TName`, `TSurname`) VALUES
  ('Alexandra', 'Richards'),
  ('Anna', 'Lee'),
  ('Elena', 'Harris'),
  ('Emily', 'Webster'),
  ('Georgia', 'Nhira'),
  ('Hannah', 'Tully'),
  ('Harvey', 'Davis'),
  ('Isabelle', 'Russell'),
  ('Jackson', 'Hoaton'),
  ('Jasmin', 'Cooper'),
  ('Joe', 'Mandiveyi'),
  ('Kiera', 'Roy'),
  ('Laura', 'Dobel'),
  ('Leo', 'Hernandez'),
  ('Lexi', 'Tapson'),
  ('Liam', 'Walker'),
  ('Lily', 'O\'Reilly'),
  ('Logan', 'Smit'),
  ('Luke', 'Walsh'),
  ('Madison', 'Haukozi'),
  ('Maria', 'Cason'),
  ('Martha', 'Gonzalez'),
  ('Megan', 'Rose'),
  ('Mohammad', 'Lewis'),
  ('Molly', 'Watson'),
  ('Nicole', 'Mudarikwa'),
  ('Noah', 'Li'),
  ('Oliver', 'Commerford'),
  ('Oscar', 'Gutierrez'),
  ('Owen', 'Wratten'),
  ('Penelope', 'Lopez'),
  ('Poppy', 'Hughes'),
  ('Rhys', 'Bismark'),
  ('Riley', 'Burrowes'),
  ('Rose', 'Pamberi'),
  ('Ryan', 'George'),
  ('Sam', 'Venter'),
  ('Samuel', 'Chitsiga'),
  ('Sarah', 'Reid'),
  ('Scarlett', 'Bennett'),
  ('Sebastian', 'Morrison'),
  ('Sienna', 'Thomson'),
  ('Sophia', 'Khan'),
  ('Tia', 'Williams'),
  ('Victoria', 'Cote'),
  ('William', 'Easterbrook'),
  ('Zachary', 'Netshamulivho'),
  ('Zoe', 'Byrne');
  
CREATE TABLE `times` (
  `num` INTEGER,
  `start` CHAR(5),
  PRIMARY KEY(`num`)
);

INSERT INTO `lessons` (`cNum`, `abc`, `lessonNum`, `lesson`) VALUES
  (1, 'a', '', '2020-05-04 08:40:00'),
  (1, 'b', '', '2020-05-04 08:40:00'),
  (1, 'c', '', '2020-05-04 08:40:00'),
  (1, 'd', '', '2020-05-04 08:40:00'),
  (2, 'a', '', '2020-05-04 08:40:00'),
  (2, 'b', '', '2020-05-04 08:40:00'),
  (2, 'c', '', '2020-05-04 08:40:00'),
  (2, 'd', '', '2020-05-04 08:40:00'),
  (3, 'a', '', '2020-05-04 08:40:00'),
  (3, 'b', '', '2020-05-04 08:40:00'),
  (3, 'c', '', '2020-05-04 08:40:00'),
  (3, 'd', '', '2020-05-04 08:40:00'),
  (4, 'a', '', '2020-05-04 08:40:00'),
  (4, 'b', '', '2020-05-04 08:40:00'),
  (4, 'c', '', '2020-05-04 08:40:00'),
  (4, 'd', '', '2020-05-04 08:40:00'),
  (5, 'a', '', '2020-05-04 08:40:00'),
  (5, 'b', '', '2020-05-04 08:40:00'),
  (5, 'c', '', '2020-05-04 08:40:00'),
  (5, 'd', '', '2020-05-04 08:40:00'),
  (6, 'a', '', '2020-05-04 08:40:00'),
  (6, 'b', '', '2020-05-04 08:40:00'),
  (6, 'c', '', '2020-05-04 08:40:00'),
  (6, 'd', '', '2020-05-04 08:40:00'),
  (7, 'a', '', '2020-05-04 08:40:00'),
  (7, 'b', '', '2020-05-04 08:40:00'),
  (7, 'c', '', '2020-05-04 08:40:00'),
  (7, 'd', '', '2020-05-04 08:40:00'),
  (8, 'a', '', '2020-05-04 08:40:00'),
  (8, 'b', '', '2020-05-04 08:40:00'),
  (8, 'c', '', '2020-05-04 08:40:00'),
  (8, 'd', '', '2020-05-04 08:40:00'),
  (9, 'a', '', '2020-05-04 08:40:00'),
  (9, 'b', '', '2020-05-04 08:40:00'),
  (9, 'c', '', '2020-05-04 08:40:00'),
  (9, 'd', '', '2020-05-04 08:40:00'),
  (10, 'a', '', '2020-05-04 08:40:00'),
  (10, 'b', '', '2020-05-04 08:40:00'),
  (10, 'c', '', '2020-05-04 08:40:00'),
  (10, 'd', '', '2020-05-04 08:40:00'),
  (11, 'a', '', '2020-05-04 08:40:00'),
  (11, 'b', '', '2020-05-04 08:40:00'),
  (11, 'c', '', '2020-05-04 08:40:00'),
  (11, 'd', '', '2020-05-04 08:40:00'),
  (12, 'a', '', '2020-05-04 08:40:00'),
  (12, 'b', '', '2020-05-04 08:40:00'),
  (12, 'c', '', '2020-05-04 08:40:00'),
  (12, 'd', '', '2020-05-04 08:40:00');