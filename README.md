# Learn Golang

## Start the developer environment

1. Install (use [docker](./compose.yaml) or install natively) Postgres Database server.
2. Start the postgres database server.
3. Create a database named `learn-golang`.
4. Create a new file `.env` and populate it referring the [.env.example](./.env.example).
5. Run the [main.go](./main.go) with the command below:
```sh
go run main.go
```

## Todo

* Authorization (casbin)
* CORS
* Admin role should be allowed to access user role's APIs as well
* Logging (logrus)
* Dockerize
* Modules, packages, multi module
* MongoDB
