## COMMAND

### DB Migration

```
migrate -database postgres://DB_USER:DB_PASSWORD@DB_HOST:DB_PORT/DB_NAME?sslmode=disable -path db/migrations up

```

### Build

```
env GOOS=linux GOARCH=amd64 go build -o ./build

```
