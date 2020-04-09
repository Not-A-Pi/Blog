# Not-A-Pi Blog
Db finalize Branch
=======

## Introduction

> A simple blog created utilizing Golang and PostgreSQL. Our purpose is to share tech, programming, and security news as well as our insights.

## Installation

> Make sure go is installed and your $GOPATH is setup

Installing docker container
- docker run --name notapi_db -p 5432:5432 -e POSTGRES_PASSWORD={password} -d postgres
- docker exec -it notapi_db psql -U postgres
- create database notapi;

Setup Environment file
- cp sample.env .env
- set the correct password after PASSWD=
- update DBUSER is optional but recommended

Initializing Web App
- go run init.go

Running Web App
- go run main.go (The server will start on 8080)
