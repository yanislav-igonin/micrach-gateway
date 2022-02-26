# micrach-gateway
Gateway to handle many nodes of [micrach](https://github.com/yanislav-igonin/micrach).

## Motivation
Always wanted to make microservice imageboard, where each board is separate microservice.

## Prerequisites
1. Go 1.13+.
2. MongoDB.
3. Create `.env` file from `.env.example`, change env vars to your needs.

## Run
Just run:
```sh
go run main.go
```

**In development** I prefer to run it with [nodemon](https://github.com/remy/nodemon) for live reload. You can check the command in the `Makefile`.
