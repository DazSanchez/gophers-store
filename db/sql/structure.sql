-- MySQL

CREATE DATABASE IF NOT EXISTS gopher_store COLLATE 'utf8mb4_unicode_ci';

USE gopher_store;

CREATE TABLE IF NOT EXISTS tag(
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS category(
  id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS gopher(
  id INT UNSIGNED PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  category_id INT UNSIGNED,
  status ENUM('pending', 'available', 'sold') DEFAULT 'pending',
  FOREIGN KEY(category_id) REFERENCES category(id)
);

CREATE TABLE IF NOT EXISTS gopher_tag(
  gopher_id INT UNSIGNED,
  tag_id INT UNSIGNED,
  PRIMARY KEY(gopher_id, tag_id),
  FOREIGN KEY(tag_id) REFERENCES tag(id) ON DELETE CASCADE,
  FOREIGN KEY(gopher_id) REFERENCES gopher(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS gopher_photo_url(
  gopher_id INT UNSIGNED,
  photo_url TEXT(2048),
  FOREIGN KEY(gopher_id) REFERENCES gopher(id) ON DELETE CASCADE
);