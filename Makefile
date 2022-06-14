include .env
export

init:
	# install CLI for managing db migrations
	go get -v github.com/rubenv/sql-migrate/...

	# install CLI for generating db queries
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

	# install all application dependencies
	go mod tidy

	# initialize docker stuff
	docker-compose up -d

# run all code generation tools
gen:
	sqlc generate

# starting app after closing IDE
start:
	docker-compose start

stop:
	docker-compose stop

rebuild:
	docker-compose down -v
	docker-compose up -d

status:
	sql-migrate status

migrate:
	sql-migrate up
	sql-migrate status

rollback:
	sql-migrate down
	sql-migrate status

.PHONY: init start status stop rebuild migrate
