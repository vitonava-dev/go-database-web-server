CREATE DATABASE book_store;

USE book_store;

CREATE TABLE book (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publication VARCHAR(255) NOT NULL
);


SELECT * FROM books;