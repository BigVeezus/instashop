# Assessment: Build an E-commerce API

## Objective:

Create a RESTful API for an e-commerce application. This API will handle basic CRUD operations for products and orders, and provide user management and authentication.

## Requirements

This RESTFUL API was built using:

- GIN (Golang Backend framework)
- MySQL (Database)
- GORM (Golang MySQL Driver)

## To START

- Make sure you have a `.env` file inside the app folder that has the neccessary values (Look at the .env.default file for the variable examples)

- Make sure that if you choose to run the SQL locally, make sure its installed or if hosted, make sure the url is correct.

- enter into the app directory on your terminal `cd app`

- Run `go build -o app` to compile the app or Run `go run main.go` for dev mode (Make sure you are in the app folder on your terminal while running this)

- Finally run `./app`

## Documentation

Link to the postman documentation is here

https://documenter.getpostman.com/view/39858378/2sAYJ6CfTJ

## Bonus

Used SQL transactions on create orders endpoint to show how transactions can be used to exceute multiple queries and make sure they all pass before they are all stored in the DB.
