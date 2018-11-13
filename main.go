// todo.go
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kwangchin/go-echo-vue/handlers"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/engine/standard"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()

	e.Static("/", "public")
	e.GET("/", func(c echo.Context) (err error) {
		pusher, ok := c.Response().Writer.(http.Pusher)
		if ok {
			if err = pusher.Push("/css/bootstrap.min.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/css/font-awesome.min.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/fonts/fontawesome-webfont.woff2", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/fonts/fontawesome-webfont.woff", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/fonts/fontawesome-webfont.ttf", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/js/bootstrap.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/js/jquery.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/js/vue-resource.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/js/vue.min.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err = pusher.Push("/favicon.ico", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		return c.File("public/index.html")
	})
	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PostTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start as a web server
	// e.Logger.Fatal(e.Start(":1323"))
	e.Logger.Fatal(e.StartTLS(":443", "cert.pem", "key.pem"))
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
