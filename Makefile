.PHONY:

run:
	go run cmd/main.go

migrate_up:
	migrate -path=migrations -database=postgres://user:secret@localhost/myapp?sslmode=disable up


migrate_down:
	migrate -path=migrations -database=postgres://user:secret@localhost/myapp?sslmode=disable down

migrate_goto:
	migrate -path=migrations -database=postgres://user:secret@localhost/myapp?sslmode=disable force 1

migrate_drop:
	migrate -path=migrations -database=postgres://user:secret@localhost/myapp?sslmode=disable drop