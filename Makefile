create-db:
	psql postgres -c "create database weather_monster;"

migrate-dev-db:
	bin/db-migrate development

migrate-test-db:
	bin/db-migrate test

test:
	make migrate-test-db && go test ./tests/controllers/ -v

run:
	make migrate-dev-db && go build && ./weather_monster
