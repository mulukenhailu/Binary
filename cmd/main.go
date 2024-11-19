package main

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mulukenhailu/Binary/api/route"
	"github.com/mulukenhailu/Binary/bootstrap"
)

func main() {

	app := bootstrap.App()
	env := app.Env
	conn := app.Conn

	db := sql.OpenDB(conn)
	defer db.Close()
	
	timeout := time.Duration(env.ContextTimeout) * time.Second
	r := gin.Default()

	route.Setup(env, timeout, db, r)

	r.Run(env.ServerAddress)
}
