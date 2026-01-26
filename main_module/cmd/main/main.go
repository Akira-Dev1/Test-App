package main

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"log"

	"main_module/internal/app"
)


func main() {
	db, err := sql.Open(
		"postgres", 
		"host=postgres port=5432 user=core password=core dbname=core_db sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(db)
	app.Run()
}