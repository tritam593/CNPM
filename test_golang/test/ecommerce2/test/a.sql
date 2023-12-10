
CREATE DATABASE temp;
USE temp;

CREATE TABLE a (
    id INT PRIMARY KEY,
    name VARCHAR(20),
    depart INT,
    FOREIGN KEY (depart) REFERENCES b(id)
);

CREATE TABLE b (
    id INT PRIMARY KEY,
    name VARCHAR(20)
);

SELECT * FROM  users
JOIN posts
ON user_id = users.id;