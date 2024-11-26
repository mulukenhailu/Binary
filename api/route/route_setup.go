package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/Binary/bootstrap"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {

	protectedRoute := gin.Group("")
	NewRoleRouter(env, timeout, db, protectedRoute)
	NewDeviceRouter(env, timeout, db, protectedRoute)
	NewSignupRouter(env, timeout, db, protectedRoute)
	NewUserManagerRouter(env, timeout, db, protectedRoute)
	NewLoginRouter(env, timeout, db, protectedRoute)
}