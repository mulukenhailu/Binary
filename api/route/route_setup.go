package route

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sql.DB, gin *gin.Engine) {

	proctedRoute := gin.Group("")
	NewRoleRouter(env, timeout, db, proctedRoute)
}