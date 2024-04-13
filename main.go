package main

import (
	"com.github.dazsanchez/gophers-store/db"
	"com.github.dazsanchez/gophers-store/router"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	app := gin.Default()

	db.Init(db.GetMySQLDBManager(mysql.Config{
		User:   "gopher",
		Passwd: "my-password",
		DBName: "gopher_store",
	}))

	defer db.Close()

	router.Init(app)

	app.Run() // listen and serve on 0.0.0.0:8080
}
