package main

import (
	"database/sql"
	"go-todos/handlers"
	"github.com/labstack/echo"
	"go-todos/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB(config.DB_NAME)
	migrate(db)

	e := echo.New()

	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.NewTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Logger.Fatal(e.Start(config.PORT))

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
		    CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL
		    );
		`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}