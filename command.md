# Useful command
Here a list of commands that usually used in this project development


### Create Migration File
```bash
migrate create -ext sql -dir database/migrations create_users_table
```

Note: use must install the package first on host machine [golang migrate](https://github.com/golang-migrate/migrate)

This command will create a file with .sql extension on directory database/migrations with postfix create_users_table and datetime as prefix filename

### Run Migration
Format
```bash
migrate -path [path] -database [databaseconnection] [action] N
```

Example
```bash
migrate -path database/migrations -database "mysql://root:adminlocal@tcp(127.0.0.1:3306)/fra" up 1
```
This command will run migration on database/migrations, and the database mysql location in 127.0.0.1:3306 with username root and password adminlocal by using fra database. The up action will run up file and we specify only 1 migration.

Note: If we didn't specify how many migration that run, it will run all pending migration. This also applies on down migration. If you did'nt specify, it will run down migration until the first migration.

### Build
Build main app
```bash
env GOOS=linux GOARCH=amd64 go build .
```

Build commands
ensure you change directory to "commands"
```bash
cd ./commands
```
and then build
```bash
env GOOS=linux GOARCH=amd64 go build .
```

