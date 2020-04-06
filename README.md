# User Service Simple App

## Prerequisite

1. Go v13.x
2. PostgreSQL v11 or later
3. Dep v0.5.x

## Installation & Setup

### 1. Copy the project folder to the `$GOPATH/src/`

### 2. Create new database: `users`

### 3. Copy `.env.example` to `.env`, change the value with the real ones

### 4. Install Dependencies with dep

    $ dep ensure

### 5. Run the app

    $ go run main.go

## for test the API, import postman collection from this file:

    $ ./api/collections/Users.postman_collection.json

***

&copy; 2020 SNVL