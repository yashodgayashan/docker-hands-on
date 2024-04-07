-- init.sql
CREATE DATABASE IF NOT EXISTS book_store;
USE book_store;

CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date DATE
);

-- Inserting some initial data
INSERT INTO books (title, author, published_date) VALUES ('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10');
INSERT INTO books (title, author, published_date) VALUES ('1984', 'George Orwell', '1949-06-08');
