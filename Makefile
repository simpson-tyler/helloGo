migrate:
	export POSTGRESQL_URL='postgres://admin:my_password@localhost:5432/postgres'
	migrate -database ${POSTGRESQL_URL} -path database/migrations/ up

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down
