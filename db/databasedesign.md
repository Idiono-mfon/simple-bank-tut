<!-- Docker command to run postgress image -->

docker run --name postgres12 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres:12-alpine

<!-- To exec into postgres image -->

docker exec -it postgres-simpleBank psql -U root

<!-- installing Golang Migrate for Linux Ubuntu 20.04 -->

curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 $GOPATH/bin/migrate

install homebrew from here

https://brew.sh/

Run the command after installation:

brew install golang-migrate

<!-- Creating migration command for Golang -->
<!-- seq (sequential version number for the migration file) -->
<!-- init_schema (migration file name) -->
<!-- -ext sql -dir ./  means the extension should be sql and the migration directory should be the current directory-->

migrate create -ext sql -dir ./ -seq init_schema

<!--Migrate Up command-->

migrate --path db/migration --database "postgresql://root:secret@localhost:5439/simple_bank?sslmode=disable" --verbose up

(note: change the path to what you want)

<!-- To create DB in postgres docker -->

createdb --username=root --owner=root simple_bank (dBName)

<!-- make migratedown -->

<!-- Install SQLC ORM to manage interaction with DB -->

sudo snap install sqlc (ubuntu)
brew install sqlc (macOs)

<!-- Note we have
GORM
SQLX
SQLC
Default Database interaction
 -->

<!-- Seeting up Go Modules to manage dependencies -->

go mod init github.com/simplebank

<!-- Tidy cleans up the go.mod and go mod verify
will verify that your dependencies
match the requirements specified in the go.mod
-->

go mod tidy
go mod verify

<!-- This generates a vendor folder for packages -->

go mod vendor

<!-- To clean dependencies cache -->
<!-- github.com/simplebank -->

go clean -modcache

 <!-- To run Go test -->

 <!-- cd into the folders where we have the test files and execute -->

go test

 <!-- Database driver to helps to communicate with database like postgres is lib/pq -->

 <!-- how to install it is -->

go get github.com/lib/pq (this is a database drive for postgresql)

go get github.com/stretchr/testify (Test package for Go)

https://youtu.be/m9gYy5U0edQ (to switch to pgx)

go get package@none (this command is used to remove a package)

go test -v -cover ./... (This runs unit tests in all the packages(folders) at once)
