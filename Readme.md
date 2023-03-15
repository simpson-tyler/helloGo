

# Small Project to Test Go


I wrote this project to start learning Go and see what this language can do for a simple HTTP Server. 

I used a variety of open source tooling to make this project work and tried to document where I took code directly from the internet. 

This project uses:
- Gorilla /mux for routing
- Gorm for ORM
- Postgresql as the backing DB
- Docker as the container strategry
- Go Migrate to handle Databse Migrations
- Makefile to make it all happen

## Spinning up this project


1. git clone this repository
1. `make up`
1. `make migrate`

