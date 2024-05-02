#migrate create -ext sql -dir database/migrations -seq <migration-name>
postgres:
	 docker run --name Database -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it Database createdb --username=root --owner=root TestDB

dropdb:
	docker exec -it Database dropdb TestDB

migrateup:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/TestDB?sslmode=disable" -verbose up


migratedown:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/TestDB?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server :
	go run main.go