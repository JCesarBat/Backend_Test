package main

import (
	"Goolang_backend_Project/database/GORM"
	database "Goolang_backend_Project/database/sqlc"
	"Goolang_backend_Project/server"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	dsn       = "host=localhost user= root password=secret dbname= TestDB  port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB_DRIVER = "postgres"
	db_source = "postgresql://root:secret@localhost:5432/TestDB?sslmode=disable"
)

func main() {
	GormStore, err := GORM.NewGormStore(dsn)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(DB_DRIVER, db_source)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	SqlStore := database.NewStore(conn)

	server := server.StartServer(GormStore, SqlStore)

	err = server.App.Listen(":3000")
	if err != nil {
		log.Fatal("error to connect the  port ", err)
	}
}
