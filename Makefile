create-db:
	psql postgres -c "create database weather_monster;"

migrate-dev-db:
	migrate -path ./db/migrations/ -database postgres://localhost:5432/weather_monster?sslmode=disable up

migrate-test-db:
	migrate -path ./db/migrations/ -database postgres://localhost:5432/weather_monster_test?sslmode=disable up
