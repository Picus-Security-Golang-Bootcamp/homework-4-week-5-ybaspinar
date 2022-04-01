## Installation

-Clone repo to local
```
git clone https://github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-ybaspinar.git
```

-Check port 8080

-Create .env file
```
#Postgres
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_NAME=
DB_PASSWORD=
```
## How to use

-Run application
```
go run main.go
```
-Use postman to test or use curl
```
curl localhost:8080/books
```

## Available endpoints
```
GET all books
/books
GET book with given name
/book/{name}
GET book with given id
/byid/{id}
GET deleted book with given name
/deletedbook/{name}
DELETE book with given id
/book/{id}
```