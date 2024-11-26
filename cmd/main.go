package main

import (
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

	// db := sql.OpenDB(conn)
	defer conn.Close()
	
	timeout := time.Duration(env.ContextTimeout) * time.Second
	r := gin.Default()

	route.Setup(env, timeout, conn, r)

	r.Run(env.ServerAddress)
}
