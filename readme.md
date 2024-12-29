# Assignment

# How to run locally
- you need to have docker and docker compose installed
- run `docker compose up -d mysql`
- install go version 1.23.4
- run `go mod tidy` 
- run `go mod vendor` this is just to have all dependancies in vendor folder
- the db tables will be created (migrated automatically).
- seed the spacecrafts table with some data.
- run the application using `go run cmd/main.go` 