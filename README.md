# Postgres & Golang CRUD API

This repository contains a simple RESTful API. It uses Go for the backend with Gorilla Mux for routing, GORM for database interactions, and MySQL for data storage.

## Table of Contents

-   [Routes](#routes)
-   [Dependencies](#dependencies)
-   [Getting Started](#getting-started)
    -   [Setting Up the Environment](#setting-up-the-environment)
    -   [Running the API](#running-the-api)

## Routes

Here is a list of available routes in the API and their corresponding HTTP methods:

-   `POST /book`: Create a new book
-   `GET /book`: Get all books
-   `GET /book/{bookID}`: Get a book by ID
-   `PUT /book/{bookID}`: Update a book by ID
-   `DELETE /book/{bookID}`: Delete a book by ID

## Dependencies

The following dependencies are required for this project:

-   `github.com/gorilla/mux v1.8.1`: A powerful URL router and dispatcher for Go
-   `github.com/jinzhu/gorm v1.9.16`: An ORM library for Go

Additionally, the following dependencies are used indirectly:

-   `github.com/go-sql-driver/mysql v1.5.0`: MySQL driver for Go
-   `github.com/jinzhu/inflection v1.0.0`: Inflection support for Gorm

## Getting Started

### Setting Up the Environment

This project uses Docker for environment setup. Ensure you have Docker and Docker Compose installed on your machine. The Docker Compose configuration is defined in a `docker-compose.yml` file. It includes a MySQL service with the following configuration:

-   MySQL 8.0
-   Root password: `rootpass`
-   Database: `mydb`
-   User: `user`
-   User password: `userpass`

To create the Docker environment and start MySQL, run:

```bash
docker-compose up -d
```
