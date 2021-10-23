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
# Working
* Users have to choose any number from the select list
1. List Articles
2. Create Articles
3. Approve Articles
4. Exit

# List Articles
All the published(admin approved) articles are listing togather.

# Create Article
Users can Create new article by entering article details such as tittle, description.
Before that users must login to the system, if the users haven't credetails then he/she can signup accroding to the instructions following in command prompt.

# Approve Article 
Users required authentication and also users with is_admin status is active, can only approve the articles.

        
