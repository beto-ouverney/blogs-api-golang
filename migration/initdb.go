package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schemas = [4]string{
	`CREATE TABLE IF NOT EXISTS Users(
    id INT AUTO_INCREMENT NOT NULL,
    displayName VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255),
		image VARCHAR(255),
		PRIMARY KEY (id)
);`,
	`CREATE TABLE IF NOT EXISTS BlogPosts(
	id INT AUTO_INCREMENT NOT NULL,
title VARCHAR(255),
content VARCHAR(255),
userId INT,
published DATETIME,
updated DATETIME,
PRIMARY KEY (id),
FOREIGN KEY (userId) REFERENCES Users(id));`,

	`CREATE TABLE IF NOT EXISTS Categories(
	id INT AUTO_INCREMENT NOT NULL,
	name VARCHAR(255),
	PRIMARY KEY (id));`,

	`CREATE TABLE IF NOT EXISTS PostCategories(
	postId INT,
	categoryId INT,
	FOREIGN KEY (postId) REFERENCES BlogPosts(id),
	FOREIGN KEY (categoryId) REFERENCES Categories(id));`,
}

const MYSQL_CONNECTION = "root:password@(blogs_api_db)/"
const DB_NAME = "blogs_api"

func main() {
	conn, err := sqlx.Open("mysql", MYSQL_CONNECTION)
	if err != nil {
		panic(err.Error())
	}

	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + DB_NAME + " CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
	if err != nil {
		panic(err)
	}

	_, err = conn.Exec("USE " + DB_NAME)
	if err != nil {
		panic(err)
	}

	for i, s := range schemas {
		fmt.Printf("Creating schema %d\n", i)
		_, err = conn.Exec(s)
		if err != nil {
			panic(err)
		}
	}
	defer conn.Close()
	fmt.Println("Database created successfully")
}
