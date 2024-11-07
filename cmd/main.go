package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/api/route"
	"github.com/mulukenhailu/Binary/bootstrap"
	_ "github.com/lib/pq" 
)

func main() {


	app := bootstrap.App()
	env := app.Env
	conn := app.Conn

	db := sql.OpenDB(conn)
	defer db.Close()

	r := gin.Default()

	route.Setup(env, db, r)

	r.Run(env.ServerAddress)
}
