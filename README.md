# WRG-go-machine-test


REQUIRED DATABASE TABLES;
# users table 
CREATE TABLE users
(
id int PRIMARY KEY AUTO,
name varchar(255),
email varchar(255),
password varchar(255),
is_admin char(2) DEFAULT 0
);
# article table
CREATE TABLE article
(
art_id int PRIMARY KEY AUTO,
uid int(10) FORIEN KEY,
tittle varchar(255),
description varchar(255),
timestamp date
status char(2) DEFAULT 0
);

        