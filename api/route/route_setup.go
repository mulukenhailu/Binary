package route

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/bootstrap"
)

func Setup(env *bootstrap.Env, db *sql.DB, gin *gin.Engine) {

	proctedRoute := gin.Group("")
	NewRoleRouter(env, db, proctedRoute)
}