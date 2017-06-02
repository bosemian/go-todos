package main

import (
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.POST("/tasks", func(c echo.Context) error { return c.JSON(200, "POST Tasks") })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	e.Logger.Fatal(e.Start(":1323"))

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