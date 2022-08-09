# Blogs API GO LANG with DOCKER, MYSQL, JWT, SOLID and Clean Architecture #

Development of an API and DATABASE for a blog using JWT for security

# Other Versions of this project with other languages #

[Blogs API build with Node.Js, Docker, Sequelize, JWT and MSC Architecture Version](https://github.com/beto-ouverney/blogs-api-node-js)

## Table of contents

- [General view](#general-view)
  - [The Challenge](#the-challenge)
- [The development process](#the-development-process)
  - [Tools used](#tools-used)
  - [Lessons learned](#lessons-learned)
- [Usage](#usage)
- [Author](#author)

## General view

### The challenge

Blogs API GO LANG with Docker, Sqlx, Mysql, SOLID and Clean Architecture

**The users must be capable of**

- endpoint POST /login 
  -> The request return status 200 and a token if login is success
  -> The body of the request should follow the format below:
  
   ```json
   {
     "email": "beto.ouverney@email.com",
     "password": "123456"
   }
   ```
 -> The validation rules are:

    the email field is mandatory;
    the email field must have a valid email address;
    the password field is mandatory;
    the password field must be at least 6 characters long.

- endpoint POST /user
  -> The request return status 200 and a token if user was created with success
  -> The body of the request should follow the format below:
  
   ```json
   {
     "displayName": "Carol",
     "email": "carol@email.com",
     "password": "123456",
     "image": "http://4.bp.blogspot.com/_YA50adQ-7vQ/S1gfR_6ufpI/AAAAAAAAAAk/1ErJGgRWZDg/S45/brett.png"
   }
   ```

- endpoint GET /user
  -> The request return status 200 and all users in database
  -> The request header must be contains a valid token
  
- endpoint GET /user/:id
  -> Return a user according to id
  -> The request header must be contains a valid token
  
- endpoint POST /categories
  -> Add a category.
  -> The body of the request should follow the format below:
  
  ```json
  {
    "name": "Go Lang"
  }
  ```
   -> The validation rules are:

    the name field is mandatory;
    
- endpoint GET /categories
  -> Return all categories.
  -> The request header must be contains a valid token
  
- endpoint POST /post
  -> Add a blog post
  -> The request header must be contains a valid token
  -> The body of the request should follow the format below:
  
  ```json
  {
    "title": "Latest updates, August 1st",
    "content": "The whole text for the blog post goes here in this key",
    "categoryIds": [1, 2]
  }
  ```
  
- endpoint GET /post
  -> Return all blogs post, user owner of it and database categories
  -> The request header must be contains a valid token
  
- endpoint GET /post/:id
  -> Return a blog post, user owner of it and database categories according to id
  -> The request header must be contains a valid token
  
- endpoint PUT /post/:id
  -> Update a blog post
  -> The request header must be contains a valid token
  -> The body of the request should follow the format below:
  
  ```json
  {
    "title": "Latest updates, August 1st",
    "content": "The whole text for the blog post goes here in this key"
  }
  ```

- endpoint DELETE /post/:id
  -> Delete a blog post
  -> The request header must be contains a valid token
  -> Only the blog post creator can delete it
  
- endpoint DELETE /user/me
  -> Delete you from the database, based on the id in your token
  -> The request header must be contains a valid token
  
- endpoint GET /post/search?q=:searchTerm
  -> Return an array of blog posts that contain in their title or content the term passed in the URL
  -> The request header must be contains a valid token
  -> The query params of the request should follow the format below:
     
     http://localhost:PORT/post/search?q=wow
  
  
## The development process

### Tools used

#### Back-end

- Go Lang

### Lessons learned

In this project I could improve my knowledge in back-end, by:

- Make my own router without other frameworks.
- Using SOLID principles and Clean Architecture
- Docker
- Mysql
- Sqlx
- JWT

## Usage

- You will have access to various scripts, that will help you achieving what you want to do.
  - Before you start, your docker-compose needs to be at version 1.29 or higher
  - docker-compose up -d --build
  
- To init the database you must be run the file initdb.go in migration

cd migration

go run init.db

- After you can run the api with:

go run main.go
  
-   
## Author

- LinkedIn - [Alberto Ouverney Paz](https://www.linkedin.com/in/beto-ouverney-paz/)
