include .env
export

install:
	# install CLI for managing db migrations
	go get -v github.com/rubenv/sql-migrate/...
	# install CLI for generating db queries
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	# install all application dependencies
	go mod tidy

# run all code generation tools
gen: wait_db
	sqlc generate
	sqlc compile

# starting app after closing IDE
start:
	docker-compose start

stop:
	docker-compose stop

status: wait_db
	sql-migrate status

rebuild:
	docker-compose down -v
	docker-compose up --build -d

migrate: wait_db status
	sql-migrate up

rollback: wait_db status
	sql-migrate down

wait_db:
	./wait-for.sh ${DB_HOST}:${DB_PORT}

fresh : rebuild gen

.PHONY: install gen start stop status rebuild migrate rollback wait_db fresh
