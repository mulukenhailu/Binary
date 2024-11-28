package bootstrap

import "github.com/jackc/pgx/v5/pgxpool"

type Application struct {
	Env        *Env
	Conn      *pgxpool.Pool
}

func App() Application{
	app := &Application{}
	app.Env = NewEnv()
	app.Conn = NewPgConn(app.Env)
	return *app
}


