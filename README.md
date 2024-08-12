# Learn Golang

## Start the developer environment

1. Install (use [docker](./compose.yaml) or install natively) Postgres Database server
2. Start the postgres database server
3. Create a database named `learn-golang`
4. Rename [.env.example](./.env.example) to `.env` and provide the variable values
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
