package bootstrap

import "github.com/lib/pq"

type Application struct {
	Env        *Env
	Conn      *pq.Connector
}

func App() Application{
	app := &Application{}
	app.Env = NewEnv()
	app.Conn = NewPgConn(app.Env)
	return *app
}


