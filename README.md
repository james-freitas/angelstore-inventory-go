# Angelstore Inventory (Golang)
Application that manages a product inventory fed by product entry or outflow 

## ✔️ Requirements
- Golang 1.21

## 🍔 Stack
 - Golang 1.2
 - MySQL


## ✈️ How to run locally
```shell

## start the infra
docker-compose up -d

### connect to the application container command line
docker-compose exec goapp bash

### execute application
go run cmd/app/main.go

### test using curl
curl http://localhost:8000/categories
```
