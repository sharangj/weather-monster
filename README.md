# Weather Monster

An app to forecast weather of a city.

## Prerequisites

- Postgres
- Golang
- Go Migrations. Refer `https://github.com/golang-migrate/migrate`.

## Running the app

- Create the db by running `make create-db`.
- Add the database string to `config/development.yaml`.
- Run `make run` to run the migrations, build the app and run the server.

## Running with docker

- Run `docker-compose up` to run postgres and the app.
- The dev config files need to point to the bootstrapped db.

## Running tests

- Run `make test` to run all the tests
- Run only controller tests with `go test ./test/controllers/ -v`
- Make sure you run test migrations using `make migrate-test-db`

## Running migrations

- Run `make migrate-dev-db` for dev db and Run `make migrate-test-db` before running tests.

## Create the database

- Run `psql postgres -c "create database weather_monster;"` or `make create-db`

## Important Libraries

- Gorm for db connections
- Go Migration. Refer `https://github.com/golang-migrate/migrate`
- Gin for Building API's.
- Ginkgo as a BDD test framework.
- Gomega for test matchers.
